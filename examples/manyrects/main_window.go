package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

/*
type MainWindow struct {
	manyRectChildren    []sg.Node // method 1
	howManyRectChildren int       // method 2
	sz                  sg.Vec2
	fpsLastUpdated      time.Duration
	fpsLabel            string
}
*/

// these really belong in state perhaps
var manyRectChildren []sg.Node // method 1
var howManyRectChildren int    // method 2
var fpsLastUpdated time.Duration
var fpsLabel string

const method1 = true

var localRand *rand.Rand = rand.New(rand.NewSource(1234))

func ManyRectRender(props components.PropType, state *components.RenderState) sg.Node {
	//func (this *MainWindow) Render(w sg.Windowable) sg.Node {
	sz := state.Window.GetSize()
	frameTime := state.Window.FrameTime()

	const childSize = 200
	const maxNodes = 10000
	const minNodes = 5000
	const nodeDebug = false

	addChance := localRand.Intn(99)

	if addChance > 20 && len(manyRectChildren) < maxNodes {
		addNodes := int(float64(maxNodes-len(manyRectChildren)) * 0.05)
		if nodeDebug {
			log.Printf("Adding %d nodes", addNodes)
		}
		if method1 {
			for i := 0; i < addNodes; i++ {
				manyRectChildren = append(manyRectChildren, nodes.Transform{
					Matrix: sg.Translate2D(0, 0),
					Children: []sg.Node{
						nodes.Rectangle{Size: sg.Vec2{childSize, childSize}, Color: sg.Color{1, 1, 1, 0}},
					},
				})
			}
		} else {
			howManyRectChildren += addNodes
		}
	}
	remChance := localRand.Intn(99)
	if remChance > 90 && len(manyRectChildren) > minNodes {
		delNodes := int(float64(minNodes * 0.05))
		if nodeDebug {
			log.Printf("Removing %d nodes", delNodes)
		}
		if method1 {
			manyRectChildren = manyRectChildren[delNodes:]
		} else {
			howManyRectChildren -= delNodes
		}
	}

	if !method1 {
		childs := []sg.Node{}
		for i := 0; i < howManyRectChildren; i++ {
			childs = append(manyRectChildren, nodes.Transform{
				Matrix: sg.Translate2D(0, 0),
				Children: []sg.Node{
					nodes.Rectangle{Size: sg.Vec2{childSize, childSize}, Color: sg.Color{1, 1, 1, 0}},
				},
			})
		}
		manyRectChildren = childs
	}

	for idx, child := range manyRectChildren {
		tchild := child.(nodes.Transform)
		//tchild.Matrix = sg.Translate2D(float32(idx), float32(idx))
		tchild.Matrix = sg.Translate2D(localRand.Float32()*(sz.X-childSize), localRand.Float32()*(sz.Y-childSize))

		rchild := tchild.Children[0].(nodes.Rectangle)
		const blend = false
		if blend {
			rchild.Color.X = localRand.Float32()
		}
		rchild.Color.Y = localRand.Float32()
		rchild.Color.Z = localRand.Float32()
		rchild.Color.W = localRand.Float32()
		tchild.Children[0] = rchild
		manyRectChildren[idx] = tchild
	}

	fpsLastUpdated += frameTime
	if fpsLabel == "" || fpsLastUpdated > 1*time.Second {
		fpsLastUpdated = 0
		div := frameTime / time.Millisecond
		if div == 0 {
			div = 1
		}
		fps := math.Ceil(float64(1000 / div))
		fpsLabel = fmt.Sprintf("%d children rendered, %g FPS", len(manyRectChildren), fps)
		const fpsDebug = true
		if fpsDebug {
			log.Printf(fpsLabel)
		}
	}
	ret := nodes.Rectangle{
		Color:    sg.Color{1, 0, 1, 0},
		Size:     sz,
		Children: manyRectChildren,
		/*[]sg.Node{
		&sg.RectangleNode{ // just a way to get an extra child.. no actual scaling..
			Children: manyRectChildren,
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
						Text:       fpsLabel,
						FontFamily: "../shared/Barlow/Barlow-Regular.ttf",
						Color:      sg.Color{1, 1, 1, 1},
					},
				},
			},
		},*/
	}
	return ret
}
