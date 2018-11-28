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
func (this Vec4) XY() Vec2 {
	return Vec2{this.X, this.Y}
}
func (this Vec4) ZW() Vec2 {
	return Vec2{this.Z, this.W}
}
func (this Vec4) String() string {
	return fmt.Sprintf("%gx%gx%gx%g", this.X, this.Y, this.Z, this.W)
}
