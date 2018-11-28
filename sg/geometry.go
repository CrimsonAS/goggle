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

import (
	"fmt"
)

// Geometry represents a rectangle with a top-left origin position
// and a size.
type Geometry struct {
	Origin Position
	Size   Size
}

func (g Geometry) BottomRight() Position {
	return g.Origin.Translate(g.Size.ToPosition())
}

func (g Geometry) SetBottomRight(p Position) Geometry {
	return Geometry{
		g.Origin,
		p.Sub(g.Origin).ToSize(),
	}
}

func (g Geometry) Contains(p Position) bool {
	bottomRight := g.BottomRight()
	return (p.X >= g.Origin.X && p.X < bottomRight.X) &&
		(p.Y >= g.Origin.Y && p.Y < bottomRight.Y)
}

func (g Geometry) ContainsGeometry(g2 Geometry) bool {
	return g.Contains(g2.Origin) && g.Contains(g2.BottomRight())
}

func (g Geometry) Union(g2 Geometry) Geometry {
	bottomRight := g.BottomRight().Max(g2.BottomRight())
	g.Origin = g.Origin.Min(g2.Origin)
	return g.SetBottomRight(bottomRight)
}

// TransformedBounds returns a bounding box around this Geometry after
// applying the given transformation matrix.
//
// It's not possible to represent non-rectangular transformations in a
// Geometry. This function is a substitute for proper transformation only
// with trivial (translate+scale) transforms.
func (g Geometry) TransformedBounds(transform Mat4) Geometry {
	bottomRight := g.BottomRight()
	points := [4]Vec2{
		transform.MulV2(Vec2{g.Origin.X, g.Origin.Y}),
		transform.MulV2(Vec2{bottomRight.X, g.Origin.Y}),
		transform.MulV2(Vec2{bottomRight.X, bottomRight.Y}),
		transform.MulV2(Vec2{g.Origin.X, bottomRight.Y}),
	}

	tl, br := points[0], points[0]
	for i := 1; i < len(points); i++ {
		if points[i].X < tl.X {
			tl.X = points[i].X
		}
		if points[i].Y < tl.Y {
			tl.Y = points[i].Y
		}
		if points[i].X > br.X {
			br.X = points[i].X
		}
		if points[i].Y > br.Y {
			br.Y = points[i].Y
		}
	}

	return Geometry{
		Position{tl.X, tl.Y},
		Size{br.X - tl.X, br.Y - tl.Y},
	}
}

func (g Geometry) String() string {
	return fmt.Sprintf("[%v %v]", g.Origin, g.Size)
}
