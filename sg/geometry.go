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
