package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/layouts"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func RowLayout(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
	remainingWidth := c.Max.Width
	var x, maxChildY float32

	for i, child := range children {
		childConstraint := sg.Constraints{
			Max: sg.Size{
				Width:  remainingWidth / float32(len(children)-i),
				Height: c.Max.Height,
			},
		}
		childSize := child.Render(childConstraint)
		remainingWidth -= childSize.Width
		child.SetPosition(sg.Position{x, 0})
		x += childSize.Width
		if childSize.Height > maxChildY {
			maxChildY = childSize.Height
		}
	}

	return sg.Size{x, maxChildY}
}

func LayoutWindow(props components.PropType, state *components.RenderState) sg.Node {
	return nodes.Rectangle{
		Color: sg.Color{1, 0, 0, 1},
		Children: []sg.Node{
			layouts.Box{
				Layout: RowLayout,
				Children: []sg.Node{
					layouts.Box{
						Layout: layouts.Fill,
						Child: nodes.Rectangle{
							Color: sg.Color{1, 1, 1, 0},
						},
					},
					components.Component{
						Type: components.Rectangle,
						Props: components.RectangleProps{
							Color: sg.Color{1, 0.5, 0.5, 0},
							Size:  sg.Size{200, 200},
						},
					},
				},
			},
		},
	}
}

func main() {
	r, err := sdlsoftware.NewRenderer()
	if err != nil {
		panic(err)
	}

	w, err := r.CreateWindow()
	defer w.Destroy()

	if err != nil {
		panic(err)
	}

	for r.IsRunning() {
		r.ProcessEvents()
		w.Render(components.Component{Type: LayoutWindow})
	}

	r.Quit()
}
