package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

type HoverableRectProps struct {
	sg.Geometry
	color        sg.Color
	hoveredColor sg.Color
}

type HoverableRectState struct {
	IsHovered bool
}

func HoverableRectRender(props components.PropType, state *components.RenderState) sg.Node {
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

	return components.Component{
		Type: components.Rectangle,
		Props: components.RectangleProps{
			dprops.Geometry,
			color,
		},
		Children: []sg.Node{
			nodes.Input{
				Geometry: dprops.Geometry.ZeroOrigin(),
				OnEnter: func(input nodes.InputState) {
					log.Printf("hoverable rect OnEnter")
					dstate.IsHovered = true
				},
				OnLeave: func(input nodes.InputState) {
					log.Printf("hoverable rect OnLeave")
					dstate.IsHovered = false
				},
				OnPress: func(input nodes.InputState) {
					log.Printf("hoverable rect OnPressed")
				},
				OnRelease: func(input nodes.InputState) {
					log.Printf("hoverable rect OnReleased")
				},
				OnMove: func(input nodes.InputState) {
					log.Printf("hoverable rect OnMove")
				},
			},
		},
	}
}

func MainWindowRender(props components.PropType, state *components.RenderState) sg.Node {
	return components.Component{
		Type: HoverableRectRender,
		Props: HoverableRectProps{
			Geometry:     sg.Geometry{25, 25, 100, 100},
			color:        sg.Color{1, 1, 0, 0},
			hoveredColor: sg.Color{1, 1, 1, 0},
		},
		Children: []sg.Node{
			components.Component{
				Type: components.Rectangle,
				Props: components.RectangleProps{
					sg.Geometry{5, 5, 10, 10},
					sg.Color{1, 0, 1, 0},
				},
			},
		},
	}
}
