package main

import "github.com/CrimsonAS/goggle/sg"

type HoverableRectProps struct {
	sg.GeometryProps
	color        sg.Color
	hoveredColor sg.Color
}

type HoverableRectState struct {
	sg.HoverableState
	isHovered bool
}

func HoverableRectRender(props sg.Proppable, state sg.Stateable, w sg.Windowable) sg.Node {
	if state == nil {
		state = &HoverableRectState{}
		state.OnEnter = func(state sg.Stateable) {
			dstate := state.(HoverableRectState)
			dstate.isHovered = true
		}
		state.OnLeave = func(state sg.Stateable) {
			dstate := state.(HoverableRectState)
			dstate.isHovered = false
		}
	}
	dstate := state.(HoverableRectState)
	color := props.Color
	if dstate.isHovered {
		color = sg.Color{1, 1, 0, 0}
	}

	return createElement(
		&sg.RectangleNodeRender,
		sg.RectangleProps{
			props.Geometry(), /* ??? */
			color,
		},
	)
}

func MainWindowRender(props sg.Proppable, state sg.Stateable, w sg.Windowable) sg.Node {
	return createElement(
		&HoverableRectRender,
		sg.HoverableRectProps{
			sg.RectangleProps{
				sg.GeometryProps{0, 0, 100, 100},
				sg.Color{1, 1, 0, 0},
			},
		},
	)
}
