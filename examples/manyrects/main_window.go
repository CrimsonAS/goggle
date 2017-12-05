package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

type MainWindow struct {
	manyRectChildren    []sg.Node // method 1
	howManyRectChildren int       // method 2
	sz                  sg.Vec2
	fpsLastUpdated      time.Duration
	fpsLabel            string
}

func (this *MainWindow) Size() sg.Vec2 {
	return this.sz
}

func (this *MainWindow) SetSize(sz sg.Vec2) {
	// can't alter; we are top level
}

const method1 = true

var localRand *rand.Rand = rand.New(rand.NewSource(1234))

func (this *MainWindow) Render(w sg.Windowable) sg.Node {
	this.sz = w.GetSize()

	const childSize = 200
	const maxNodes = 10000
	const minNodes = 5000
	const nodeDebug = false

	addChance := localRand.Intn(99)

	if addChance > 20 && len(this.manyRectChildren) < maxNodes {
		addNodes := int(float64(maxNodes-len(this.manyRectChildren)) * 0.05)
		if nodeDebug {
			log.Printf("Adding %d nodes", addNodes)
		}
		if method1 {
			for i := 0; i < addNodes; i++ {
				this.manyRectChildren = append(this.manyRectChildren, &sg.RectangleNode{X: 0, Y: 0, Width: childSize, Height: childSize, Color: sg.Color{1, 1, 1, 0}})
			}
		} else {
			this.howManyRectChildren += addNodes
		}
	}
	remChance := localRand.Intn(99)
	if remChance > 90 && len(this.manyRectChildren) > minNodes {
		delNodes := int(float64(minNodes * 0.05))
		if nodeDebug {
			log.Printf("Removing %d nodes", delNodes)
		}
		if method1 {
			this.manyRectChildren = this.manyRectChildren[delNodes:]
		} else {
			this.howManyRectChildren -= delNodes
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
		rchild.X = localRand.Float32() * (this.sz.X - childSize)
		rchild.Y = localRand.Float32() * (this.sz.Y - childSize)

		const blend = false
		if blend {
			rchild.Color.X = localRand.Float32()
		}
		rchild.Color.Y = localRand.Float32()
		rchild.Color.Z = localRand.Float32()
		rchild.Color.W = localRand.Float32()
	}

	this.fpsLastUpdated += w.FrameTime()
	if this.fpsLabel == "" || this.fpsLastUpdated > 1*time.Second {
		this.fpsLastUpdated = 0
		div := w.FrameTime() / time.Millisecond
		if div == 0 {
			div = 1
		}
		fps := math.Ceil(float64(1000 / div))
		this.fpsLabel = fmt.Sprintf("%d children rendered, %g FPS", len(this.manyRectChildren), fps)
		const fpsDebug = true
		if fpsDebug {
			log.Printf(this.fpsLabel)
		}
	}
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
						Text:       this.fpsLabel,
						FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
						Color:      sg.Color{1, 1, 1, 1},
					},
				},
			},
		},
	}
	return ret
}
