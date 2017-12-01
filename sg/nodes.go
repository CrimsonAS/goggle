package sg

import (
	"fmt"
)

// Node is the basic element in a node tree
type Node interface{}

// Parentable is a node that can have child nodes in the tree
type Parentable interface {
	GetChildren() []Node
}

// Renderable is a node that can be rendered by the engine,
// returning a tree of nodes to draw.
type Renderable interface {
	Node
	// Render is expected to return a tree of nodes representing the
	// current graphical state of this node.
	Render() Node
}

// Drawable is a node that the rendering backend implements
// primitive drawing functions for. Only drawable nodes have a
// visual representation on the surface.
type Drawable interface {
	Node
	// CopyDrawable must return a copy of this instance with all information
	// necessary for drawing preserved. Drawable copies should not preserve
	// children or other unnecessary state.
	CopyDrawable() Drawable
}

type TouchPoint struct {
	X float32
	Y float32
}

// A Hoverable is a node that will get events when a point's coordintes are
// above the item. Note tht Hoverable must also implement GeometryNode for the
// scenegraph to know that the point is inside the item's boundaries.
//
// ### unsolved problems: we should also probably block propagation of hover.
// We could have a return code to block hover propagating further down the tree,
// letting someone write code like:
//
// Root UI node
//     Sidebar PointerEnter() { return true; /* block */ }
//         Button Hoverable // to highlight as need be
//     UI page
//
// This would also imply that we need to invert delivery such that we deliver to
// children first up to parents, and we also need to deliver based on paint
// order.
type Hoverable interface {
	PointerEnter(TouchPoint)
	PointerLeave(TouchPoint)
}

// A Pressable is a node that will get events when a point is pressed or
// released in its boundary.
type Pressable interface {
	PointerPressed(TouchPoint)
	PointerReleased(TouchPoint)
}

// A Movable is a node that will get events when a mouse is inside its boundary.
type Moveable interface {
	PointerMoved(TouchPoint)
}

// A Tappable is a node that will get events when a touch is pressed and released.
type Tappable interface {
	PointerTapped(TouchPoint)
}

// ### Pressable: OnPressed, OnReleased
// ### Tappable: OnTap
// ... etc. These are a bit more complicated though, as they need to "grab"
// pointers, rather than just acting as state notifiers like Hoverable.

// GeometryNode is a node with a position and size in the canvas.
// They do not necessarily have a visual representation, but will
// translate the coordinate space of any rendered or child nodes.
type GeometryNode interface {
	Node
	Geometry() (x, y, w, h float32)
	SetGeometry(x, y, w, h float32)
}

// ParentNode is a node that acts as a container for child nodes,
// provided for your tree-building convenience.
type ParentNode struct {
	Children []Node
}

func (node *ParentNode) GetChildren() []Node {
	return node.Children
}

// A R G B
type Color [4]float32

// Useful debug/info functions for nodes
func NodeName(node Node) string {
	return fmt.Sprintf("%T", node)
}

func NodeInterfaces(node Node) []string {
	var re []string
	if _, yes := node.(Parentable); yes {
		re = append(re, "parentable")
	}
	if _, yes := node.(Drawable); yes {
		re = append(re, "drawable")
	}
	if _, yes := node.(Renderable); yes {
		re = append(re, "renderable")
	}
	if _, yes := node.(GeometryNode); yes {
		re = append(re, "geometry")
	}
	return re
}
