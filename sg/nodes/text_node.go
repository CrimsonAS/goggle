package nodes

import "github.com/CrimsonAS/goggle/sg"

type Text struct {
	Size       sg.Vec2
	PixelSize  int
	FontFamily string
	Color      sg.Color
	Text       string
	Children   []sg.Node
}

func (this Text) GetChildren() []sg.Node {
	return this.Children
}
