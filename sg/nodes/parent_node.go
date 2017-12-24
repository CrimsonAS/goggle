package nodes

import "github.com/CrimsonAS/goggle/sg"

type Parent struct {
	Children []sg.Node
}

var _ sg.Parentable = Parent{}

func (node Parent) GetChildren() []sg.Node {
	return node.Children
}
