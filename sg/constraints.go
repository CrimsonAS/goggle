package sg

import "math"

type Constraints struct {
	Min Size
	Max Size
}

func Unconstrained() Constraints {
	return Constraints{
		Min: Size{},
		Max: Size{float32(math.Inf(+1)), float32(math.Inf(+1))},
	}
}

func FixedConstraint(sz Size) Constraints {
	return Constraints{sz, sz}
}

func (c Constraints) BoundedConstraints(o Constraints) Constraints {
	o.Min = o.Min.Max(c.Min)
	o.Max = o.Max.Min(c.Max)
	return o
}

func (c Constraints) BoundedSize(sz Size) Size {
	return sz.Min(c.Max).Max(c.Min)
}

// Fit the given geometry within constraints (meaning, x+width must
// be within MaxWidth), bounding the size first if necessary.
//
// Really we need a much more interesting suite of these that can
// handle alignment, etc.
func (c Constraints) BoundedGeometrySize(geo Geometry) Geometry {
	maxSz := c.BoundedSize(geo.BottomRight().ToSize())
	return Geometry{
		geo.Origin,
		maxSz.SubPosition(geo.Origin),
	}
}
