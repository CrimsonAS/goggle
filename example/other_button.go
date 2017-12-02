package main

import (
	"time"

	"github.com/CrimsonAS/goggle/animation"
	"github.com/CrimsonAS/goggle/sg"
)

type OtherButton struct {
	containsPointer bool
	scaleAnimation  *animation.FloatAnimation
	w, h            float32
}

func (this *OtherButton) Size() (w, h float32) {
	return this.w, this.h
}

func (this *OtherButton) SetSize(w, h float32) {
	this.w, this.h = w, h
}

// hoverable
func (this *OtherButton) PointerEnter(tp sg.TouchPoint) {
	this.containsPointer = true
}

// hoverable
func (this *OtherButton) PointerLeave(tp sg.TouchPoint) {
	this.containsPointer = false
}

func (this *OtherButton) Render() sg.Node {
	if this.scaleAnimation == nil {
		this.scaleAnimation = &animation.FloatAnimation{
			From:     1.0,
			To:       5.0,
			Duration: 1000 * time.Millisecond,
		}
	}
	scale := float32(1.0)
	if this.containsPointer {
		scale = this.scaleAnimation.Get()
	}
	return &sg.ScaleNode{
		Scale: scale,
		Children: []sg.Node{
			&sg.ImageNode{
				X:      10,
				Y:      10,
				Width:  this.w,
				Height: this.h,
				Texture: &sg.FileTexture{
					Source: "solid.png",
				},
			},
		},
	}
}
