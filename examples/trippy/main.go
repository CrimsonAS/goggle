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
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	color           sg.Color
	inverseColor    sg.Color
	containsPointer bool
	active          bool
	scaleAnimation  *animation.FloatAnimation
	colorAnimation  *animation.ColorAnimation
	otherButton     *OtherButton
	windowable      sg.Windowable
}

func (this *Button) Size() sg.Vec2 {
	if this.windowable != nil {
		return this.windowable.GetSize()
	}
	return sg.Vec2{0, 0}
}

func (this *Button) SetSize(sz sg.Vec2) {

}

// hoverable
func (this *Button) PointerEnter(tp sg.Vec2) {
	this.containsPointer = true
}

// hoverable
func (this *Button) PointerLeave(tp sg.Vec2) {
	this.containsPointer = false
}

func (this *Button) Render(w sg.Windowable) sg.Node {
	this.windowable = w
	if this.scaleAnimation == nil {
		this.scaleAnimation = &animation.FloatAnimation{
			From:     0.0,
			To:       2.0,
			Duration: 3000 * time.Millisecond,
			Easing:   easing.EaseInOutCubic,
		}
		this.scaleAnimation.Restart()
		this.colorAnimation = &animation.ColorAnimation{
			From:     sg.Color{1, 1, 0, 0},
			To:       sg.Color{1, 0, 1, 0},
			Duration: 5000 * time.Millisecond,
		}
		this.colorAnimation.Restart()
		this.otherButton = &OtherButton{w: 100, h: 100}
	}
	this.scaleAnimation.Advance(w.FrameTime())

	if this.active {
		this.color = sg.Color{1, 0, 0, 1}
	} else {
		if this.containsPointer {
			this.color = sg.Color{1, 0, 1, 0}
		} else {
			this.colorAnimation.Advance(w.FrameTime())
			this.color = this.colorAnimation.Get()
		}
	}

	this.inverseColor = sg.Color{this.color.X, 1.0 - this.color.Y, 1.0 - this.color.Z, 1.0 - this.color.W}

	sz := w.GetSize()

	return &sg.RectangleNode{
		Width:  sz.X,
		Height: sz.Y,
		Color:  sg.Color{1, 0, 0, 0},
		Children: []sg.Node{
			this.otherButton,
			&sdlsoftware.DrawNode{
				Draw: func(renderer *sdl.Renderer, node *sdlsoftware.DrawNode, transform sg.Transform) {
					// custom drawing here
				},
			},
			&sg.TextNode{
				X:          float32(sz.X / 2 * this.scaleAnimation.Get()),
				Width:      300,
				Height:     42,
				Text:       "Hello, world",
				Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
				PixelSize:  42,
				FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
			},
			&sg.RectangleNode{
				X:      10,
				Y:      200,
				Width:  200,
				Height: 50,
				Color:  sg.Color{0, 0, 0, 0},
				Children: []sg.Node{
					&sg.ScaleNode{
						Scale: this.scaleAnimation.Get(),
						Children: []sg.Node{
							&sg.Row{
								Children: []sg.Node{
									&sg.RectangleNode{
										Width:  50 * this.scaleAnimation.Get(),
										Height: 50 * this.scaleAnimation.Get(),
										Color:  this.color,
									},
									&sg.RectangleNode{
										Width:  50 / 2 * this.scaleAnimation.Get() * this.scaleAnimation.Get(),
										Height: 50 / 2 * this.scaleAnimation.Get() * this.scaleAnimation.Get(),
										Color:  this.inverseColor,
									},
								},
							},
						},
					},
				},
			},
		},
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

	thing := &Button{}

	for r.IsRunning() {
		r.ProcessEvents()
		w.Render(thing)
	}

	r.Quit()
}
