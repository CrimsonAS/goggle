package sg

import (
	"fmt"
	"time"
)

// Node is the basic element used in a component tree
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
	GetSize() Size
}

// Useful debug/info functions for nodes
func NodeName(node Node) string {
	return fmt.Sprintf("%T", node)
}
