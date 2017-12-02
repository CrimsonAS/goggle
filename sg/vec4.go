package sg

import "fmt"

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (this Vec4) Add(other Vec4) Vec4 {
	return Vec4{
		X: this.X + other.X,
		Y: this.Y + other.Y,
		Z: this.Z + other.Z,
		W: this.W + other.W,
	}
}
func (this Vec4) Sub(other Vec4) Vec4 {
	return Vec4{
		X: this.X - other.X,
		Y: this.Y - other.Y,
		Z: this.Z - other.Z,
		W: this.W - other.W,
	}
}
func (this Vec4) Mul(other Vec4) Vec4 {
	return Vec4{
		X: this.X * other.X,
		Y: this.Y * other.Y,
		Z: this.Z * other.Z,
		W: this.W * other.W,
	}
}
func (this Vec4) Div(other Vec4) Vec4 {
	return Vec4{
		X: this.X / other.X,
		Y: this.Y / other.Y,
		Z: this.Z / other.Z,
		W: this.W / other.W,
	}
}
func (this Vec4) String() string {
	return fmt.Sprintf("%gx%gx%gx%g", this.X, this.Y, this.Z, this.W)
}
