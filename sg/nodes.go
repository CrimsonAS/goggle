package sg

// TreeNode is a basic node; embed BasicNode for standard implementation
type TreeNode interface {
	GetChildren() []TreeNode
	Pos() (x, y float32)
	SetPos(x, y float32)
	Size() (w, h float32)
	SetSize(w, h float32)
}

// BasicNode provides a default embeddable implementation of TreeNode
type BasicNode struct {
	ObjectName    string
	Children      []TreeNode
	X, Y          float32
	Width, Height float32
}

func (node *BasicNode) GetObjectName() string {
	return node.ObjectName
}

func (node *BasicNode) GetChildren() []TreeNode {
	return node.Children
}

func (node *BasicNode) Pos() (x, y float32) {
	return node.X, node.Y
}

func (node *BasicNode) SetPos(x, y float32) {
	node.X, node.Y = x, y
}

func (node *BasicNode) Size() (w, h float32) {
	return node.Width, node.Height
}

func (node *BasicNode) SetSize(w, h float32) {
	node.Width, node.Height = w, h
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
	Color Color
}
