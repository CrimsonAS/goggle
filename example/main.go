package main

import (
	"fmt"
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type OtherButton struct {
	sg.BasicNode
}

func (this *OtherButton) Render() sg.TreeNode {
	return &sg.Rectangle{
		BasicNode: sg.BasicNode{
			ObjectName: "OtherButtonRect",
			X:          10,
			Y:          10,
			Width:      180,
			Height:     180,
		},
		Color: sg.Color{0.5, 0.0, 1.0, 0.0},
	}
}

type Button struct {
	sg.BasicNode
	color sg.Color
}

func (this *Button) Render() sg.TreeNode {
	return &sg.Rectangle{
		BasicNode: sg.BasicNode{
			ObjectName: "Rect",
			X:          100,
			Y:          100,
			Width:      200,
			Height:     200,
			Children: []sg.TreeNode{
				&OtherButton{
					sg.BasicNode{ObjectName: "OtherButton"},
				},
				&sdlsoftware.SDLDrawNode{
					Draw: func(surface *sdl.Surface, node *sdlsoftware.SDLDrawNode) {
						fmt.Printf("custom drawing here\n")
					},
				},
				&sg.Rectangle{
					BasicNode: sg.BasicNode{
						ObjectName: "Rect2",
						X:          0,
						Y:          100,
						Width:      50,
						Height:     50,
					},
					Color: sg.Color{0.5, 1.0, 0, 0},
				},
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

	thing := &Button{BasicNode: sg.BasicNode{ObjectName: "Button"}}
	for {
		thing.SetColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		w.Render(thing)
		r.ProcessEvents()
	}

	r.Quit()
}
