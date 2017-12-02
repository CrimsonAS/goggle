package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"

	"github.com/CrimsonAS/goggle/animation"
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	color           sg.Color
	containsPointer bool
	active          bool
	rectAnimation   *animation.FloatAnimation
	bgAnimation     *animation.FloatAnimation
	scaleAnimation  *animation.FloatAnimation
	otherButton     *OtherButton
}

func (this *Button) Size() (w, h float32) {
	if this.rectAnimation == nil {
		this.rectAnimation = &animation.FloatAnimation{
			From:     0,
			To:       150,
			Duration: 1000 * time.Millisecond,
		}
		this.rectAnimation.Restart()
		this.bgAnimation = &animation.FloatAnimation{
			From:     200,
			To:       1000,
			Duration: 2000 * time.Millisecond,
		}
		this.bgAnimation.Restart()
		this.scaleAnimation = &animation.FloatAnimation{
			From:     0.0,
			To:       10.0,
			Duration: 5000 * time.Millisecond,
		}
		this.scaleAnimation.Restart()
		this.otherButton = &OtherButton{w: 100, h: 100}
	}
	return 200, this.bgAnimation.Get()
}

func (this *Button) SetSize(w, h float32) { // why does Sizeable require this?

}

// hoverable
func (this *Button) PointerEnter(tp sg.TouchPoint) {
	this.containsPointer = true
}

// hoverable
func (this *Button) PointerLeave(tp sg.TouchPoint) {
	this.containsPointer = false
}

// tapable
func (this *Button) PointerTapped(tp sg.TouchPoint) {
	this.active = !this.active
}

// touchable
func (this *Button) PointerPressed(tp sg.TouchPoint) {
}

// touchable
func (this *Button) PointerReleased(tp sg.TouchPoint) {
}

// moveable
func (this *Button) PointerMoved(tp sg.TouchPoint) {
}

func (this *Button) Render(w sg.Windowable) sg.Node {
	this.rectAnimation.Advance(w.FrameTime())
	this.scaleAnimation.Advance(w.FrameTime())
	this.bgAnimation.Advance(w.FrameTime())

	if this.active {
		this.color = sg.Color{1, 0, 0, 1}
	} else {
		if this.containsPointer {
			this.color = sg.Color{1, 0, 1, 0}
		} else {
			this.color = sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()}
		}
	}

	width, height := this.Size()

	return &sg.RectangleNode{
		Color:  this.color,
		Width:  width,
		Height: height,
		Children: []sg.Node{
			this.otherButton,
			&sdlsoftware.DrawNode{
				Draw: func(renderer *sdl.Renderer, node *sdlsoftware.DrawNode) {
					// custom drawing here
				},
			},
			&sg.ScaleNode{
				Scale: this.scaleAnimation.Get(),
				Children: []sg.Node{
					&sg.RectangleNode{
						X:      float32(this.rectAnimation.Get()),
						Y:      height / 2,
						Width:  50,
						Height: 50,
						Color:  sg.Color{0.5, 1.0, 0, 0},
					},
				},
			},
			&sg.TextNode{
				X:          float32(width - this.rectAnimation.Get()),
				Width:      width,
				Height:     42,
				Text:       "Hello, world",
				Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
				PixelSize:  42,
				FontFamily: "Barlow/Barlow-Regular.ttf",
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
