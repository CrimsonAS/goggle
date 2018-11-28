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
