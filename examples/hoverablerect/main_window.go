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
	IsHovered bool
}

func HoverableRectRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
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

	return sg2.RenderableNode{
		Type: sg2.RectangleNodeRender,
		Props: sg2.RectangleProps{
			dprops.Geometry,
			color,
		},
		Children: []sg.Node{
			sg2.InputNode{
				Geometry: dprops.Geometry.ZeroOrigin(),
				OnEnter: func(input sg2.InputState) {
					log.Printf("hoverable rect OnEnter")
					dstate.IsHovered = true
				},
				OnLeave: func(input sg2.InputState) {
					log.Printf("hoverable rect OnLeave")
					dstate.IsHovered = false
				},
				OnPress: func(input sg2.InputState) {
					log.Printf("hoverable rect OnPressed")
				},
				OnRelease: func(input sg2.InputState) {
					log.Printf("hoverable rect OnReleased")
				},
				OnMove: func(input sg2.InputState) {
					log.Printf("hoverable rect OnMove")
				},
			},
		},
	}
}

func MainWindowRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
	return sg2.RenderableNode{
		Type: HoverableRectRender,
		Props: HoverableRectProps{
			Geometry:     sg2.Geometry{25, 25, 100, 100},
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
