package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg2"
)

type HoverableRectProps struct {
	sg2.Geometry
	color        sg.Color
	hoveredColor sg.Color
}

type HoverableRectState struct {
	sg2.HoverableState
	isHovered bool
}

func HoverableRectRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
	if state.NodeState == nil {
		state.NodeState = HoverableRectState{}
		log.Printf("No node state. Created new. State is %+v", state)
		dstate := state.NodeState.(HoverableRectState)
		dstate.OnEnter = func(state sg2.StateType) {
			//dstate := *state.(*HoverableRectState)
			//dstate.isHovered = true
		}
		dstate.OnLeave = func(state sg2.StateType) {
			//dstate := *state.(*HoverableRectState)
			//dstate.isHovered = false
		}
	}
	dstate := state.NodeState.(HoverableRectState)
	dprops := props.(HoverableRectProps)
	color := dprops.color
	if dstate.isHovered {
		color = sg.Color{1, 1, 0, 0}
	}

	return sg2.RenderableNode{
		Type: sg2.RectangleNodeRender,
		Props: sg2.RectangleProps{
			dprops.Geometry,
			color,
		},
	}
}

func MainWindowRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
	return sg2.RenderableNode{
		Type: HoverableRectRender,
		Props: HoverableRectProps{
			Geometry:     sg2.Geometry{0, 0, 100, 100},
			color:        sg.Color{1, 1, 0, 0},
			hoveredColor: sg.Color{1, 1, 0, 0},
		},
		Children: []sg.Node{
			sg2.RenderableNode{
				Type: sg2.RectangleNodeRender,
				Props: sg2.RectangleProps{
					sg2.Geometry{5, 5, 10, 10},
					sg.Color{1, 0, 1, 0},
				},
			},
		},
	}
}
