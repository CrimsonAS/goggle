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
	"testing"
)

func testV2AddInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Add(v2)
	if v3 != expected {
		t.Fatalf("Add %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV2Add(t *testing.T) {
	testV2AddInternal(t, Vec2{X: 100, Y: 100}, Vec2{X: 50, Y: 50}, Vec2{X: 150, Y: 150})
	testV2AddInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -50, Y: -150})
	testV2AddInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 150, Y: -250})
}

func testV2SubInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Sub(v2)
	if v3 != expected {
		t.Fatalf("Sub %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV2Sub(t *testing.T) {
	testV2SubInternal(t, Vec2{X: 100, Y: 100}, Vec2{X: 50, Y: 50}, Vec2{X: 50, Y: 50})
	testV2SubInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -150, Y: -250})
	testV2SubInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 50, Y: -150})
}

func testV2MulInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Mul(v2)
	if v3 != expected {
		t.Fatalf("Mul %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV2Mul(t *testing.T) {
	testV2MulInternal(t, Vec2{X: 2, Y: 2}, Vec2{X: 2, Y: 4}, Vec2{X: 4, Y: 8})
	testV2MulInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -5000, Y: -10000})
	testV2MulInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 5000, Y: 10000})
}

func testV2DivInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Div(v2)
	if v3 != expected {
		t.Fatalf("Div %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV2Div(t *testing.T) {
	testV2DivInternal(t, Vec2{X: 4, Y: 8}, Vec2{X: 2, Y: 4}, Vec2{X: 2, Y: 2})
}

func TestV2String(t *testing.T) {
	v := Vec2{X: 100, Y: 200}
	str := fmt.Sprintf("%gx%g", v.X, v.Y)
	if v.String() != str {
		t.Fatalf("Wanted %s, got %s", str, v.String())
	}
}
