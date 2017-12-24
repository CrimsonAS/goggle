package components

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/layouts"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func Rectangle(cprops PropType, state *RenderState) sg.Node {
	rp := cprops.(RectangleProps)
	layoutType := layouts.Fixed
	if rp.Size.IsNil() {
		layoutType = layouts.Fill
	}

	return layouts.Box{
		Layout: layoutType,
		Props:  sg.Geometry{Size: rp.Size},
		Child: nodes.Rectangle{
			Color: rp.Color,
		},
	}
}

type RectangleProps struct {
	Size  sg.Size
	Color sg.Color
}
