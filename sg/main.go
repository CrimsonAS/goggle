package main

import (
	"math/rand"
)

type OtherButton struct{}

func (this *OtherButton) Render() TreeNode {
	return Rectangle{
		X:      5,
		Y:      5,
		Width:  5,
		Height: 5,
		Color:  Color{0.3, 0.4, 0.5, 0.6},
	}
}

type Button struct {
	color       Color
	secondColor Color
	scene       *Scene
}

func (this *Button) Render() TreeNode {
	return Rectangle{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
		Color:  this.color,
		Children: []TreeNode{
			OtherButton{},
			Rectangle{
				X:      5,
				Y:      5,
				Width:  5,
				Height: 5,
				Color:  this.secondColor,
			},
		},
	}
}

func (this *Button) SetColor(col Color) {
	this.color = col
	this.scene.MarkDirty(this)
}
func (this *Button) SetSecondColor(col Color) {
	this.secondColor = col
	this.scene.MarkDirty(this)
}

func main() {
	s := &Scene{}
	thing := &Button{scene: s}
	thing.SetSecondColor(Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
	for {
		thing.SetColor(Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()})
		s.Sync()
	}
}
