package nodes

import "github.com/CrimsonAS/goggle/sg"

type InputState struct {
	IsHovered bool
	IsGrabbed bool
	IsPressed bool

	MousePos sg.Vec2
}

// An InputNode has a size and can get input events. The current component state
// is passed in to the InputNode.
type Input struct {
	Size     sg.Vec2
	Children []sg.Node

	OnEnter   func(state InputState)
	OnLeave   func(state InputState)
	OnMove    func(state InputState)
	OnPress   func(state InputState)
	OnRelease func(state InputState)
}

var _ sg.Parentable = Input{}

func (node Input) GetChildren() []sg.Node {
	return node.Children
}
