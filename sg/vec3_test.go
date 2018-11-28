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

func testV3AddInternal(t *testing.T, v1 Vec3, v2 Vec3, expected Vec3) {
	v3 := v1.Add(v2)
	if v3 != expected {
		t.Fatalf("Add %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV3Add(t *testing.T) {
	testV3AddInternal(t, Vec3{X: 100, Y: 100, Z: 10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: 150, Y: 150, Z: 60})
	testV3AddInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: -50, Y: -150, Z: 40})
	testV3AddInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: -50, Y: -50, Z: -50}, Vec3{X: -150, Y: -250, Z: -60})
}

func testV3SubInternal(t *testing.T, v1 Vec3, v2 Vec3, expected Vec3) {
	v3 := v1.Sub(v2)
	if v3 != expected {
		t.Fatalf("Sub %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV3Sub(t *testing.T) {
	testV3SubInternal(t, Vec3{X: 100, Y: 100, Z: 10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: 50, Y: 50, Z: -40})
	testV3SubInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: -150, Y: -250, Z: -60})
	testV3SubInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: -50, Y: -50, Z: -50}, Vec3{X: -50, Y: -150, Z: 40})
}

func testV3MulInternal(t *testing.T, v1 Vec3, v2 Vec3, expected Vec3) {
	v3 := v1.Mul(v2)
	if v3 != expected {
		t.Fatalf("Mul %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV3Mul(t *testing.T) {
	testV3MulInternal(t, Vec3{X: 100, Y: 100, Z: 10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: 5000, Y: 5000, Z: 500})
	testV3MulInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: 50, Y: 50, Z: 50}, Vec3{X: -5000, Y: -10000, Z: -500})
	testV3MulInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: -50, Y: -50, Z: -50}, Vec3{X: 5000, Y: 10000, Z: 500})
}

func testV3DivInternal(t *testing.T, v1 Vec3, v2 Vec3, expected Vec3) {
	v3 := v1.Div(v2)
	if v3 != expected {
		t.Fatalf("Div %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestV3Div(t *testing.T) {
	testV3DivInternal(t, Vec3{X: 100, Y: 100, Z: 10}, Vec3{X: 50, Y: 50, Z: 20}, Vec3{X: 2, Y: 2, Z: 0.5})
	testV3DivInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: 50, Y: 50, Z: 20}, Vec3{X: -2, Y: -4, Z: -0.5})
	testV3DivInternal(t, Vec3{X: -100, Y: -200, Z: -10}, Vec3{X: -50, Y: -50, Z: -20}, Vec3{X: 2, Y: 4, Z: 0.5})
}

func TestV3String(t *testing.T) {
	v := Vec3{X: 100, Y: 200, Z: 300}
	str := fmt.Sprintf("%gx%gx%g", v.X, v.Y, v.Z)
	if v.String() != str {
		t.Fatalf("Wanted %s, got %s", str, v.String())
	}
}
