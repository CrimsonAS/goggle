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
