package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"

	"github.com/CrimsonAS/goggle/animation"
	"github.com/CrimsonAS/goggle/animation/easing"
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
)

type TrippyState struct {
	color           sg.Color
	inverseColor    sg.Color
	containsPointer bool
	active          bool
	scaleAnimation  *animation.FloatAnimation
	colorAnimation  *animation.ColorAnimation
}

func TrippyRender(props sg.PropType, state *sg.RenderState) sg.Node {
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

	return sg.SimpleRectangleNode{
		Size:  sg.Vec2{sz.X, sz.Y},
		Color: sg.Color{1, 0, 0, 0},
		Children: []sg.Node{
			sg.InputNode{
				Geometry: sg.Geometry{0, 0, sz.X, sz.Y},
				OnEnter: func(input sg.InputState) {
					log.Printf("hoverable rect OnEnter")
					dstate.containsPointer = true
				},
				OnLeave: func(input sg.InputState) {
					log.Printf("hoverable rect OnLeave")
					dstate.containsPointer = false
				},
			},

			sg.TransformNode{
				Matrix: sg.Scale2D(dstate.scaleAnimation.Get(), dstate.scaleAnimation.Get()),
				Children: []sg.Node{
					sg.SimpleRectangleNode{
						Size:  sg.Vec2{100, 100},
						Color: sg.Color{1, 0, 1, 0},
					},
				},
			},
			/*
					sg.ScaleNode{
						Scale: dstate.scaleAnimation.Get(),
						Children: []sg.Node{
							sg.ImageNode{
								Width:  100,
								Height: 100,
								Texture: sg.FileTexture{
									Source: "solid.png",
								},
							},
						},
					},
					sg.TextNode{
						X:          float32(sz.X / 2 * dstate.scaleAnimation.Get()),
						Width:      300,
						Height:     42,
						Text:       "Hello, world",
						Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
						PixelSize:  42,
						FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
					},
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
				sg.Repeater{
					Model: 200,
					New: func(index int) sg.Node {
						findex := float32(index)
						return sg.RectangleNode{
							X:      findex * dstate.scaleAnimation.Get(),
							Y:      findex * dstate.scaleAnimation.Get(),
							Width:  findex * dstate.scaleAnimation.Get(),
							Height: findex * dstate.scaleAnimation.Get(),
							Color: sg.Color{dstate.color.A(),
								findex * dstate.color.R() * dstate.scaleAnimation.Get(),
								findex * dstate.color.G() * dstate.scaleAnimation.Get(),
								findex * dstate.color.B() * dstate.scaleAnimation.Get(),
							},
						}
					},
				},
				sg.Repeater{
					X:     sz.X,
					Model: 200,
					New: func(index int) sg.Node {
						findex := float32(index)
						return sg.RectangleNode{
							X:      -findex * dstate.scaleAnimation.Get(),
							Y:      findex * dstate.scaleAnimation.Get(),
							Width:  findex * dstate.scaleAnimation.Get(),
							Height: findex * dstate.scaleAnimation.Get(),
							Color: sg.Color{dstate.color.A(),
								findex * dstate.color.R() * dstate.scaleAnimation.Get(),
								findex * dstate.color.G() * dstate.scaleAnimation.Get(),
								findex * dstate.color.B() * dstate.scaleAnimation.Get(),
							},
						}
					},
				},
			*/
		},
	}
}

func MainWindowRender(props sg.PropType, state *sg.RenderState) sg.Node {
	return sg.RenderableNode{
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
		w.Render(MainWindowRender(nil, &sg.RenderState{Window: w}))
	}

	r.Quit()
}
