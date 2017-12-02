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

// An Image is a node that renders a texture.
type Image struct {
	Children      []Node
	X, Y          float32
	Width, Height float32
	Texture       Texture
}

func (this *Image) GetChildren() []Node {
	return this.Children
}

func (this *Image) Position() (x, y float32) {
	return this.X, this.Y
}

func (this *Image) SetPosition(x, y float32) {
	this.X, this.Y = x, y
}

func (this *Image) Size() (w, h float32) {
	return this.Width, this.Height
}

func (this *Image) SetSize(w, h float32) {
	this.Width, this.Height = w, h
}

func (this *Image) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
