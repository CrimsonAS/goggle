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
	"fmt"
	"strconv"
)

// A R G B
type Color Vec4

func twoDigitHex(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%s", strconv.FormatInt(int64(n), 16))
	} else {
		return fmt.Sprintf("%s", strconv.FormatInt(int64(n), 16))
	}
}

func (this Color) String() string {
	return fmt.Sprintf("#%s%s%s%s",
		twoDigitHex(this.Aint()),
		twoDigitHex(this.Rint()),
		twoDigitHex(this.Gint()),
		twoDigitHex(this.Bint()))
}

func (this Color) A() float32 {
	return this.X
}
func (this Color) Aint() int {
	return int(this.X * 255.0)
}
func (this Color) R() float32 {
	return this.Y
}
func (this Color) Rint() int {
	return int(this.Y * 255.0)
}
func (this Color) G() float32 {
	return this.Z
}
func (this Color) Gint() int {
	return int(this.Z * 255.0)
}
func (this Color) B() float32 {
	return this.W
}
func (this Color) Bint() int {
	return int(this.W * 255.0)
}
