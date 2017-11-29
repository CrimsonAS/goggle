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

// A Rectangle is a node that is rendered as a rectangle.
type Rectangle struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Color         Color
}

func (rect *Rectangle) GetChildren() []Node {
	return rect.Children
}

func (rect *Rectangle) Geometry() (x, y, w, h float32) {
	return rect.X, rect.Y, rect.Width, rect.Height
}

func (rect *Rectangle) SetGeometry(x, y, w, h float32) {
	rect.X, rect.Y, rect.Width, rect.Height = x, y, w, h
}

func (rect *Rectangle) CopyDrawable() Drawable {
	re := *rect
	re.Children = nil
	return &re
}

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
