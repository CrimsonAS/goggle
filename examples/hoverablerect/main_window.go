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
	log.Printf("HoverableRectRender, Renderable for Rectangle is %+v", sg2.RectangleNodeRender)
	/*
		log.Printf("Entry state is %+v", state)
		if *state == nil {
			*state = &HoverableRectState{}
			log.Printf("Created, state is %+v", state)
			dstate := state.(*HoverableRectState)
			dstate.OnEnter = func(state *sg2.StateType) {
				dstate := *state.(*HoverableRectState)
				dstate.isHovered = true
			}
			dstate.OnLeave = func(state *sg2.StateType) {
				dstate := *state.(*HoverableRectState)
				dstate.isHovered = false
			}
		} else {
			log.Printf("No need to create state is %+v", state)
		}
		log.Printf("Post creation state is %+v", state)
		dstate := state.(*HoverableRectState)
		color := dprops.color
		if dstate.isHovered {
			color = sg.Color{1, 1, 0, 0}
		}*/
	dprops := props.(HoverableRectProps)
	color := dprops.color

	return sg2.RenderableNode{
		Type: sg2.RectangleNodeRender,
		Props: sg2.RectangleProps{
			dprops.Geometry,
			color,
		},
	}
}

func MainWindowRender(props sg2.PropType, state *sg2.RenderState) sg.Node {
	log.Printf("MainWindowRender, Renderable for Hoverable is %+v", HoverableRectRender)
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
