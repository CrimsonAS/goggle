package sg

// Text renders a piece of text with a given color
// It does so with no limitations on the size, or anything else. Such complexity
// should be managed some other way.
// ### how?
// ### hints like quality, kerning, shaping?
// ### hints like "this text is dynamic?, so please don't batch it ffs"?
type Text struct {
	Children            []Node
	X, Y, Width, Height float32
	Text                string
	Color               Color
	PixelSize           int

	// ### right now a path, I think the Renderer interface will also have to
	// make use of a font database type provider thing.
	FontFamily string
}

func (this *Text) GetChildren() []Node {
	return this.Children
}

func (this *Text) Geometry() (x, y, w, h float32) {
	return this.X, this.Y, this.Width, this.Height
}

func (this *Text) SetGeometry(x, y, w, h float32) {
	this.X, this.Y, this.Width, this.Height = x, y, w, h
}

func (this *Text) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
