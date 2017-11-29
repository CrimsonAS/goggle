package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
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
	return &sg.Rectangle{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
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
		},
		Color: this.color,
	}
}

func (this *Button) SetColor(col sg.Color) {
	this.color = col
}

func main() {
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
		thing.SetColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		w.Render(thing)
		r.ProcessEvents()
	}

	r.Quit()
}
