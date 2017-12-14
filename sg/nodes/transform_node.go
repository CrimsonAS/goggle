package nodes

import "github.com/CrimsonAS/goggle/sg"

type Transform struct {
	Matrix   sg.Mat4
	Children []sg.Node
}

var _ sg.Parentable = Transform{}

func (this Transform) GetChildren() []sg.Node {
	return this.Children
}
