package sg

// A RectangleNode is a node that is rendered as a rectangle.
type RectangleNode struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Color         Color
}

func (this *RectangleNode) GetChildren() []Node {
	return this.Children
}

func (this *RectangleNode) Position() (x, y float32) {
	return this.X, this.Y
}

func (this *RectangleNode) SetPosition(x, y float32) {
	this.X, this.Y = x, y
}

func (this *RectangleNode) Size() (w, h float32) {
	return this.Width, this.Height
}

func (this *RectangleNode) SetSize(w, h float32) {
	this.Width, this.Height = w, h
}

func (this *RectangleNode) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
