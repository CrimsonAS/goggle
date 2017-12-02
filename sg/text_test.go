package sg

import "testing"

var textBench *TextNode

func BenchmarkText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textBench = &TextNode{
			X:          0,
			Y:          1,
			Width:      2,
			Height:     3,
			Text:       "Hello, world",
			FontFamily: "Some/TTF.ttf",
			Children:   []Node{},
		}
	}
}

func BenchmarkTextWithChild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textBench = &TextNode{
			X:          0,
			Y:          1,
			Width:      2,
			Height:     3,
			Text:       "Hello, world",
			FontFamily: "Some/TTF.ttf",
			Children: []Node{
				&TextNode{
					X:          0,
					Y:          1,
					Width:      2,
					Height:     3,
					Text:       "Hello, world",
					FontFamily: "Some/TTF.ttf",
				},
			},
		}
	}
}
func BenchmarkTextWithColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textBench = &TextNode{
			X:          0,
			Y:          1,
			Width:      2,
			Height:     3,
			Color:      Color{1, 2, 3, 4},
			Text:       "Hello, world",
			FontFamily: "Some/TTF.ttf",
			Children:   []Node{},
		}
	}
}
