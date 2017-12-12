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

type RenderState struct {
	Window    sg.Windowable
	NodeState StateType
	Transform sg.Mat4
}

type RenderableNode struct {
	Type     func(PropType, *RenderState) sg.Node
	Props    PropType
	Children []sg.Node
}

var _ sg.Parentable = RenderableNode{}

func (this RenderableNode) GetChildren() []sg.Node {
	return this.Children
}

type RectangleProps struct {
	Geometry Geometry
	Color    sg.Color
}

type TransformNode struct {
	Matrix   sg.Mat4
	Children []sg.Node
}

var _ sg.Parentable = TransformNode{}

func (this TransformNode) GetChildren() []sg.Node {
	return this.Children
}

type GeometryNode struct {
	Material interface{} // ColorMaterial or others.. cast as needed.
	// ### mesh data...
}

type ColorMaterial sg.Color

func RectangleNodeRender(props PropType, state *RenderState) sg.Node {
	rp := props.(RectangleProps)
	return TransformNode{
		Matrix: sg.Translate2D(rp.Geometry.X, rp.Geometry.Y),
		Children: []sg.Node{
			GeometryNode{
				Material: ColorMaterial(rp.Color),
			},
		},
	}
}
