package layouts

import (
	"github.com/CrimsonAS/goggle/sg"
)

func None(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	var maxChildSize sg.Size
	for _, child := range children {
		childSize := child.Render(c)
		maxChildSize = maxChildSize.Max(childSize)
		child.SetPosition(sg.Position{})
	}
	if len(children) > 0 {
		return maxChildSize
	} else {
		return c.Max
	}
}

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

func FillAspect(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	var maxSize sg.Size
	ratio, _ := props.(float32)
	if ratio == 0 {
		return sg.Size{0, 0}
	}

	// Ratio is w/h; find largest size of that aspect ratio in constraint
	newMax := c.Max
	newMax.Height = c.Max.Width / ratio
	if newMax.Height > c.Max.Height {
		newMax.Height = c.Max.Height
		newMax.Width = newMax.Height * ratio
	}

	for _, child := range children {
		size := child.Render(sg.FixedConstraint(newMax))
		maxSize = maxSize.Max(size)
		child.SetPosition(sg.Position{})
	}

	if len(children) > 0 {
		return maxSize
	} else {
		return newMax
	}
}
