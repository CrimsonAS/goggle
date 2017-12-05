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
