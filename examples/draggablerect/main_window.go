package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/layouts"
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

func DraggableLayout(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
	geo := props.(sg.Geometry)
	for _, child := range children {
		child.Render(sg.FixedConstraint(geo.Size))
		child.SetPosition(geo.Origin)
	}
	return c.Max
}

func DraggableRectRender(props components.PropType, state *components.RenderState) sg.Node {
	dstate, _ := state.NodeState.(*DraggableRectState)
	if dstate == nil {
		dstate = &DraggableRectState{}
		dstate.Geometry = sg.Geometry{sg.Position{25, 25}, sg.Size{100, 100}}
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

	return layouts.Box{
		Layout: DraggableLayout,
		Props:  dstate.Geometry,
		Children: []sg.Node{
			components.Component{
				Type:  components.Rectangle,
				Props: components.RectangleProps{Color: color},
				Children: []sg.Node{
					nodes.Input{
						OnEnter: func(input nodes.InputState) {
							//log.Printf("rect OnEnter")
							dstate.IsHovered = true
						},
						OnLeave: func(input nodes.InputState) {
							//log.Printf("rect OnLeave")
							dstate.IsHovered = false
						},
						OnPress: func(input nodes.InputState) {
							//log.Printf("rect OnPressed")
							dstate.IsPressed = true
							dstate.Geometry.Origin = input.SceneMousePos.Sub(dstate.Geometry.Size.Div(2).ToPosition())
						},
						OnRelease: func(input nodes.InputState) {
							//log.Printf("rect OnReleased")
							dstate.IsPressed = false
						},
						OnMove: func(input nodes.InputState) {
							//log.Printf("rect OnMove %s", input.MousePos)
							dstate.Geometry.Origin = input.SceneMousePos.Sub(dstate.Geometry.Size.Div(2).ToPosition())
						},
					},
				},
			},
		},
	}
}

func MainWindowRender(props components.PropType, state *components.RenderState) sg.Node {
	return nodes.Rectangle{
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
