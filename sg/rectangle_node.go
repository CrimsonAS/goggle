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

func (this *RectangleNode) Position() Vec2 {
	return Vec2{this.X, this.Y}
}

func (this *RectangleNode) SetPosition(pos Vec2) {
	this.X, this.Y = pos.X, pos.Y
}

func (this *RectangleNode) Size() Vec2 {
	return Vec2{this.Width, this.Height}
}

func (this *RectangleNode) SetSize(sz Vec2) {
	this.Width, this.Height = sz.X, sz.Y
}

func (this *RectangleNode) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
