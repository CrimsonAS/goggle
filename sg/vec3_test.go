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
