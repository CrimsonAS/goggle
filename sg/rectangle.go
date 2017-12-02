package sg

// A Rectangle is a node that is rendered as a rectangle.
type Rectangle struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Color         Color
}

func (this *Rectangle) GetChildren() []Node {
	return this.Children
}

func (this *Rectangle) Position() (x, y float32) {
	return this.X, this.Y
}

func (this *Rectangle) SetPosition(x, y float32) {
	this.X, this.Y = x, y
}

func (this *Rectangle) Size() (w, h float32) {
	return this.Width, this.Height
}

func (this *Rectangle) SetSize(w, h float32) {
	this.Width, this.Height = w, h
}

func (this *Rectangle) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
