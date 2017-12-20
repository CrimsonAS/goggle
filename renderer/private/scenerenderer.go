package private

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

const sceneDebug = false

// DrawableNode contains a primitive node type that is directly drawable by a
// backend, fully resolved and processed in the scenegraph.
type DrawableNode struct {
	// Node is a drawable primitive node from the scene. The node is not
	// automatically copied, so this reference is not safe to keep after the
	// rendering functions return to user code unless explicitly deep copied.
	Node      sg.Node
	Transform nodes.Transform
}

// shadowNode is an entry in the shadow tree, generated by SceneRenderer during
// passes over the node tree. The shadow tree has a node for every node in the input
// scene, including rendered trees. The exact structure of the tree may differ slightly.
type shadowNode struct {
	// Original node from the input scene
	sceneNode sg.Node
	// Resolved result of calling Render() on sceneNode, if the sceneNode is a Component
	rendered *shadowNode
	// Resolved children of the node, if the sceneNode is Parentable. When a Component
	// returns a non-nil render tree, children of the node will appear under the rendered node
	// instead.
	shadowChildren []*shadowNode
	// Persisted state for the node
	state components.StateType
	// Accumulated transform for this node
	transform sg.Mat4
}

type SceneRenderer struct {
	Window          sg.Windowable
	InputHelper     *InputHelper
	EnableParallel  bool
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
	r.InputHelper.EndPointerEvents()
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
	if inputNode, ok := shadow.sceneNode.(nodes.Input); ok {
		state, _ := shadow.state.(*nodes.InputState)
		if state == nil {
			state = &nodes.InputState{}
			shadow.state = state
		}
		r.InputHelper.ProcessPointerEvents(&inputNode, shadow.transform, inputNode.Size, state)
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

// resolveNode populates a shadowNode by resolving the sceneNode it represents.
// When an oldShadow is also provided, the new tree is correlated with the old
// tree to preserve node state.
//
// resolveNode expects a shadowNode with sceneNode and transform defined,
// and all other fields to have the default value.
//
// resolveNode does not populate the list of child nodes, and does not resolve
// the rendered tree.
func (r *SceneRenderer) resolveNode(shadow *shadowNode, oldShadow *shadowNode) {
	node := shadow.sceneNode

	// If the node's actual type is the same as the old tree, we need to preserve state.
	// For any node where we're not preserving state, oldShadow is discarded here.
	if oldShadow != nil && reflect.TypeOf(oldShadow.sceneNode) != reflect.TypeOf(shadow.sceneNode) {
		if sceneDebug {
			log.Printf("Node type changed from '%T' to '%T', discarding subtree state", oldShadow.sceneNode, shadow.sceneNode)
		}
		oldShadow = nil
	}

	// Copy state from the old shadow tree
	if oldShadow != nil {
		shadow.state = oldShadow.state
	}

	switch n := node.(type) {
	case components.Component:
		// Render node
		state := components.RenderState{
			Window:    r.Window,
			NodeState: shadow.state,
			Transform: shadow.transform,
		}
		renderedNode := n.Type(n.Props, &state)
		shadow.state, shadow.transform = state.NodeState, state.Transform

		if renderedNode != nil {
			shadow.rendered = &shadowNode{sceneNode: renderedNode, transform: shadow.transform}
		}

	case nodes.Transform:
		shadow.transform = shadow.transform.MulM4(n.Matrix)

	case nodes.Rectangle:
	case nodes.Image:
	case nodes.Input:
	default:
		panic(fmt.Sprintf("unknown node %T %+v", node, node))
	}
}

// resolveTree recursively resolves the node in 'shadow', its rendered tree, and
// its children. If oldShadow is provided, state will be preserved from that tree.
func (r *SceneRenderer) resolveTree(shadow *shadowNode, oldShadow *shadowNode) {
	// Resolve this node
	r.resolveNode(shadow, oldShadow)

	if shadow.rendered != nil {
		// When there is a rendered tree (meaning, this node is a Component), any
		// children of this node are reparented to become children of the rendered
		// node. As of now, the rendered node hasn't been resolved yet. We'll
		// pre-populate its shadowChildren with the reparented children, and it can
		// pick those up during its own resolveTree pass.
		//
		// There's no need to handle oldShadow here; the state for these reparented
		// nodes will be under their final parent in the old tree, and it will
		// handle that as usual when it resolves.

		// Create placeholder shadow nodes in rendered for all children
		if pnode, ok := shadow.sceneNode.(sg.Parentable); ok {
			children := pnode.GetChildren()
			shadowChildren := make([]*shadowNode, len(children))
			for index, child := range children {
				shadowChildren[index] = &shadowNode{sceneNode: child}
			}
			shadow.rendered.shadowChildren = append(shadow.rendered.shadowChildren, shadowChildren...)
		}

		// There may also be nodes that were reparented to us; they are forwarded
		// on to this rendered node as well.
		shadow.rendered.shadowChildren = append(shadow.rendered.shadowChildren, shadow.shadowChildren...)
		shadow.shadowChildren = nil

		// Finally, recurse to resolve the rendered tree
		if oldShadow != nil {
			r.resolveTree(shadow.rendered, oldShadow.rendered)
		} else {
			r.resolveTree(shadow.rendered, nil)
		}
	} else if pnode, ok := shadow.sceneNode.(sg.Parentable); ok {
		// If the node is Parentable, recursively resolve all children. If there is an
		// oldShadow, correlate the children with the old shadow tree to preserve state.
		//
		// If this node was the head of a rendered tree, shadowChildren may already have
		// placeholder entries for children that are being reparented here. These are
		// stacked _after_ the native children of this node, but otherwise they are
		// resolved in the same way.

		// Populate shadowChildren with placeholder entries for all children
		{
			children := pnode.GetChildren()
			shadowChildren := make([]*shadowNode, len(children))
			for index, child := range children {
				shadowChildren[index] = &shadowNode{sceneNode: child}
			}
			// Prepend this to anything that is already in shadowChildren
			shadow.shadowChildren = append(shadowChildren, shadow.shadowChildren...)
		}

		// Now iterate the full list of children, copy over any render state (like the
		// transform), and recursively resolve their tree. If there are equivalent
		// nodes in the oldShadow, pass them along to preserve state. They will be
		// type-checked before use in renderNode.
		var subtreeWg sync.WaitGroup
		for index, shadowChild := range shadow.shadowChildren {
			var oldShadowChild *shadowNode
			if oldShadow != nil && len(oldShadow.shadowChildren) > index {
				oldShadowChild = oldShadow.shadowChildren[index]
			}

			shadowChild.transform = shadow.transform

			if !r.EnableParallel {
				r.resolveTree(shadowChild, oldShadowChild)
			} else {
				subtreeWg.Add(1)
				go func(c, o *shadowNode) {
					defer subtreeWg.Done()
					r.resolveTree(c, o)
				}(shadowChild, oldShadowChild)
			}
		}

		// Wait for subtrees to resolve
		if r.EnableParallel {
			subtreeWg.Wait()
		}
	}
}
