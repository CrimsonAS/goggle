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

type OtherButton struct{}

func (this *OtherButton) Render() sg.Node {
	return &sg.Image{
		X:        10,
		Y:        10,
		Width:    180,
		Height:   180,
		Scale:    1.0,
		Rotation: 1.0,
		Texture: &sg.FileTexture{
			Source: "solid.png",
		},
	}
}

type Button struct {
	color           sg.Color
	containsPointer bool
	active          bool
	rectAnimation   *animation.FloatAnimation
	bgAnimation     *animation.FloatAnimation
	scaleAnimation  *animation.FloatAnimation
}

func (this *Button) Geometry() (x, y, w, h float32) {
	if this.rectAnimation == nil {
		this.rectAnimation = &animation.FloatAnimation{
			From:     0,
			To:       150,
			Duration: 1000 * time.Millisecond,
		}
		this.bgAnimation = &animation.FloatAnimation{
			From:     200,
			To:       1000,
			Duration: 2000 * time.Millisecond,
		}
		this.scaleAnimation = &animation.FloatAnimation{
			From:     0.0,
			To:       10.0,
			Duration: 5000 * time.Millisecond,
		}
	}
	return 0, 0, 200, this.bgAnimation.Get()
}

func (this *Button) SetGeometry(x, y, w, h float32) { // why does GeometryNode require this?

}

// hoverable
func (this *Button) PointerEnter(tp sg.TouchPoint) {
	this.containsPointer = true
	log.Printf("Pointer entered: %+v", this)
}

// hoverable
func (this *Button) PointerLeave(tp sg.TouchPoint) {
	this.containsPointer = false
	log.Printf("Pointer left: %+v", this)
}

// tapable
func (this *Button) PointerTapped(tp sg.TouchPoint) {
	this.active = !this.active
	log.Printf("Pointer tapped: %+v", this)
}

// touchable
func (this *Button) PointerPressed(tp sg.TouchPoint) {
	log.Printf("Pointer touched: %+v", this)
}

// touchable
func (this *Button) PointerReleased(tp sg.TouchPoint) {
	log.Printf("Pointer released: %+v", this)
}

// moveable
func (this *Button) PointerMoved(tp sg.TouchPoint) {
	log.Printf("Pointer moved: %+v", this)
}

func (this *Button) Render() sg.Node {
	if this.active {
		this.color = sg.Color{1, 0, 0, 1}
	} else {
		if this.containsPointer {
			this.color = sg.Color{1, 0, 1, 0}
		} else {
			this.color = sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()}
		}
	}

	_, _, width, height := this.Geometry()

	return &sg.Rectangle{
		Color:    this.color,
		Width:    width,
		Height:   height,
		Scale:    1.0,
		Rotation: 1.0,
		Children: []sg.Node{
			&OtherButton{},
			&sdlsoftware.DrawNode{
				Draw: func(renderer *sdl.Renderer, node *sdlsoftware.DrawNode) {
					// custom drawing here
				},
			},
			&sg.Rectangle{
				X:        float32(this.rectAnimation.Get()),
				Y:        height / 2,
				Width:    50,
				Height:   50,
				Color:    sg.Color{0.5, 1.0, 0, 0},
				Scale:    this.scaleAnimation.Get(),
				Rotation: 1.0,
			},
			&sg.Text{
				X:          float32(width - this.rectAnimation.Get()),
				Width:      width,
				Height:     42,
				Text:       "Hello, world",
				Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
				PixelSize:  42,
				FontFamily: "Barlow/Barlow-Regular.ttf",
				Scale:      1.0,
				Rotation:   1.0,
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
