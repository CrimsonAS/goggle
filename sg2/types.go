package sg2

import (
	"github.com/CrimsonAS/goggle/sg"
)

type StateType interface{}
type PropType interface{}

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

type SimpleRectangleNode struct {
	Size     sg.Vec2
	Color    sg.Color
	Children []sg.Node
}

func (this SimpleRectangleNode) GetChildren() []sg.Node {
	return this.Children
}

func RectangleNodeRender(props PropType, state *RenderState) sg.Node {
	rp := props.(RectangleProps)
	return TransformNode{
		Matrix: sg.Translate2D(rp.Geometry.X, rp.Geometry.Y),
		Children: []sg.Node{
			SimpleRectangleNode{
				Size:  rp.Geometry.Size(),
				Color: rp.Color,
			},
		},
	}
}

type InputState struct {
	IsHovered bool
	IsGrabbed bool
	IsPressed bool

	MousePos sg.Vec2
}

// An InputNode has a size and can get input events. The current component state
// is passed in to the InputNode.
type InputNode struct {
	Geometry Geometry
	Children []sg.Node

	OnEnter   func(state InputState)
	OnLeave   func(state InputState)
	OnMove    func(state InputState)
	OnPress   func(state InputState)
	OnRelease func(state InputState)
}

var _ sg.Parentable = InputNode{}

func (node InputNode) GetChildren() []sg.Node {
	return node.Children
}
