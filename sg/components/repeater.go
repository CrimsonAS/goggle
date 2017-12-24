package components

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func Repeater(cprops PropType, state *RenderState) sg.Node {
	rp := cprops.(RepeaterProps)
	dstate, _ := state.NodeState.(*repeaterState)
	if dstate == nil {
		dstate = &repeaterState{}
		state.NodeState = dstate
	}
	dstate.childs = []sg.Node{}

	for i := 0; i < rp.Model; i++ {
		dstate.childs = append(dstate.childs, rp.New(i))
	}

	return nodes.Parent{Children: dstate.childs}
}

type FactoryFunction func(index int) sg.Node

type RepeaterProps struct {
	New   FactoryFunction
	Model int // ### interface{}?
}

type repeaterState struct {
	childs []sg.Node
}
