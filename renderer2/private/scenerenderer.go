package private

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg2"
)

const sceneDebug = true

// DrawableNode contains a primitive node type that is directly drawable by a
// backend, fully resolved and processed in the scenegraph.
type DrawableNode struct {
	// Node is a drawable primitive node from the scene. The node is not
	// automatically copied, so this reference is not safe to keep after the
	// rendering functions return to user code unless explicitly deep copied.
	Node      sg.Node
	Transform sg2.TransformNode
}

type shadowNode struct {
	NodeInstruction sg.Node // may be a TransformNode, GeometryNode, or RenderableNode etc.
	NodeInstance    sg.Node // TransformNode, GeometryNode, etc. The resolved form of RenderableNode.
	stateInstance   *sg2.StateType
	shadowChildren  []*shadowNode
}

type SceneRenderer struct {
	Window          sg.Windowable
	InputHelper     *InputHelper
	DisableParallel bool
	FullSecondPass  bool
	resolveDrawable bool
	resolveInputs   bool

	shadowRoot shadowNode
}

func (r *SceneRenderer) Render(root sg.Node) []DrawableNode {
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
	newShadowRoot := &shadowNode{NodeInstruction: root}
	log.Printf("PRE-RENDER, tree is: %+v", r.shadowRoot)
	r.resolveTree(newShadowRoot, &r.shadowRoot)

	log.Printf("POST-RENDER, tree is: %+v", r.shadowRoot)

	if sceneDebug {
		now := time.Now()
		dCompile = now.Sub(tmPass)
		dTotal := now.Sub(tmStart)
		log.Printf("scene: resolved in %s (resolve/events/compile: %s/%s/%s)", dTotal, dResolve, dEvents, dCompile)
	}
	/*
		drawables := make([]DrawableNode, len(newShadowRoot.DrawableList))
		for i, node := range newShadowRoot.DrawableList {
			drawables[i].Node = node.Node
			drawables[i].Transform = node.Transform
		}
	*/
	return nil /* drawables */
}

func (r *SceneRenderer) resolveTree(shadow *shadowNode, oldShadow *shadowNode) {
	var node sg.Node
	var oldNode sg.Node

	node = shadow.NodeInstruction
	if oldShadow != nil {
		oldNode = oldShadow.NodeInstruction
	}

	if newRenderableNode, ok := node.(sg2.RenderableNode); ok {
		// Check if the old one was also Renderable.
		if oldRenderable, ok := oldNode.(sg2.RenderableNode); ok {
			// Check if they are the same type. If they are, then it is likely
			// that this is the same instance, so we don't need to create a new
			// node.
			newFn := fmt.Sprintf("%+v", newRenderableNode.Type) // super lame.. ###
			oldFn := fmt.Sprintf("%+v", oldRenderable.Type)
			if newFn == oldFn {
				// Set new props on the old instance, keep the old state.
				log.Printf("Preserving old node instance")
				*shadow = *oldShadow
			} else {
				log.Printf("Pointers differ (new %+v vs old %+v)", newRenderableNode.Type, oldRenderable.Type)
			}
		} else {
			log.Printf("Rendering from new node instance")
		}

		// Keep re-rendering until we fully resolve this node.
		instructionRoot := shadow.NodeInstruction
		for {
			log.Printf("Resolving instruction root %+v with props %+v state %+v", newRenderableNode.Type, newRenderableNode.Props, shadow.stateInstance)
			instructionRoot = newRenderableNode.Type(newRenderableNode.Props, &shadow.stateInstance, r.Window)
			newRenderableNode, ok = instructionRoot.(sg2.RenderableNode)
			if !ok {
				shadow.NodeInstance = instructionRoot
				break
			}
		}
	} else {
		switch node.(type) {
		case *sg2.TransformNode:
		case *sg2.GeometryNode:
		default:
			panic(fmt.Sprintf("unknown node %+v", node))
		}
		shadow.NodeInstance = node
	}

	// Recurse to resolve children
	// ### should also fetch old child
	if pnode, ok := node.(sg.Parentable); ok {
		children := pnode.GetChildren()
		shadow.shadowChildren = make([]*shadowNode, len(children))

		var subtreeWg sync.WaitGroup
		for index, child := range children {
			if shadow.shadowChildren[index] == nil {
				shadow.shadowChildren[index] = new(shadowNode)
			}
			childShadow := shadow.shadowChildren[index]
			childShadow.NodeInstruction = child

			if r.DisableParallel {
				r.resolveTree(childShadow, nil)
			} else {
				subtreeWg.Add(1)
				go func(c *shadowNode) {
					defer subtreeWg.Done()
					r.resolveTree(childShadow, nil)
				}(childShadow)
			}
		}
		// Wait for subtrees to resolve
		if !r.DisableParallel {
			subtreeWg.Wait()
		}
	}

	*oldShadow = *shadow
}
