package sg

import (
	"testing"
)

var rootNode Node

type simpleTestNode struct {
	X, Y, Width, Height float32
	Children            []Node
}

// Some simple benchmarks testing the performance impact of creating trees of
// items.

func renderSingleRectangle() Node {
	return &simpleTestNode{
		X:      5,
		Y:      5,
		Height: 100,
		Width:  100,
	}
}
func BenchmarkSingleRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSingleRectangle()
	}
}

func renderSingleRectangleWithChildrenRectangles() Node {
	return &simpleTestNode{
		X:      5,
		Y:      5,
		Height: 100,
		Width:  100,
		Children: []Node{
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
			},
		},
	}
}
func BenchmarkSingleRectangleWithChildrenRectangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSingleRectangleWithChildrenRectangles()
	}
}

func renderRectangleTree() Node {
	return &simpleTestNode{
		X:      5,
		Y:      5,
		Height: 100,
		Width:  100,
		Children: []Node{
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
				},
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
				},
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
				},
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
				},
			},
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
					},
				},
			},
		},
	}
}
func BenchmarkRectangleTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderRectangleTree()
	}
}

func renderRectangleDeepTree() Node {
	return &simpleTestNode{
		X:      5,
		Y:      5,
		Height: 100,
		Width:  100,
		Children: []Node{
			&simpleTestNode{
				X:      5,
				Y:      5,
				Height: 100,
				Width:  100,
				Children: []Node{
					&simpleTestNode{
						X:      5,
						Y:      5,
						Height: 100,
						Width:  100,
						Children: []Node{
							&simpleTestNode{
								X:      5,
								Y:      5,
								Height: 100,
								Width:  100,
								Children: []Node{
									&simpleTestNode{
										X:      5,
										Y:      5,
										Height: 100,
										Width:  100,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func BenchmarkRectangleDeepTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderRectangleDeepTree()
	}
}
