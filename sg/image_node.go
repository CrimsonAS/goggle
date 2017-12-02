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

func (this *ImageNode) Position() Vec2 {
	return Vec2{this.X, this.Y}
}

func (this *ImageNode) SetPosition(pos Vec2) {
	this.X, this.Y = pos.X, pos.Y
}

func (this *ImageNode) Size() Vec2 {
	return Vec2{this.Width, this.Height}
}

func (this *ImageNode) SetSize(sz Vec2) {
	this.Width, this.Height = sz.X, sz.Y
}

func (this *ImageNode) CopyDrawable() Drawable {
	re := *this
	re.Children = nil
	return &re
}
