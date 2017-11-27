package sg

// Syntactic sugar.
type TreeNode interface{}

// Nodeable is something that can have children, used by the engine.
type Nodeable interface {
	// Return a list of children (either Nodeable or Renderable) that are in
	// this thing.
	GetChildren() []TreeNode
}

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
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Color    Color
	Children []TreeNode
}

// See Nodeable.
func (this *Rectangle) GetChildren() []TreeNode {
	return this.Children
}
