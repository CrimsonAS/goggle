package main

import "github.com/CrimsonAS/goggle/sg"

type MainWindow struct {
	windowable     sg.Windowable
	draggable_rect *DraggableRect
}

func (this *MainWindow) Size() sg.Vec2 {
	if this.windowable != nil {
		return this.windowable.GetSize()
	}
	return sg.Vec2{0, 0}
}

func (this *MainWindow) SetSize(sz sg.Vec2) {

}

func (this *MainWindow) Render(w sg.Windowable) sg.Node {
	if this.draggable_rect == nil {
		this.draggable_rect = &DraggableRect{sg.RectangleNode{X: 0, Y: 0, Width: 100, Height: 100, Color: sg.Color{1, 1, 0, 0}}}
	}

	this.windowable = w
	sz := w.GetSize()

	return &sg.RectangleNode{
		Color:  sg.Color{1, 0, 1, 0},
		Width:  sz.X,
		Height: sz.Y,
		Children: []sg.Node{
			this.draggable_rect,
		},
	}
}
