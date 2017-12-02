package sg

import (
	"math"
)

func Translate2D(dx float32, dy float32) Mat4 {
	return NewMat4(
		1, 0, 0, dx,
		0, 1, 0, dy,
		0, 0, 1, 0,
		0, 0, 0, 1,
		Translation2DType)
}
func Translate2DV2(d Vec2) Mat4 {
	return Translate2D(d.X, d.Y)
}
func Scale2D(sx float32, sy float32) Mat4 {
	return NewMat4(
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
		Scale2DType)
}
func Scale2DV2(s Vec2) Mat4 {
	return Scale2D(s.X, s.Y)
}
func Rotate2D(radians float32) Mat4 {
	s := float32(math.Sin(float64(radians)))
	c := float32(math.Cos(float64(radians)))
	return NewMat4(
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
		Rotation2DType)
}
func Translate(dx float32, dy float32, dz float32) Mat4 {
	return NewMat4(
		1, 0, 0, dx,
		0, 1, 0, dy,
		0, 0, 1, dz,
		0, 0, 0, 1,
		Translation2DType)
}
func RotateAroundZ(radians float32) Mat4 {
	return Rotate2D(radians)
}
func RotateAroundX(radians float32) Mat4 {
	s := float32(math.Sin(float64(radians)))
	c := float32(math.Cos(float64(radians)))
	return NewMat4(
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
		GenericType)
}
func RotateAroundY(radians float32) Mat4 {
	s := float32(math.Sin(float64(radians)))
	c := float32(math.Cos(float64(radians)))
	return NewMat4(
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
		GenericType)
}
func Scale(sx float32, sy float32, sz float32) Mat4 {
	return NewMat4(
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
		GenericType)
}

func (this Mat4) IsIdentity() bool {
	return this.Type == IdentityType
}

/*
###
    float operator()(int c, int r) {
        return m[c*4 + r];
    }
*/
