package nodes

import "github.com/CrimsonAS/goggle/sg"

type Rectangle struct {
	Size     sg.Vec2
	Color    sg.Color
	Children []sg.Node
}

func (this Rectangle) GetChildren() []sg.Node {
	return this.Children
}
