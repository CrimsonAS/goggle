package sg

// A Rectangle is a node that is rendered as a rectangle.
type Rectangle struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Color         Color
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
