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
