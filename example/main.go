package main

import (
	"github.com/CrimsonAS/goggle/sg"
	"math/rand"
)

type OtherButton struct{}

func (this *OtherButton) Render() sg.TreeNode {
	return sg.Rectangle{
		X:      5,
		Y:      5,
		Width:  5,
		Height: 5,
		Color:  sg.Color{0.3, 0.4, 0.5, 0.6},
	}
}

type Button struct {
	color       sg.Color
	secondColor sg.Color
	scene       *sg.Scene
}

func (this *Button) Render() sg.TreeNode {
	return sg.Rectangle{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
		Color:  this.color,
		Children: []sg.TreeNode{
			OtherButton{},
			sg.Rectangle{
				X:      5,
				Y:      5,
				Width:  5,
				Height: 5,
				Color:  this.secondColor,
			},
		},
	}
}

func (this *Button) SetColor(col sg.Color) {
	this.color = col
	this.scene.MarkDirty(this)
}
func (this *Button) SetSecondColor(col sg.Color) {
	this.secondColor = col
	this.scene.MarkDirty(this)
}

func main() {
	s := &sg.Scene{}
	thing := &Button{scene: s}
	thing.SetSecondColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
	for {
		thing.SetColor(sg.Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		s.Sync()
	}
}
