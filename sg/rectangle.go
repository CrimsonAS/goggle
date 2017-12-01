package sg

// A Rectangle is a node that is rendered as a rectangle.
type Rectangle struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Color         Color
	Scale         float32
	Rotation      float32
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

func (this *Rectangle) GetScale() float32 {
	return this.Scale
}

func (this *Rectangle) SetScale(r float32) {
	this.Scale = r
}

func (this *Rectangle) GetRotation() float32 {
	return this.Rotation
}

func (this *Rectangle) SetRotation(r float32) {
	this.Rotation = r
}
