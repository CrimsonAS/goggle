/*
 * Copyright 2017 Crimson AS <info@crimson.no>
 * Author: Robin Burchell <robin.burchell@crimson.no>
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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
