package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"math/rand"
)

type OtherButton struct {
	ObjectName string
}

func (this *OtherButton) Render() sg.TreeNode {
	return &sg.Rectangle{
		ObjectName: "OtherButtonRect",
		X:          5,
		Y:          5,
		Width:      5,
		Height:     5,
		Color:      sg.Color{0.3, 0.4, 0.5, 0.6},
	}
}

type Button struct {
	ObjectName string
	color      sg.Color
}

func (this *Button) Render() sg.TreeNode {
	return &sg.Rectangle{
		ObjectName: "Rect",
		X:          100,
		Y:          100,
		Width:      200,
		Height:     200,
		Color:      this.color,
		Children: []sg.TreeNode{
			&OtherButton{
				ObjectName: "OtherButton",
			},
			&sg.Rectangle{
				ObjectName: "Rect2",
				X:          100,
				Y:          200,
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

	thing := &Button{ObjectName: "Button"}
	for {
		thing.SetColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		w.Render(thing)
		r.ProcessEvents()
	}

	r.Quit()
}
