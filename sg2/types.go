package sg2

import (
	"github.com/CrimsonAS/goggle/sg"
)

type Geometry sg.Vec4

type StateType interface{}
type PropType interface{}

// RenderableType is an internal definition for node types which
// can dynamically render nodes, e.g. RenderableNode and TouchNode.
//
// It is not intended for user code to implement RenderableType.
type RenderableType interface {
	Render(state *RenderState) sg.Node
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
var _ RenderableType = RenderableNode{}

func (this RenderableNode) GetChildren() []sg.Node {
	return this.Children
}

func (node RenderableNode) Render(state *RenderState) sg.Node {
	return node.Type(node.Props, state)
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
				Size:  sg.Vec2{rp.Geometry.Z, rp.Geometry.W},
				Color: rp.Color,
			},
		},
	}
}

// TouchNode is equivalent to a RenderableNode, except that a TouchNode
// also receives touch input events. The current touch state is passed
// as part of the TouchRenderState, which also sets event callbacks.
type TouchNode struct {
	Type     func(PropType, *TouchRenderState) sg.Node
	Props    PropType
	Children []sg.Node
}

type TouchState struct {
	// Geometry is the rectangular area in which input events are
	// accepted, specified in the coordinates of the TouchNode.
	TouchGeometry Geometry

	IsHovered bool

	OnEnter func(state *TouchState)
	OnLeave func(state *TouchState)

	userState StateType
}

type TouchRenderState struct {
	*RenderState
	*TouchState
}

var _ sg.Parentable = TouchNode{}
var _ RenderableType = TouchNode{}

func (node TouchNode) GetChildren() []sg.Node {
	return node.Children
}

func (node TouchNode) Render(rs *RenderState) sg.Node {
	// The NodeState is a TouchState. Pull that out alongside
	// the RenderState in a TouchRenderState, then swap in the
	// userState as the NodeState.
	//
	// This matches RenderableNode's API more closely, with the
	// added TouchState data in TouchRenderState.
	if rs.NodeState == nil {
		rs.NodeState = &TouchState{}
	}
	state := TouchRenderState{
		rs,
		rs.NodeState.(*TouchState),
	}
	state.NodeState = state.userState

	// Render with user function
	rendered := node.Type(node.Props, &state)

	// Reverse the above to keep the TouchState
	state.userState = state.NodeState
	rs.NodeState = state.TouchState

	return rendered
}
