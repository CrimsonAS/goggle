package nodes

import "github.com/CrimsonAS/goggle/sg"

type Texture interface{}

// A FileTexture is a local file represented by a URI.
type FileTexture string

type Image struct {
	Size     sg.Vec2
	Color    sg.Color
	Children []sg.Node
	Texture  Texture
}

func (this Image) GetChildren() []sg.Node {
	return this.Children
}
