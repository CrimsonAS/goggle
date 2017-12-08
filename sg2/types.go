package sg2

import (
	"github.com/CrimsonAS/goggle/sg"
)

type Geometry sg.Vec4

type Proppable interface {
	IsProppable()
}
type Stateable interface {
	IsStateable()
}

type HoverableState struct {
	OnEnter func(state Stateable)
	OnLeave func(state Stateable)
}

func CreateElement(...interface{}) sg.Node {
	return nil
}

type RectangleProps struct {
	Geometry Geometry
	Color    sg.Color
}

func RectangleNodeRender(props Proppable, state Stateable, w sg.Windowable) sg.Node {
	return nil
}
