package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

type MainWindow struct {
	manyRectChildren    []sg.Node // method 1
	howManyRectChildren int       // method 2
	sz                  sg.Vec2
}

func (this *MainWindow) Size() sg.Vec2 {
	return this.sz
}

func (this *MainWindow) SetSize(sz sg.Vec2) {
	// can't alter; we are top level
}

const method1 = false

func (this *MainWindow) Render(w sg.Windowable) sg.Node {
	this.sz = w.GetSize()

	const childSize = 200
	addChance := rand.Intn(99)

	if addChance > 20 && len(this.manyRectChildren) < 10000 {
		if method1 {
			for i := 0; i < 10; i++ {
				this.manyRectChildren = append(this.manyRectChildren, &sg.RectangleNode{X: 0, Y: 0, Width: childSize, Height: childSize, Color: sg.Color{1, 1, 1, 0}})
			}
		} else {
			this.howManyRectChildren += 10
		}
	}
	remChance := rand.Intn(99)
	if remChance > 90 && len(this.manyRectChildren) > 500 {
		if method1 {
			this.manyRectChildren = this.manyRectChildren[1:]
		} else {
			this.howManyRectChildren -= 1
		}
	}

	if !method1 {
		childs := []sg.Node{}
		for i := 0; i < this.howManyRectChildren; i++ {
			childs = append(childs, &sg.RectangleNode{X: 0, Y: 0, Width: childSize, Height: childSize, Color: sg.Color{1, 1, 1, 0}})
		}
		this.manyRectChildren = childs
	}

	for _, child := range this.manyRectChildren {
		rchild := child.(*sg.RectangleNode)
		rchild.X = rand.Float32() * (this.sz.X - childSize)
		rchild.Y = rand.Float32() * (this.sz.Y - childSize)
		rchild.Color.X = rand.Float32()
		rchild.Color.Y = rand.Float32()
		rchild.Color.Z = rand.Float32()
		rchild.Color.W = rand.Float32()
	}

	div := w.FrameTime() / time.Millisecond
	if div == 0 {
		div = 1
	}
	fps := math.Ceil(float64(1000 / div))
	ret := &sg.RectangleNode{
		Color:  sg.Color{1, 0, 1, 0},
		Width:  this.sz.X,
		Height: this.sz.Y,
		Children: []sg.Node{
			&sg.RectangleNode{ // just a way to get an extra child.. no actual scaling..
				Children: this.manyRectChildren,
			},
			&sg.RectangleNode{
				X:      0,
				Y:      0,
				Width:  400,
				Height: 42,
				Color:  sg.Color{1, 1, 0, 0},
				Children: []sg.Node{
					&sg.TextNode{
						X:          0,
						Y:          0,
						Width:      400,
						Height:     42,
						PixelSize:  42,
						Text:       fmt.Sprintf("%d children rendered, %g FPS", len(this.manyRectChildren), fps),
						FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
						Color:      sg.Color{1, 1, 1, 1},
					},
				},
			},
		},
	}
	return ret
}
