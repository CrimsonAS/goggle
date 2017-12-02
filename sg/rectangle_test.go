package sg

import "testing"

var rectangleBench *RectangleNode

func BenchmarkRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectangleBench = &RectangleNode{
			X:        0,
			Y:        1,
			Width:    2,
			Height:   3,
			Children: []Node{},
		}
	}
}

func BenchmarkRectangleWithChild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectangleBench = &RectangleNode{
			X:      0,
			Y:      1,
			Width:  2,
			Height: 3,
			Children: []Node{
				&RectangleNode{
					X:      0,
					Y:      1,
					Width:  2,
					Height: 3,
				},
			},
		}
	}
}
func BenchmarkRectangleWithColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectangleBench = &RectangleNode{
			X:        0,
			Y:        1,
			Width:    2,
			Height:   3,
			Color:    Color{1, 2, 3, 4},
			Children: []Node{},
		}
	}
}
