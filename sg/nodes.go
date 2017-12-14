package sg

import (
	"fmt"
	"time"
)

// Node is the basic element used in a component tree
type Node interface{}

// Syntactic sugar: state is an abstract type that holds data a componet calculates
// for its own use (or use in children), in its own implementation.
type StateType interface{}

// Syntactic sugar: props are passed from parent components to child components.
// They are things like the color, or the text that the parent wants a reusable
// button component to use.
type PropType interface{}

// Parentable is a node that can have child nodes in the tree
type Parentable interface {
	GetChildren() []Node
}

type TransformNode struct {
	Matrix   Mat4
	Children []Node
}

var _ Parentable = TransformNode{}

func (this TransformNode) GetChildren() []Node {
	return this.Children
}

// A Windowable is the surface a renderer paints the node tree onto.
type Windowable interface {
	// Return the time since the last frame. This should be used for advancing
	// animations, to ensure that all animations advance in synchronisation.
	FrameTime() time.Duration
	GetSize() Vec2
}

// Useful debug/info functions for nodes
func NodeName(node Node) string {
	return fmt.Sprintf("%T", node)
}
