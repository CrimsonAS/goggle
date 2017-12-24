package layouts

import (
	"github.com/CrimsonAS/goggle/sg"
)

func Fixed(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	geo := c.BoundedGeometrySize(props.(sg.Geometry))
	for _, child := range children {
		child.Render(sg.FixedConstraint(geo.Size))
		child.SetPosition(geo.Origin)
	}
	return geo.Size
}

func Fill(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	var maxChildSize sg.Size

	for _, child := range children {
		childSize := child.Render(sg.FixedConstraint(c.Max))
		maxChildSize = maxChildSize.Max(childSize)
		child.SetPosition(sg.Position{})
	}

	if len(children) > 0 {
		return maxChildSize
	} else {
		return c.Max
	}
}
