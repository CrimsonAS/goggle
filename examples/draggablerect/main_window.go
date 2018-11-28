/*
 * Copyright 2017 Crimson AS <info@crimson.no>
 * Author: Robin Burchell <robin.burchell@crimson.no>
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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
				Type: components.Rectangle, Props: components.RectangleProps{Color: color},
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
