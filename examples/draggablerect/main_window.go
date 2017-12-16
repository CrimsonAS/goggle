package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

type DraggableRectProps struct {
	color        sg.Color
	hoveredColor sg.Color
	pressedColor sg.Color
}

type DraggableRectState struct {
	sg.Geometry
	IsHovered bool
	IsPressed bool
}

func DraggableRectRender(props components.PropType, state *components.RenderState) sg.Node {
	dstate, _ := state.NodeState.(*DraggableRectState)
	if dstate == nil {
		dstate = &DraggableRectState{}
		dstate.Geometry = sg.Geometry{25, 25, 100, 100}
		state.NodeState = dstate
		log.Printf("No node state. Created new. State is %+v", state.NodeState)
	}

	dprops := props.(DraggableRectProps)
	color := dprops.color
	if dstate.IsPressed {
		color = dprops.pressedColor
	} else if dstate.IsHovered {
		color = dprops.hoveredColor
	}
	log.Printf("rect %s", dstate.Geometry)

	return components.Component{
		Type: components.Rectangle,
		Props: components.RectangleProps{
			Geometry: dstate.Geometry,
			Color:    color,
		},
		Children: []sg.Node{
			nodes.Input{
				Geometry: dstate.Geometry.ZeroOrigin(),
				OnEnter: func(input nodes.InputState) {
					log.Printf("rect OnEnter")
					dstate.IsHovered = true
				},
				OnLeave: func(input nodes.InputState) {
					log.Printf("rect OnLeave")
					dstate.IsHovered = false
				},
				OnPress: func(input nodes.InputState) {
					log.Printf("rect OnPressed")
					dstate.IsPressed = true
					dstate.Geometry.X = input.MousePos.X
					dstate.Geometry.Y = input.MousePos.Y
				},
				OnRelease: func(input nodes.InputState) {
					log.Printf("rect OnReleased")
					dstate.IsPressed = false
				},
				OnMove: func(input nodes.InputState) {
					log.Printf("rect OnMove %s", input.MousePos)
				},
			},
		},
	}
}

func MainWindowRender(props components.PropType, state *components.RenderState) sg.Node {
	return nodes.Rectangle{
		Size:  state.Window.GetSize(),
		Color: sg.Color{1, 0, 0, 0},
		Children: []sg.Node{
			components.Component{
				Type: DraggableRectRender,
				Props: DraggableRectProps{
					color:        sg.Color{1, 1, 0, 0},
					hoveredColor: sg.Color{1, 1, 1, 0},
					pressedColor: sg.Color{1, 1, 1, 1},
				},
			},
		},
	}
}
