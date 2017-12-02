package sg

import "fmt"

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (this Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: this.X + other.X,
		Y: this.Y + other.Y,
		Z: this.Z + other.Z,
	}
}
func (this Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: this.X - other.X,
		Y: this.Y - other.Y,
		Z: this.Z - other.Z,
	}
}
func (this Vec3) Mul(other Vec3) Vec3 {
	return Vec3{
		X: this.X * other.X,
		Y: this.Y * other.Y,
		Z: this.Z * other.Z,
	}
}
func (this Vec3) Div(other Vec3) Vec3 {
	return Vec3{
		X: this.X / other.X,
		Y: this.Y / other.Y,
		Z: this.Z / other.Z,
	}
}
func (this Vec3) Project2D(farPlane float32) Vec2 {
	zScale := (farPlane - this.Z) / farPlane
	return Vec2{this.X / zScale, this.Y / zScale}
}
func (this Vec3) String() string {
	return fmt.Sprintf("%gx%gx%g", this.X, this.Y, this.Z)
}
