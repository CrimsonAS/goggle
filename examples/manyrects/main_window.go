package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/layouts"
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
				manyRectChildren = append(manyRectChildren,
					components.Component{
						Type:  components.Rectangle,
						Props: components.RectangleProps{Color: sg.Color{1, 1, 1, 0}},
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
			childs = append(manyRectChildren,
				components.Component{
					Type:  components.Rectangle,
					Props: components.RectangleProps{Color: sg.Color{1, 1, 1, 0}},
				})
		}
		manyRectChildren = childs
	}

	// Randomize colors of all children
	for idx, child := range manyRectChildren {
		tchild := child.(components.Component)
		rprops := tchild.Props.(components.RectangleProps)

		const blend = false
		if blend {
			rprops.Color.X = localRand.Float32()
		}
		rprops.Color.Y = localRand.Float32()
		rprops.Color.Z = localRand.Float32()
		rprops.Color.W = localRand.Float32()

		tchild.Props = rprops
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
		Color: sg.Color{1, 0, 1, 0},
		Children: []sg.Node{
			layouts.Box{
				Layout: func(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
					// All children have a fixed size and a randomized position
					for _, child := range children {
						child.Render(sg.FixedConstraint(sg.Size{childSize, childSize}))
						child.SetPosition(sg.Position{
							localRand.Float32() * (c.Max.Width - childSize),
							localRand.Float32() * (c.Max.Height - childSize),
						})
					}
					return c.Max
				},
				Children: manyRectChildren,
			},
		},
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
