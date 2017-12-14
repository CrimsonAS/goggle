package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
)

type HoverableRectProps struct {
	sg.Geometry
	color        sg.Color
	hoveredColor sg.Color
}

type HoverableRectState struct {
	IsHovered bool
}

func HoverableRectRender(props sg.PropType, state *sg.RenderState) sg.Node {
	dstate, _ := state.NodeState.(*HoverableRectState)
	if dstate == nil {
		dstate = &HoverableRectState{}
		state.NodeState = dstate
		log.Printf("No node state. Created new. State is %+v", state.NodeState)
	}

	dprops := props.(HoverableRectProps)
	color := dprops.color
	if dstate.IsHovered {
		color = dprops.hoveredColor
	}

	return sg.RenderableNode{
		Type: sg.RectangleNodeRender,
		Props: sg.RectangleProps{
			dprops.Geometry,
			color,
		},
		Children: []sg.Node{
			sg.InputNode{
				Geometry: dprops.Geometry.ZeroOrigin(),
				OnEnter: func(input sg.InputState) {
					log.Printf("hoverable rect OnEnter")
					dstate.IsHovered = true
				},
				OnLeave: func(input sg.InputState) {
					log.Printf("hoverable rect OnLeave")
					dstate.IsHovered = false
				},
				OnPress: func(input sg.InputState) {
					log.Printf("hoverable rect OnPressed")
				},
				OnRelease: func(input sg.InputState) {
					log.Printf("hoverable rect OnReleased")
				},
				OnMove: func(input sg.InputState) {
					log.Printf("hoverable rect OnMove")
				},
			},
		},
	}
}

func MainWindowRender(props sg.PropType, state *sg.RenderState) sg.Node {
	return sg.RenderableNode{
		Type: HoverableRectRender,
		Props: HoverableRectProps{
			Geometry:     sg.Geometry{25, 25, 100, 100},
			color:        sg.Color{1, 1, 0, 0},
			hoveredColor: sg.Color{1, 1, 1, 0},
		},
		Children: []sg.Node{
			sg.RenderableNode{
				Type: sg.RectangleNodeRender,
				Props: sg.RectangleProps{
					sg.Geometry{5, 5, 10, 10},
					sg.Color{1, 0, 1, 0},
				},
			},
		},
	}
}
