package main

import (
	"github.com/CrimsonAS/goggle/sg"
)

type Draggable struct {
	Child     sg.Geometryable
	IsHovered bool
	IsPressed bool
}

// interface assertions
var _ sg.Positionable = (*Draggable)(nil)
var _ sg.Sizeable = (*Draggable)(nil)
var _ sg.Hoverable = (*Draggable)(nil)

func (this *Draggable) Render(w sg.Windowable) sg.Node {
	return this.Child
}

func (this *Draggable) PointerEnter(pos sg.Vec2) { // ### Entered? Left?
	this.IsHovered = true
}
func (this *Draggable) PointerLeave(pos sg.Vec2) {
	this.IsHovered = false
}

func (this *Draggable) PointerPressed(pos sg.Vec2) {
	this.IsPressed = true
}

func (this *Draggable) PointerMoved(pos sg.Vec2) {
	this.SetPosition(pos)
}

func (this *Draggable) PointerReleased(pos sg.Vec2) {
	this.IsPressed = false
}

func (this *Draggable) Position() sg.Vec2 {
	return this.Child.Position()
}

func (this *Draggable) SetPosition(pos sg.Vec2) {
	this.Child.SetPosition(pos)
}

func (this *Draggable) Size() sg.Vec2 {
	return this.Child.Size()
}

func (this *Draggable) SetSize(sz sg.Vec2) {
	this.Child.SetSize(sz)
}
