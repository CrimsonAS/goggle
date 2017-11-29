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

func (this *Image) Geometry() (x, y, w, h float32) {
	return this.X, this.Y, this.Width, this.Height
}

func (this *Image) SetGeometry(x, y, w, h float32) {
	this.X, this.Y, this.Width, this.Height = x, y, w, h
}

func (this *Image) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
