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

type Size struct {
	Width, Height float32
}

func SizeV2(v Vec2) Size {
	return Size{v.X, v.Y}
}

func (s Size) Min(o Size) Size {
	if o.Width < s.Width {
		s.Width = o.Width
	}
	if o.Height < s.Height {
		s.Height = o.Height
	}
	return s
}

func (s Size) Max(o Size) Size {
	if o.Width > s.Width {
		s.Width = o.Width
	}
	if o.Height > s.Height {
		s.Height = o.Height
	}
	return s
}

func (s Size) Div(f float32) Size {
	return Size{s.Width / f, s.Height / f}
}

func (s Size) IsNil() bool {
	return s.Width == 0 && s.Width == 0
}

func (s Size) IsEmpty() bool {
	return s.Width <= 0 || s.Height <= 0
}

func (s Size) ToPosition() Position {
	return Position{s.Width, s.Height}
}

func (s Size) Sub(o Size) Size {
	return Size{s.Width - o.Width, s.Height - o.Height}
}

func (s Size) SubPosition(p Position) Size {
	return Size{s.Width - p.X, s.Height - p.Y}
}

func (s Size) String() string {
	return fmt.Sprintf("%gx%g", s.Width, s.Height)
}
