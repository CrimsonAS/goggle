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

func CreateElement(vars ...interface{}) sg.Node {
	var props PropType
	var fptr (func(PropType, StateType, sg.Windowable) sg.Node)
	for idx, avar := range vars {
		if idx == 0 {
			fptr = avar.(func(PropType, StateType, sg.Windowable) sg.Node)
		} else if idx == 1 {
			props = avar
		}
	}

	// ### ehhhh this should really just be turned into an instruction and
	// called by the renderer with state
	return fptr(props, nil, nil)
}

type RectangleProps struct {
	Geometry Geometry
	Color    sg.Color
}

type TransformNode struct {
	Matrix   sg.Mat4
	Children []sg.Node
}

type GeometryNode struct {
	Material interface{} // ColorMaterial or others.. cast as needed.
	// ### mesh data...
}

type ColorMaterial sg.Color

func RectangleNodeRender(props PropType, state StateType, w sg.Windowable) sg.Node {
	rp := props.(RectangleProps)
	return TransformNode{
		Matrix: sg.NewIdentity(), // ### calculate from rp props
		Children: []sg.Node{
			GeometryNode{
				Material: ColorMaterial(rp.Color),
			},
		},
	}
}
