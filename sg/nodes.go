package sg

import (
	"fmt"
	"time"
)

// Node is the basic element in a node tree
// Nodes may implement interfaces to gain capabilities like the ability to have
// a size in the scenegraph through the Sizeable interface, or to be
// notified about input events through interfaces like Hoverable.
type Node interface{}

// Parentable is a node that can have child nodes in the tree
type Parentable interface {
	GetChildren() []Node
}

// A Windowable is the surface a renderer paints the node tree onto.
type Windowable interface {
	// Return the time since the last frame. This should be used for advancing
	// animations, to ensure that all animations advance in synchronisation.
	FrameTime() time.Duration
	GetSize() Vec2
}

// Renderable is a node that can be rendered by the engine,
// returning a tree of nodes to draw.
type Renderable interface {
	Node
	// Render is expected to return a tree of nodes representing the
	// current graphical state of this node.
	Render(w Windowable) Node
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

// Positionable is a node with a position in the scene.
// They do not necessarily have a visual representation, but will
// translate the coordinate space of any rendered or child nodes.
type Positionable interface {
	Node
	Position() Vec2
	SetPosition(Vec2)
}

// Sizeable is a node with a size in the scene.
// The size may be used for rendering, or things like hit testing.
type Sizeable interface {
	Node
	Size() Vec2
	SetSize(Vec2)
}

// Geometryable is both Positionable and Sizeable
type Geometryable interface {
	Positionable
	Sizeable
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
	if _, yes := node.(Positionable); yes {
		re = append(re, "positionable")
	}
	if _, yes := node.(Sizeable); yes {
		re = append(re, "sizeable")
	}
	if _, yes := node.(Hoverable); yes {
		re = append(re, "hoverable")
	}
	if _, yes := node.(Pressable); yes {
		re = append(re, "pressable")
	}
	if _, yes := node.(Moveable); yes {
		re = append(re, "moveable")
	}
	if _, yes := node.(Tappable); yes {
		re = append(re, "tappable")
	}
	if _, yes := node.(Scaleable); yes {
		re = append(re, "scaleable")
	}
	if _, yes := node.(Rotateable); yes {
		re = append(re, "rotateable")
	}
	return re
}
