package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"math/rand"
)

type OtherButton struct {
	sg.BasicNode
}

func (this *OtherButton) Render() sg.TreeNode {
	return &sg.Rectangle{
		BasicNode: sg.BasicNode{
			ObjectName: "OtherButtonRect",
		},
		X:      110,
		Y:      110,
		Width:  180,
		Height: 180,
		Color:  sg.Color{0.5, 0.0, 1.0, 0.0},
	}
}

type Button struct {
	sg.BasicNode
	color sg.Color
}

func (this *Button) Render() sg.TreeNode {
	return &sg.Rectangle{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
		Color:  this.color,
		BasicNode: sg.BasicNode{
			ObjectName: "Rect",
			Children: []sg.TreeNode{
				&OtherButton{
					sg.BasicNode{ObjectName: "OtherButton"},
				},
				&sg.Rectangle{
					BasicNode: sg.BasicNode{ObjectName: "Rect2"},
					X:         100,
					Y:         200,
					Width:     50,
					Height:    50,
					Color:     sg.Color{0.5, 1.0, 0, 0},
				},
			},
		},
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
