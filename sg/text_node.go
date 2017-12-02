package sg

// TextNode renders a piece of text with a given color
// It does so with no limitations on the size, or anything else. Such complexity
// should be managed some other way.
// ### how?
// ### hints like quality, kerning, shaping?
// ### hints like "this text is dynamic?, so please don't batch it ffs"?
type TextNode struct {
	Children            []Node
	X, Y, Width, Height float32
	Text                string
	Color               Color
	PixelSize           int

	// ### right now a path, I think the Renderer interface will also have to
	// make use of a font database type provider thing.
	FontFamily string
}

func (this *TextNode) GetChildren() []Node {
	return this.Children
}

func (this *TextNode) Position() Vec2 {
	return Vec2{this.X, this.Y}
}

func (this *TextNode) SetPosition(pos Vec2) {
	this.X, this.Y = pos.X, pos.Y
}

func (this *TextNode) Size() (w, h float32) {
	return this.Width, this.Height
}

func (this *TextNode) SetSize(w, h float32) {
	this.Width, this.Height = w, h
}

func (this *TextNode) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
