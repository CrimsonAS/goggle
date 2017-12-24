package components

import "github.com/CrimsonAS/goggle/sg"

// Syntactic sugar: state is an abstract type that holds data a componet calculates
// for its own use (or use in children), in its own implementation.
type StateType interface{}

// Syntactic sugar: props are passed from parent components to child components.
// They are things like the color, or the text that the parent wants a reusable
// button component to use.
type PropType interface{}

type RenderState struct {
	Window    sg.Windowable
	NodeState StateType
}

type Component struct {
	Type     func(PropType, *RenderState) sg.Node
	Props    PropType
	Children []sg.Node
}

var _ sg.Parentable = Component{}

func (this Component) GetChildren() []sg.Node {
	return this.Children
}
