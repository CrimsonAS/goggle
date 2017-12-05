package private

import (
	"log"
	"sync"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

const sceneDebug = false

// DrawableNode contains a primitive node type that is directly drawable by a
// backend, fully resolved and processed in the scenegraph.
type DrawableNode struct {
	// Node is a drawable primitive node from the scene. The node is not
	// automatically copied, so this reference is not safe to keep after the
	// rendering functions return to user code unless explicitly deep copied.
	Node sg.Node
	// Transform is the full accumulated transformation of Node's geometry in
	// the scene. The transform includes X,Y translation for relative
	// coordinates as well as scale and rotate effects.
	Transform sg.Transform
}

type shadowNode struct {
	Node sg.Node
	// Transform describes coordinate transformations necessary to size and
	// position this node in the scene. Transforms are absolute and accumulated
	// from parent nodes.
	//
	// Transforms must be used for position and size during drawing and hit tests,
	// rather than a node's geometry.
	Transform sg.Transform

	Rendered *shadowNode
	Children []*shadowNode

	DrawableList []*shadowNode
	InputList    []*shadowNode
}

type SceneRenderer struct {
	Window          sg.Windowable
	InputHelper     InputHelper
	DisableParallel bool
	FullSecondPass  bool
	resolveDrawable bool
	resolveInputs   bool
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
	shadowRoot := &shadowNode{Node: root, Transform: sg.IdentityTransform}
	r.resolveTree(shadowRoot)

	// Deliver events
	if sceneDebug {
		now := time.Now()
		dResolve = now.Sub(tmPass)
		tmPass = now
		log.Printf("scene: (2/3) delivering events to %d input nodes", len(shadowRoot.InputList))
	}
	// shadowRoot.InputList is a full list of input nodes from the graph in render
	// order (bottom-to-top). Event order is strictly reversed (top-to-bottom), so
	// iterate backwards over the list to deliver events
	dirty := false
	for i := len(shadowRoot.InputList) - 1; i >= 0; i-- {
		shadowNode := shadowRoot.InputList[i]
		size := shadowNode.Transform.Size(shadowNode.Node.(sg.Sizeable).Size())
		dirty = r.InputHelper.ProcessPointerEvents(shadowNode.Transform.Translate, size.X, size.Y, shadowNode.Node) || dirty
	}

	// Compile
	if sceneDebug {
		now := time.Now()
		dEvents = now.Sub(tmPass)
		tmPass = now
	}
	if dirty || r.FullSecondPass {
		if sceneDebug {
			log.Printf("scene: (3/3) re-resolving tree after event delivery")
		}
		r.resolveInputs = false
		// ### As we implement dirty states, it'll become possible to optimize this
		// pass significantly, by reusing the shadow tree and walking to resolve only
		// dirty subtrees.
		shadowRoot = &shadowNode{Node: root, Transform: sg.IdentityTransform}
		r.resolveTree(shadowRoot)
	} else if sceneDebug {
		log.Printf("scene: (3/3) tree is clean; skipping second resolve pass")
	}

	if sceneDebug {
		now := time.Now()
		dCompile = now.Sub(tmPass)
		dTotal := now.Sub(tmStart)
		log.Printf("scene: resolved %d drawables in %s (resolve/events/compile: %s/%s/%s)", len(shadowRoot.DrawableList), dTotal, dResolve, dEvents, dCompile)
	}

	drawables := make([]DrawableNode, len(shadowRoot.DrawableList))
	for i, node := range shadowRoot.DrawableList {
		drawables[i].Node = node.Node
		drawables[i].Transform = node.Transform
	}
	return drawables
}

func (r *SceneRenderer) resolveTree(shadow *shadowNode) {
	node := shadow.Node

	// If transformable or positionable, update transform
	if snode, ok := node.(sg.Scaleable); ok {
		shadow.Transform.Scale *= snode.GetScale()
	}
	if pnode, ok := node.(sg.Positionable); ok {
		pos := pnode.Position()
		shadow.Transform.Translate = shadow.Transform.Translate.Add(pos)
	}

	// If node accepts input events and render pass is interested, add to list
	if r.resolveInputs && NodeAcceptsInputEvents(node) {
		shadow.InputList = append(shadow.InputList, shadow)
	}

	// If drawable and render pass is drawable, add to list
	if _, ok := node.(sg.Drawable); ok && r.resolveDrawable {
		shadow.DrawableList = append(shadow.DrawableList, shadow)
	}

	var subtreeWg sync.WaitGroup

	// If renderable, render and recurse
	if rnode, ok := node.(sg.Renderable); ok {
		if shadow.Rendered == nil {
			shadow.Rendered = new(shadowNode)
		}
		shadow.Rendered.Transform = shadow.Transform

		// Node may be cached if prerendered by parent
		if r.DisableParallel {
			if shadow.Rendered.Node == nil {
				shadow.Rendered.Node = rnode.Render(r.Window)
			}
			r.resolveTree(shadow.Rendered)
		} else {
			subtreeWg.Add(1)
			go func() {
				defer subtreeWg.Done()
				if shadow.Rendered.Node == nil {
					shadow.Rendered.Node = rnode.Render(r.Window)
				}
				r.resolveTree(shadow.Rendered)
			}()
		}

		// The rendered node's transform is inherited back to this
		// node; this is important for input events.
		shadow.Transform = shadow.Rendered.Transform
	}

	// Recurse to resolve children
	if pnode, ok := node.(sg.Parentable); ok {
		children := pnode.GetChildren()
		shadow.Children = make([]*shadowNode, len(children))

		// If this node is a layout, call LayoutChildren
		if lnode, ok := node.(sg.Layouter); ok {
			geo := make([]sg.Geometryable, len(children))

			for i, child := range children {
				switch n := child.(type) {
				case sg.Geometryable:
					geo[i] = n
				case sg.Renderable:
					rendered := n.Render(r.Window)
					// Cache rendered node in the child's shadowNode
					shadow.Children[i] = &shadowNode{Rendered: &shadowNode{Node: rendered}}
					if gn, ok := rendered.(sg.Geometryable); ok {
						geo[i] = gn
					}
				default:
				}
			}

			lnode.LayoutChildren(geo)
		}

		for index, child := range children {
			if shadow.Children[index] == nil {
				shadow.Children[index] = new(shadowNode)
			}
			childShadow := shadow.Children[index]
			childShadow.Node = child
			childShadow.Transform = shadow.Transform

			if r.DisableParallel {
				r.resolveTree(childShadow)
			} else {
				subtreeWg.Add(1)
				go func(c *shadowNode) {
					defer subtreeWg.Done()
					r.resolveTree(childShadow)
				}(childShadow)
			}
		}
	}

	// Wait for subtrees to resolve
	if !r.DisableParallel {
		subtreeWg.Wait()
	}

	// Append drawable and input lists from render and children
	if shadow.Rendered != nil {
		shadow.DrawableList = append(shadow.DrawableList, shadow.Rendered.DrawableList...)
		shadow.InputList = append(shadow.InputList, shadow.Rendered.InputList...)
	}

	for _, childShadow := range shadow.Children {
		shadow.DrawableList = append(shadow.DrawableList, childShadow.DrawableList...)
		shadow.InputList = append(shadow.InputList, childShadow.InputList...)
	}
}
