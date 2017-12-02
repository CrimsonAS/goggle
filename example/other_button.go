package main

import (
	"time"

	"github.com/CrimsonAS/goggle/animation"
	"github.com/CrimsonAS/goggle/sg"
)

type OtherButton struct {
	containsPointer bool
	scaleAnimation  *animation.FloatAnimation
	x, y, w, h      float32
	currentScale    float32
}

func (this *OtherButton) Size() sg.Vec2 {
	return sg.Vec2{this.w, this.h}
}

func (this *OtherButton) SetSize(sz sg.Vec2) {
	this.w, this.h = sz.X, sz.Y
}

func (this *OtherButton) Position() sg.Vec2 {
	return sg.Vec2{this.x, this.y}
}

func (this *OtherButton) SetPosition(pos sg.Vec2) {
	this.x, this.y = pos.X, pos.Y
}

// hoverable
func (this *OtherButton) PointerEnter(tp sg.Vec2) {
	this.containsPointer = true
}

// hoverable
func (this *OtherButton) PointerLeave(tp sg.Vec2) {
	this.containsPointer = false
}

// pressable
func (this *OtherButton) PointerPressed(tp sg.Vec2) {
}

// pressable
func (this *OtherButton) PointerReleased(tp sg.Vec2) {
}

// moveable
func (this *OtherButton) PointerMoved(tp sg.Vec2) {
	this.x = tp.X
	this.y = tp.Y
}

func (this *OtherButton) Render(w sg.Windowable) sg.Node {
	if this.scaleAnimation == nil {
		this.scaleAnimation = &animation.FloatAnimation{
			From:     1.0,
			To:       5.0,
			Duration: 1000 * time.Millisecond,
		}
		this.scaleAnimation.Restart()
	}

	if this.containsPointer {
		this.scaleAnimation.Advance(w.FrameTime())
		this.currentScale = this.scaleAnimation.Get()
	} else if this.currentScale < 0.2 {
		this.currentScale = 0.2
	}
	return &sg.ScaleNode{
		Scale: 1.0,
		//Scale: this.currentScale,
		Children: []sg.Node{
			&sg.ImageNode{
				X:      this.x,
				Y:      this.y,
				Width:  this.w,
				Height: this.h,
				Texture: &sg.FileTexture{
					Source: "solid.png",
				},
			},
		},
	}
}
