package sg

import "fmt"

type Position struct {
	X, Y float32
}

func PositionV2(v Vec2) Position {
	return Position{v.X, v.Y}
}

func (p Position) Min(o Position) Position {
	if o.X < p.X {
		p.X = o.X
	}
	if o.Y < p.Y {
		p.Y = o.Y
	}
	return p
}

func (p Position) Max(o Position) Position {
	if o.X > p.X {
		p.X = o.X
	}
	if o.Y > p.Y {
		p.Y = o.Y
	}
	return p
}

func (p Position) IsOrigin() bool {
	return p.X == 0 && p.Y == 0
}

func (p Position) Translate(o Position) Position {
	p.X += o.X
	p.Y += o.Y
	return p
}

func (p Position) Sub(o Position) Position {
	p.X -= o.X
	p.Y -= o.Y
	return p
}

func (p Position) ToSize() Size {
	return Size{p.X, p.Y}
}

func (p Position) Vec2() Vec2 {
	return Vec2{p.X, p.Y}
}

func (p Position) String() string {
	return fmt.Sprintf("%g,%g", p.X, p.Y)
}
