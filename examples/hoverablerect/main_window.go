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

func HoverableRectRender(props sg2.PropType, state *sg2.TouchRenderState) sg.Node {
	/*	if state.NodeState == nil {
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
		dstate := state.NodeState.(HoverableRectState)*/
	if state.OnEnter == nil {
		state.OnEnter = func(*sg2.TouchState) {
			log.Printf("hoverable rect OnEnter")
		}
		state.OnLeave = func(*sg2.TouchState) {
			log.Printf("hoverable rect OnLeave")
		}
	}

	dprops := props.(HoverableRectProps)
	color := dprops.color
	if state.IsHovered {
		color = dprops.hoveredColor
	}

	state.TouchGeometry = dprops.Geometry
	return sg2.RenderableNode{
		Type: sg2.RectangleNodeRender,
		Props: sg2.RectangleProps{
			dprops.Geometry,
			color,
		},
	}
}

func MainWindowRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
	return sg2.TouchNode{
		Type: HoverableRectRender,
		Props: HoverableRectProps{
			Geometry:     sg2.Geometry{0, 0, 100, 100},
			color:        sg.Color{1, 1, 0, 0},
			hoveredColor: sg.Color{1, 1, 1, 0},
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
