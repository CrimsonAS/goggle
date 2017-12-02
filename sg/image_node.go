package sg

// A Texture is a type representing a remote resource for the engine to load in.
// ### memory block option also
type Texture interface {
	GetSource() string
}

// A FileTexture is a texture represented by a URI
// ### should this also handle remote resources? I'd guess so.
type FileTexture struct {
	Source string
}

func (this *FileTexture) GetSource() string {
	return this.Source
}

// An ImageNode is a node that renders a texture.
type ImageNode struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Texture       Texture
}

func (this *ImageNode) GetChildren() []Node {
	return this.Children
}

func (this *ImageNode) Position() (x, y float32) {
	return this.X, this.Y
}

func (this *ImageNode) SetPosition(x, y float32) {
	this.X, this.Y = x, y
}

func (this *ImageNode) Size() (w, h float32) {
	return this.Width, this.Height
}

func (this *ImageNode) SetSize(w, h float32) {
	this.Width, this.Height = w, h
}

func (this *ImageNode) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
