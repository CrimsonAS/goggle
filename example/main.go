package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
)

type OtherButton struct{}

func (this *OtherButton) Render() sg.Node {
	return &sg.Image{
		X:      10,
		Y:      10,
		Width:  180,
		Height: 180,
		Texture: &sg.FileTexture{
			Source: "solid.png",
		},
	}
}

type Button struct {
	color   sg.Color
	bstep   int
	bToLeft bool
}

func (this *Button) Render() sg.Node {
	if this.bToLeft {
		this.bstep -= 1
		if this.bstep < 0 {
			this.bstep = 0
			this.bToLeft = false
		}
	} else {
		this.bstep += 1
		if this.bstep > 150 {
			/* our width - inner rect's width */
			this.bstep = 150
			this.bToLeft = true
		}
	}
	this.color = sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()}

	width := float32(200.0)
	height := float32(200.0)

	return &sg.Rectangle{
		Width:  width,
		Height: height,
		Children: []sg.Node{
			&OtherButton{},
			&sdlsoftware.DrawNode{
				Draw: func(renderer *sdl.Renderer, node *sdlsoftware.DrawNode) {
					// custom drawing here
				},
			},
			&sg.Rectangle{
				X:      float32(this.bstep),
				Y:      100,
				Width:  50,
				Height: 50,
				Color:  sg.Color{0.5, 1.0, 0, 0},
			},
			&sg.Text{
				X:          0,
				Y:          0,
				Width:      width,
				Height:     42,
				Text:       "Hello, world",
				Color:      sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()},
				PixelSize:  12,
				FontFamily: "Barlow/Barlow-Regular.ttf",
			},
		},
		Color: this.color,
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
