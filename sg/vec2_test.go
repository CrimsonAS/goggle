package sg

import (
	"fmt"
	"testing"
)

func testAddInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Add(v2)
	if v3 != expected {
		t.Fatalf("Add %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestAdd(t *testing.T) {
	testAddInternal(t, Vec2{X: 100, Y: 100}, Vec2{X: 50, Y: 50}, Vec2{X: 150, Y: 150})
	testAddInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -50, Y: -150})
	testAddInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 150, Y: -250})
}

func testSubInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Sub(v2)
	if v3 != expected {
		t.Fatalf("Sub %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestSub(t *testing.T) {
	testSubInternal(t, Vec2{X: 100, Y: 100}, Vec2{X: 50, Y: 50}, Vec2{X: 50, Y: 50})
	testSubInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -150, Y: -250})
	testSubInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 50, Y: -150})
}

func testMulInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Mul(v2)
	if v3 != expected {
		t.Fatalf("Mul %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestMul(t *testing.T) {
	testMulInternal(t, Vec2{X: 2, Y: 2}, Vec2{X: 2, Y: 4}, Vec2{X: 4, Y: 8})
	testMulInternal(t, Vec2{X: -100, Y: -200}, Vec2{X: 50, Y: 50}, Vec2{X: -5000, Y: -10000})
	testMulInternal(t, Vec2{X: 100, Y: -200}, Vec2{X: 50, Y: -50}, Vec2{X: 5000, Y: 10000})
}

func testDivInternal(t *testing.T, v1 Vec2, v2 Vec2, expected Vec2) {
	v3 := v1.Div(v2)
	if v3 != expected {
		t.Fatalf("Div %s %s gave: %s -- wanted %s", v1, v2, v3, expected)
	}
}
func TestDiv(t *testing.T) {
	testDivInternal(t, Vec2{X: 4, Y: 8}, Vec2{X: 2, Y: 4}, Vec2{X: 2, Y: 2})
}

func TestString(t *testing.T) {
	v := Vec2{X: 100, Y: 200}
	str := fmt.Sprintf("%gx%g", v.X, v.Y)
	if v.String() != str {
		t.Fatalf("Wanted %s, got %s", str, v.String())
	}
}
