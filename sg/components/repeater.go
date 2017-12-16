package components

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func Repeater(cprops PropType, state *RenderState) sg.Node {
	rp := cprops.(RepeaterProps)
	childs := []sg.Node{}

	for i := 0; i < rp.Model; i++ {
		childs = append(childs, rp.New(i))
	}

	return nodes.Transform{
		Matrix:   sg.Translate2D(0, 0),
		Children: childs,
	}
}

type FactoryFunction func(index int) sg.Node

type RepeaterProps struct {
	New   FactoryFunction
	Model int // ### interface{}?
}
