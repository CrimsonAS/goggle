package private

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg2"
)

const sceneDebug = false

// DrawableNode contains a primitive node type that is directly drawable by a
// backend, fully resolved and processed in the scenegraph.
type DrawableNode struct {
	// Node is a drawable primitive node from the scene. The node is not
	// automatically copied, so this reference is not safe to keep after the
	// rendering functions return to user code unless explicitly deep copied.
	Node      sg.Node
	Transform sg2.TransformNode
}

// shadowNode is an entry in the shadow tree, generated by SceneRenderer during
// passes over the node tree. The shadow tree has a node for every node in the input
// scene, including rendered trees. The exact structure of the tree may differ slightly.
type shadowNode struct {
	// Original node from the input scene
	sceneNode sg.Node
	// Resolved result of calling Render() on sceneNode, if the sceneNode is a RenderableNode
	rendered *shadowNode
	// Resolved children of the node, if the sceneNode is Parentable. When a RenderableNode
	// returns a non-nil render tree, children of the node will appear under the rendered node
	// instead.
	shadowChildren []*shadowNode
	// Persisted state for the node
	state sg2.StateType
	// Accumulated transform for this node
	transform sg.Mat4
}

type SceneRenderer struct {
	Window          sg.Windowable
	InputHelper     *InputHelper
	DisableParallel bool
	FullSecondPass  bool
	resolveDrawable bool
	resolveInputs   bool

	shadowRoot *shadowNode
}

func (r *SceneRenderer) Render(root sg.Node) {
	var tmStart, tmPass time.Time
	var dResolve, dEvents, dCompile time.Duration

	// Resolve tree
	r.resolveDrawable = true
	r.resolveInputs = true
	if sceneDebug {
		tmStart = time.Now()
		tmPass = tmStart
		log.Printf("scene: (1/3) resolving tree")
	}
	newShadowRoot := &shadowNode{sceneNode: root, transform: sg.NewIdentity()}

	const resolveDebug = false // extremely noisy

	if resolveDebug {
		log.Printf("PRE-RENDER, tree is: %+v", r.shadowRoot)
	}

	r.resolveTree(newShadowRoot, r.shadowRoot)
	r.shadowRoot = newShadowRoot

	if resolveDebug {
		log.Printf("POST-RENDER, tree is: %+v", r.shadowRoot)
	}

	if sceneDebug {
		now := time.Now()
		dCompile = now.Sub(tmPass)
		dTotal := now.Sub(tmStart)
		log.Printf("scene: resolved in %s (resolve/events/compile: %s/%s/%s)", dTotal, dResolve, dEvents, dCompile)
	}
}

// DeliverEvents walks the rendered shadow tree in input order and
// invokes callbacks or updates state in the tree as appropriate.
// DeliverEvents changes state but does not re-render the tree, so
// the shadow tree is considered dirty after calling this function.
func (r *SceneRenderer) DeliverEvents() {
	r.deliverEventsToTree(r.shadowRoot)
}

func (r *SceneRenderer) deliverEventsToTree(shadow *shadowNode) {
	if shadow == nil {
		return
	}

	// Input order is the inverse of draw order, so deliver events
	// to children first and in reverse.
	for i := len(shadow.shadowChildren) - 1; i >= 0; i-- {
		r.deliverEventsToTree(shadow.shadowChildren[i])
	}

	// If there's a rendered tree, go down it next
	if shadow.rendered != nil {
		r.deliverEventsToTree(shadow.rendered)
	}

	// Finally, try to deliver events to this node
	if inputNode, ok := shadow.sceneNode.(sg2.InputNode); ok {
		r.InputHelper.ProcessPointerEvents(inputNode, shadow.transform, inputNode.Geometry, &shadow.state)
	}
}

// Draw walks the rendered shadow tree in draw order and calls the
// nodeCallback function for each primitive node.
func (r *SceneRenderer) Draw(nodeCallback func(sg.Node, sg.Mat4)) {
	r.drawNode(r.shadowRoot, nodeCallback)
}

func (r *SceneRenderer) drawNode(shadow *shadowNode, nodeCallback func(sg.Node, sg.Mat4)) {
	if shadow == nil {
		return
	} else if shadow.rendered != nil {
		// Bypass this node for the callback, but move down the rendered tree
		r.drawNode(shadow.rendered, nodeCallback)
	} else {
		nodeCallback(shadow.sceneNode, shadow.transform)
		for _, child := range shadow.shadowChildren {
			r.drawNode(child, nodeCallback)
		}
	}
}

// resolveTree populates a shadowNode by recursively resolving the sceneNode it
// represents and any resolved trees or children. When an oldShadow is also provided,
// the new tree is correlated with the old tree to preserve node state.
func (r *SceneRenderer) resolveTree(shadow *shadowNode, oldShadow *shadowNode) {
	node := shadow.sceneNode

	// If the node's actual type is the same as the old tree, we need to preserve state.
	// Only RenderableNodes may have state, but all other nodes have to preserve the
	// tree state for any child RenderableNodes.
	//
	// For any node where we're not preserving state, oldShadow is discarded here.
	if oldShadow != nil && reflect.TypeOf(oldShadow.sceneNode) != reflect.TypeOf(shadow.sceneNode) {
		if sceneDebug {
			log.Printf("Node type changed from '%T' to '%T', discarding subtree state", oldShadow.sceneNode, shadow.sceneNode)
		}
		oldShadow = nil
	}

	if newRenderableNode, ok := node.(sg2.RenderableNode); ok {
		// Copy state from the old shadow tree
		if oldShadow != nil {
			shadow.state = oldShadow.state
		}

		// Render node
		state := sg2.RenderState{
			Window:    r.Window,
			NodeState: shadow.state,
			Transform: shadow.transform,
		}
		renderedNode := newRenderableNode.Render(&state)
		shadow.state, shadow.transform = state.NodeState, state.Transform

		// Recurse to resolve rendered tree
		if renderedNode != nil {
			shadow.rendered = &shadowNode{sceneNode: renderedNode, transform: shadow.transform}
			if oldShadow != nil {
				r.resolveTree(shadow.rendered, oldShadow.rendered)
			} else {
				r.resolveTree(shadow.rendered, nil)
			}
		}
	} else {
		switch n := node.(type) {
		case sg2.TransformNode:
			shadow.transform = shadow.transform.MulM4(n.Matrix)
		case sg2.GeometryNode:
		case sg2.SimpleRectangleNode:
		case sg2.InputNode:
		default:
			panic(fmt.Sprintf("unknown node %T %+v", node, node))
		}
	}

	// Recurse to resolve children
	if pnode, ok := node.(sg.Parentable); ok {
		children := pnode.GetChildren()
		shadowChildren := make([]*shadowNode, len(children))
		var oldShadowChildren []*shadowNode

		// If this is a RenderableNode, children of this node are actually under the
		// next non-renderable node in the shadow tree, appended to any children of that
		// node.
		//
		// In this case, the oldShadow tree also needs to compare from the correct offset
		// in the rendered node's shadowChildren.
		parentShadow, oldParentShadow := shadow, oldShadow
		for parentShadow.rendered != nil {
			parentShadow = parentShadow.rendered
			if oldParentShadow != nil {
				oldParentShadow = oldParentShadow.rendered
				// If the rendered nodes are not the same type, discard old state
				if oldParentShadow != nil && reflect.TypeOf(parentShadow.sceneNode) != reflect.TypeOf(oldParentShadow.sceneNode) {
					oldParentShadow = nil
				}
			}
		}

		prefix := len(parentShadow.shadowChildren)
		if oldParentShadow != nil && prefix < len(oldParentShadow.shadowChildren) {
			oldShadowChildren = oldParentShadow.shadowChildren[prefix:]
		}

		var subtreeWg sync.WaitGroup
		for index, child := range children {
			var oldChildShadow *shadowNode
			if index < len(oldShadowChildren) {
				oldChildShadow = oldShadowChildren[index]
			}

			shadowChildren[index] = &shadowNode{sceneNode: child, transform: parentShadow.transform}

			if r.DisableParallel {
				r.resolveTree(shadowChildren[index], oldChildShadow)
			} else {
				subtreeWg.Add(1)
				go func(c, o *shadowNode) {
					defer subtreeWg.Done()
					r.resolveTree(c, o)
				}(shadowChildren[index], oldChildShadow)
			}
		}

		// Wait for subtrees to resolve
		if !r.DisableParallel {
			subtreeWg.Wait()
		}

		// Store list of shadowChildren
		parentShadow.shadowChildren = append(parentShadow.shadowChildren, shadowChildren...)
	}
}
