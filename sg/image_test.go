package sg

import "testing"

var imageBench *ImageNode

func BenchmarkImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imageBench = &ImageNode{
			X:        0,
			Y:        1,
			Width:    2,
			Height:   3,
			Children: []Node{},
		}
	}
}

func BenchmarkImageWithTexture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imageBench = &ImageNode{
			X:        0,
			Y:        1,
			Width:    2,
			Height:   3,
			Texture:  &FileTexture{Source: "HelloWorld.png"},
			Children: []Node{},
		}
	}
}

func BenchmarkImageWithChild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imageBench = &ImageNode{
			X:       0,
			Y:       1,
			Width:   2,
			Height:  3,
			Texture: &FileTexture{Source: "HelloWorld.png"},
			Children: []Node{
				&ImageNode{
					X:       0,
					Y:       1,
					Width:   2,
					Height:  3,
					Texture: &FileTexture{Source: "HelloWorld.png"},
				},
			},
		}
	}
}
