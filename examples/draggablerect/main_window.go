package main

import "github.com/CrimsonAS/goggle/sg"

type MainWindow struct {
	windowable     sg.Windowable
	draggable_rect *Draggable
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
		this.draggable_rect = &Draggable{Child: &sg.RectangleNode{X: 0, Y: 0, Width: 100, Height: 100}}
	}

	if this.draggable_rect.IsPressed {
		this.draggable_rect.Child.(*sg.RectangleNode).Color = sg.Color{1, 1, 0, 0}
	} else if this.draggable_rect.IsHovered {
		this.draggable_rect.Child.(*sg.RectangleNode).Color = sg.Color{1, 0, 1, 0}
	} else {
		this.draggable_rect.Child.(*sg.RectangleNode).Color = sg.Color{1, 0, 0, 1}
	}

	this.windowable = w
	sz := w.GetSize()

	return &sg.RectangleNode{
		Color:  sg.Color{1, 0, 0, 0},
		Width:  sz.X,
		Height: sz.Y,
		Children: []sg.Node{
			this.draggable_rect,
		},
	}
}
