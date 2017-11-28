package sg

// TreeNode is a basic node; embed BasicNode for standard implementation
type TreeNode interface {
	GetChildren() []TreeNode
}

// BasicNode provides a default embeddable implementation of TreeNode
type BasicNode struct {
	ObjectName string
	Children   []TreeNode
}

func (node *BasicNode) GetObjectName() string {
	return node.ObjectName
}

func (node *BasicNode) GetChildren() []TreeNode {
	return node.Children
}

// ### The word 'render' is extremely overloaded in this API

// Renderable is something that can be rendered, used by the engine.
type Renderable interface {
	// Render is expected to return a tree of nodes representing this thing's
	// current graphical state.
	Render() TreeNode
}

// A R G B
type Color [4]float32

// A Rectangle is a node that is rendered as a rectangle.
type Rectangle struct {
	BasicNode
	X      float32
	Y      float32
	Width  float32
	Height float32
	Color  Color
}
