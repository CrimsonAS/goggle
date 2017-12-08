package main

import "github.com/CrimsonAS/goggle/sg"
import "github.com/CrimsonAS/goggle/sg2"

type HoverableRectProps struct {
	sg2.Geometry
	color        sg.Color
	hoveredColor sg.Color
}

type HoverableRectState struct {
	sg2.HoverableState
	isHovered bool
}

func HoverableRectRender(props sg2.PropType, state sg2.StateType, w sg.Windowable) sg.Node {
	if state == nil {
		state = &HoverableRectState{}
		dstate := state.(*HoverableRectState)
		dstate.OnEnter = func(state sg2.StateType) {
			dstate := state.(*HoverableRectState)
			dstate.isHovered = true
		}
		dstate.OnLeave = func(state sg2.StateType) {
			dstate := state.(*HoverableRectState)
			dstate.isHovered = false
		}
	}
	dstate := state.(*HoverableRectState)
	dprops := props.(*HoverableRectProps)
	color := dprops.color
	if dstate.isHovered {
		color = sg.Color{1, 1, 0, 0}
	}

	return sg2.CreateElement(
		sg2.RectangleNodeRender,
		sg2.RectangleProps{
			dprops.Geometry,
			color,
		},
	)
}

func MainWindowRender(props sg2.PropType, state sg2.StateType, w sg.Windowable) sg.Node {
	return sg2.CreateElement(
		HoverableRectRender,
		HoverableRectProps{
			Geometry:     sg2.Geometry{0, 0, 100, 100},
			color:        sg.Color{1, 1, 0, 0},
			hoveredColor: sg.Color{1, 1, 0, 0},
		},
	)
}
