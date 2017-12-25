package layouts

import (
	"fmt"

	"github.com/CrimsonAS/goggle/sg"
)

type FlexDirection int

const (
	FlexRow FlexDirection = iota
	FlexRowReverse
	FlexColumn
	FlexColumnReverse
)

type PixelUnit float32
type PercentUnit float32

type FlexLayoutProps struct {
	Direction FlexDirection
	// Not implemented: wrap-reverse
	// Not implemented: Wrap bool
	// Not implemented: justify-content
	// Not implemented: align-items
	// Not implemented: align-content
}

type FlexChildProps struct {
	// Not implemented: order
	Grow   int
	Shrink int
	// Not implemented: (min-|max-|fit-|)content
	// Basis is one of PixelUnit, PercentUnit, or nil (auto)
	Basis interface{}
	// Not implemented: 'flex' gives sensible defaults, like 0, 1, auto.
	//  - Could make some functions to do decent common behaviors
	// Not implemented: align-self
}

type flexState struct {
	props FlexLayoutProps
	basis []float32
}

func Flex(c sg.Constraints, children []BoxChild, props interface{}) sg.Size {
	flex := flexState{}
	flex.props, _ = props.(FlexLayoutProps)
	maxLength := flex.MainAxisLength(c.Max)

	flex.basis = make([]float32, len(children))
	var sumBasis, sumShrinkBasis float32
	var sumGrow int

	// Determine and sum basis for all children
	for i, child := range children {
		cp, _ := child.Props().(FlexChildProps)

		if pxBasis, specified := flex.PixelBasis(cp.Basis, maxLength); specified {
			flex.basis[i] = pxBasis
		} else {
			// ### 'auto' is giving the child max constraint and using returned dimension
			// as basis. I'm not sure if that really makes sense, but here we are.
			panic("re-render during layout is broken")
			size := child.Render(c)
			flex.basis[i] = flex.MainAxisLength(size)
		}

		sumBasis += flex.basis[i]
		sumGrow += cp.Grow
		sumShrinkBasis += flex.basis[i] * float32(cp.Shrink)
	}

	// Grow or shrink for available space
	avail := maxLength - sumBasis
	if avail > 0 {
		// Grow
		if sumGrow == 0 {
			// XXX This is a bad way to handle it.
			panic("no grow")
		}
		growUnit := avail / float32(sumGrow)

		for i, child := range children {
			cp, _ := child.Props().(FlexChildProps)
			flex.basis[i] += growUnit * float32(cp.Grow)
		}
	} else if avail < 0 {
		// Shrink
		for i, child := range children {
			cp, _ := child.Props().(FlexChildProps)
			shrinkFactor := (flex.basis[i] * float32(cp.Shrink)) / sumShrinkBasis
			flex.basis[i] += shrinkFactor * avail
		}
	}

	// Render everything and first pass on positioning
	var p sg.Position
	for i, child := range children {
		max := c.Max
		if flex.IsHorizontal() {
			max.Width = flex.basis[i]
		} else {
			max.Height = flex.basis[i]
		}
		size := child.Render(sg.Constraints{Max: max})

		child.SetPosition(p)
		p = p.Translate(flex.CrossAxisZero(size).ToPosition())
	}

	if flex.IsReversed() {
		// Loop back over and invert positions
		for _, child := range children {
			geo := child.Geometry()
			child.SetPosition(p.Sub(geo.Origin).Sub(flex.CrossAxisZero(geo.Size).ToPosition()))
		}
	}

	return p.ToSize()
}

func (f flexState) IsHorizontal() bool {
	return f.props.Direction == FlexRow || f.props.Direction == FlexRowReverse
}

func (f flexState) IsReversed() bool {
	return f.props.Direction == FlexRowReverse || f.props.Direction == FlexColumnReverse
}

// MainAxisSize returns a Size with v in the main axis
func (f flexState) MainAxisSize(v float32) sg.Size {
	if f.IsHorizontal() {
		return sg.Size{Width: v}
	} else {
		return sg.Size{Height: v}
	}
}

// MainAxisLength returns the main axis dimension from a Size
func (f flexState) MainAxisLength(size sg.Size) float32 {
	if f.IsHorizontal() {
		return size.Width
	} else {
		return size.Height
	}
}

// CrossAxisZero returns a copy of Size with the cross axis zeroed
func (f flexState) CrossAxisZero(size sg.Size) sg.Size {
	if f.IsHorizontal() {
		return sg.Size{Width: size.Width}
	} else {
		return sg.Size{Height: size.Height}
	}
}

func (f flexState) PixelBasis(basis interface{}, parentLength float32) (float32, bool) {
	switch v := basis.(type) {
	case PixelUnit:
		return float32(v), true
	case PercentUnit:
		return (parentLength * (float32(v) / 100)), true
	case nil:
		return 0, false
	default:
		panic(fmt.Sprintf("Unknown Basis value in Flex layout %T %+v", basis, basis))
		return 0, false
	}
}
