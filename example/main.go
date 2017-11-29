package main

import (
	"fmt"
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
	color sg.Color
}

func (this *Button) Render() sg.Node {
	return &sg.Rectangle{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
		Children: []sg.Node{
			&OtherButton{},
			&sdlsoftware.DrawNode{
				Draw: func(renderer *sdl.Renderer, node *sdlsoftware.DrawNode) {
					fmt.Printf("custom drawing here\n")
				},
			},
			&sg.Rectangle{
				X:      0,
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
	for {
		thing.SetColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		w.Render(thing)
		r.ProcessEvents()
	}

	r.Quit()
}
