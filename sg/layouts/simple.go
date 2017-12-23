package layouts

import (
	"github.com/CrimsonAS/goggle/sg"
)

func Fixed(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	geo := c.BoundedGeometrySize(props.(sg.Geometry))
	for _, child := range children {
		child.Render(sg.FixedConstraint(sg.Size{geo.Width, geo.Height}))
		child.SetPosition(sg.Position{geo.X, geo.Y})
	}
	return sg.Size{geo.Width, geo.Height}
}

func Fill(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	size := sg.Size{c.MaxWidth, c.MaxHeight}
	var maxChildSize sg.Size

	for _, child := range children {
		childSize := child.Render(sg.FixedConstraint(size))
		maxChildSize = maxChildSize.Max(childSize)
		child.SetPosition(sg.Position{0, 0})
	}

	if len(children) > 0 {
		return maxChildSize
	} else {
		return size
	}
}
