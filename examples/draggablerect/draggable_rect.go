package main

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
)

type DraggableRect struct {
	sg.RectangleNode
}

// interface assertions
var _ sg.Positionable = (*DraggableRect)(nil)
var _ sg.Sizeable = (*DraggableRect)(nil)
var _ sg.Hoverable = (*DraggableRect)(nil)

func (this *DraggableRect) PointerEnter(pos sg.Vec2) { // ### Entered? Left?
	this.Color = sg.Color{1, 0.5, 0.5, 1}
}
func (this *DraggableRect) PointerLeave(pos sg.Vec2) {
	this.Color = sg.Color{1, 1, 0, 0}
}

func (this *DraggableRect) PointerPressed(pos sg.Vec2) {
	this.Color = sg.Color{1, 0.8, 0.8, 1}
}

func (this *DraggableRect) PointerMoved(pos sg.Vec2) {
	log.Printf("Moving to %s", pos)
	this.X, this.Y = pos.X, pos.Y
}

func (this *DraggableRect) PointerReleased(pos sg.Vec2) {
	this.Color = sg.Color{1, 0.5, 0.5, 1}
}
