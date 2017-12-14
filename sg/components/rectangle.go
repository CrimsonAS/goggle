package components

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func Rectangle(cprops PropType, state *RenderState) sg.Node {
	rp := cprops.(RectangleProps)
	return nodes.Transform{
		Matrix: sg.Translate2D(rp.Geometry.X, rp.Geometry.Y),
		Children: []sg.Node{
			nodes.Rectangle{
				Size:  rp.Geometry.Size(),
				Color: rp.Color,
			},
		},
	}
}

type RectangleProps struct {
	Geometry sg.Geometry
	Color    sg.Color
}
