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

import "testing"

func assertColorString(t *testing.T, index int, c Color, s string) {
	cs := c.String()
	if cs != s {
		t.Fatalf("For color %d, expected string %s, got string %s", index, s, cs)
	}
}

type colorTest struct {
	color          Color
	expectedString string
}

func TestColorString(t *testing.T) {
	tests := []colorTest{
		colorTest{color: Color{1, 0, 0, 0}, expectedString: "#ff000000"},
		colorTest{color: Color{0, 1, 0, 0}, expectedString: "#00ff0000"},
		colorTest{color: Color{0, 0, 1, 0}, expectedString: "#0000ff00"},
		colorTest{color: Color{0, 0, 0, 1}, expectedString: "#000000ff"},

		colorTest{color: Color{0.5, 0, 0, 0}, expectedString: "#7f000000"},
		colorTest{color: Color{0, 0.5, 0, 0}, expectedString: "#007f0000"},
		colorTest{color: Color{0, 0, 0.5, 0}, expectedString: "#00007f00"},
		colorTest{color: Color{0, 0, 0, 0.5}, expectedString: "#0000007f"},

		colorTest{color: Color{0.1, 0, 0, 0}, expectedString: "#19000000"},
		colorTest{color: Color{0, 0.1, 0, 0}, expectedString: "#00190000"},
		colorTest{color: Color{0, 0, 0.1, 0}, expectedString: "#00001900"},
		colorTest{color: Color{0, 0, 0, 0.1}, expectedString: "#00000019"},

		colorTest{color: Color{1, 1, 1, 1}, expectedString: "#ffffffff"},
		colorTest{color: Color{0, 0, 0, 0}, expectedString: "#00000000"},
	}

	for idx, test := range tests {
		assertColorString(t, idx, test.color, test.expectedString)
	}
}
