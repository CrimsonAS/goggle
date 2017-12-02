package sg

import "fmt"

type Vec2 struct {
	X float32
	Y float32
}

func (this Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: this.X + other.X,
		Y: this.Y + other.Y,
	}
}

func (this Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: this.X - other.X,
		Y: this.Y - other.Y,
	}
}

func (this Vec2) Mul(other Vec2) Vec2 {
	return Vec2{
		X: this.X * other.X,
		Y: this.Y * other.Y,
	}
}

func (this Vec2) Div(other Vec2) Vec2 {
	return Vec2{
		X: this.X / other.X,
		Y: this.Y / other.Y,
	}
}

func (this Vec2) String() string {
	return fmt.Sprintf("%gx%g", this.X, this.Y)
}
