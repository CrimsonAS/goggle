package sg2

import (
	"github.com/CrimsonAS/goggle/sg"
)

type Geometry sg.Vec4

type StateType interface{}
type PropType interface{}

type HoverableState struct {
	OnEnter func(state StateType)
	OnLeave func(state StateType)
}

func CreateElement(...interface{}) sg.Node {
	return nil
}

type RectangleProps struct {
	Geometry Geometry
	Color    sg.Color
}

func RectangleNodeRender(props PropType, state StateType, w sg.Windowable) sg.Node {
	return nil
}
