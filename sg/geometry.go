package sg

import (
	"fmt"
)

type Size struct {
	Width, Height float32
}

type Position struct {
	X, Y float32
}

type Geometry struct {
	X, Y, Width, Height float32
}

type Constraints struct {
	MinWidth, MinHeight float32
	MaxWidth, MaxHeight float32
}

func (g Geometry) XYWH() Vec4 {
	return Vec4{g.X, g.Y, g.Width, g.Height}
}

func (g Geometry) XYXY() Vec4 {
	return Vec4{g.X, g.Y, g.X + g.Width, g.X + g.Height}
}

func (g Geometry) Pos() Vec2 {
	return Vec2{g.X, g.Y}
}

func (g Geometry) Size() Vec2 {
	return Vec2{g.Width, g.Height}
}

func (g Geometry) Translate(x, y float32) Geometry {
	g.X += x
	g.Y += y
	return g
}

func (g Geometry) ZeroOrigin() Geometry {
	g.X, g.Y = 0, 0
	return g
}

func (g Geometry) BottomRight() Vec2 {
	return Vec2{g.X + g.Width, g.Y + g.Height}
}

func (g Geometry) Contains(x, y float32) bool {
	return (x >= g.X && x <= g.X+g.Width) && (y >= g.Y && y <= g.Y+g.Height)
}

func (g Geometry) ContainsV2(point Vec2) bool {
	return g.Contains(point.X, point.Y)
}

func (g Geometry) ContainsGeometry(g2 Geometry) bool {
	return g.ContainsV2(g2.Pos()) && g.ContainsV2(g2.BottomRight())
}

func (g Geometry) Union(g2 Geometry) Geometry {
	if g2.X < g.X {
		g.X = g2.X
	}
	if g2.Y < g.Y {
		g.Y = g2.Y
	}
	if g2.Width > g.Width {
		g.Width = g2.Width
	}
	if g2.Height > g.Height {
		g.Height = g2.Height
	}
	return g
}

// TransformedBounds returns a bounding box around this Geometry after
// applying the given transformation matrix.
//
// It's not possible to represent non-rectangular transformations in a
// Geometry. This function is a substitute for proper transformation only
// with trivial (translate+scale) transforms.
func (g Geometry) TransformedBounds(transform Mat4) Geometry {
	points := [4]Vec2{
		transform.MulV2(Vec2{g.X, g.Y}),
		transform.MulV2(Vec2{g.X + g.Width, g.Y}),
		transform.MulV2(Vec2{g.X + g.Width, g.Y + g.Height}),
		transform.MulV2(Vec2{g.X, g.Y + g.Height}),
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

	return Geometry{tl.X, tl.Y, br.X - tl.X, br.Y - tl.Y}
}

func (g Geometry) String() string {
	return fmt.Sprintf("[%g,%g %gx%g]", g.X, g.Y, g.Width, g.Height)
}
