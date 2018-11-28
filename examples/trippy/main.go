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
	"math/rand"
	"os"
	"runtime/trace"
	"time"

	"github.com/CrimsonAS/goggle/animation"
	"github.com/CrimsonAS/goggle/animation/easing"
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/layouts"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

type TrippyState struct {
	color           sg.Color
	inverseColor    sg.Color
	containsPointer bool
	active          bool
	scaleAnimation  *animation.FloatAnimation
	colorAnimation  *animation.ColorAnimation
	solidTexture    nodes.Texture
}

func TrippyRender(props components.PropType, state *components.RenderState) sg.Node {
	dstate, _ := state.NodeState.(*TrippyState)
	if dstate == nil {
		dstate = &TrippyState{}
		state.NodeState = dstate
		dstate.scaleAnimation = &animation.FloatAnimation{
			From:     0.0,
			To:       2.0,
			Duration: 3000 * time.Millisecond,
			Easing:   easing.InOutCubic,
		}
		dstate.scaleAnimation.Restart()
		dstate.colorAnimation = &animation.ColorAnimation{
			From:     sg.Color{1, 1, 0, 0},
			To:       sg.Color{1, 0, 1, 0},
			Duration: 5000 * time.Millisecond,
		}
		dstate.colorAnimation.Restart()
		log.Printf("Created %+v", dstate)

		dstate.solidTexture = nodes.FileTexture("solid.png")
	}
	dstate.scaleAnimation.Advance(state.Window.FrameTime())

	if dstate.active {
		dstate.color = sg.Color{1, 0, 0, 1}
	} else {
		if dstate.containsPointer {
			dstate.color = sg.Color{1, 0, 1, 0}
		} else {
			dstate.colorAnimation.Advance(state.Window.FrameTime())
			dstate.color = dstate.colorAnimation.Get()
		}
	}

	dstate.inverseColor = sg.Color{dstate.color.X, 1.0 - dstate.color.Y, 1.0 - dstate.color.Z, 1.0 - dstate.color.W}

	sz := state.Window.GetSize()

	// Two of these will be included in the node tree below
	animatedRectStack := layouts.Box{
		Layout: func(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
			factor := props.(float32)
			for i, child := range children {
				v := float32(i) * factor
				child.Render(c.BoundedConstraints(sg.FixedConstraint(sg.Size{v, v})))
				child.SetPosition(sg.Position{v, v})
			}
			max := float32(len(children)-1) * factor * 2
			return sg.Size{max, max}
		},
		Props: dstate.scaleAnimation.Get(),
		Child: components.Component{
			Type: components.Repeater,
			Props: components.RepeaterProps{
				Model: 200,
				New: func(index int) sg.Node {
					findex := float32(index)
					return components.Component{
						Type: components.Rectangle,
						Props: components.RectangleProps{
							Color: sg.Color{dstate.color.A(),
								findex * dstate.color.R() * dstate.scaleAnimation.Get(),
								findex * dstate.color.G() * dstate.scaleAnimation.Get(),
								findex * dstate.color.B() * dstate.scaleAnimation.Get(),
							},
						},
					}
				},
			},
		},
	}

	return nodes.Rectangle{
		Color: sg.Color{1, 0, 0, 0},
		Children: []sg.Node{
			nodes.Input{
				OnEnter: func(input nodes.InputState) {
					log.Printf("hoverable rect OnEnter")
					dstate.containsPointer = true
				},
				OnLeave: func(input nodes.InputState) {
					log.Printf("hoverable rect OnLeave")
					dstate.containsPointer = false
				},
			},

			layouts.Box{
				Layout:    layouts.None,
				Transform: sg.Scale2D(dstate.scaleAnimation.Get(), dstate.scaleAnimation.Get()),
				Children: []sg.Node{
					components.Component{
						Type: components.Rectangle,
						Props: components.RectangleProps{
							Size:  sg.Size{200, 200},
							Color: sg.Color{1, 0, 1, 0},
						},
					},
					layouts.Box{
						Layout: layouts.Fixed,
						Props:  sg.Geometry{Size: sg.Size{100, 100}},
						Child: nodes.Image{
							Texture: dstate.solidTexture,
						},
					},
				},
			},
			animatedRectStack,
			layouts.Box{
				Layout:    layouts.None,
				Transform: sg.Translate2D(sz.Width, 0).MulM4(sg.Scale(-1, 1, 1)),
				Child:     animatedRectStack,
			},
			layouts.Box{
				Layout:    layouts.None,
				Transform: sg.Translate2D(sz.Width/2*dstate.scaleAnimation.Get(), 0),
				Children: []sg.Node{
					nodes.Text{
						Size:       sg.Vec2{300, 42},
						Text:       "Hello, world",
						Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
						PixelSize:  42,
						FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
					},
				},
			},
			/*
				sg.RectangleNode{
					X:      10,
					Y:      200,
					Width:  200,
					Height: 50,
					Color:  sg.Color{0, 0, 0, 0},
					Children: []sg.Node{
						sg.ScaleNode{
							Scale: dstate.scaleAnimation.Get(),
							Children: []sg.Node{
								sg.Row{
									Children: []sg.Node{
										sg.RectangleNode{
											Width:  50 * dstate.scaleAnimation.Get(),
											Height: 50 * dstate.scaleAnimation.Get(),
											Color:  dstate.color,
										},
										sg.RectangleNode{
											Width:  50 / 2 * dstate.scaleAnimation.Get() * dstate.scaleAnimation.Get(),
											Height: 50 / 2 * dstate.scaleAnimation.Get() * dstate.scaleAnimation.Get(),
											Color:  dstate.inverseColor,
										},
									},
								},
							},
						},
					},
				},
			*/
		},
	}
}

func MainWindowRender(props components.PropType, state *components.RenderState) sg.Node {
	return components.Component{
		Type: TrippyRender,
	}
}

func main() {
	const shouldTrace = false
	if shouldTrace {
		traceFile, err := os.OpenFile("traceFile.out", os.O_RDWR|os.O_CREATE, 0600)
		traceFile.Truncate(0)
		if err != nil {
			log.Println("Can't trace: %s", err.Error())
		} else {
			trace.Start(traceFile)
			defer func() {
				trace.Stop()
				traceFile.Close()
			}()
		}
	}

	r, err := sdlsoftware.NewRenderer()
	if err != nil {
		panic(err)
	}

	w, err := r.CreateWindow()
	defer w.Destroy()

	if err != nil {
		panic(err)
	}

	for r.IsRunning() {
		r.ProcessEvents()
		// ### I do not like user code calling render functions at all. Avoid.
		w.Render(MainWindowRender(nil, &components.RenderState{Window: w}))
	}

	r.Quit()
}
