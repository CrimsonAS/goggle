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

func renderSingleTestNode() Node {
	return &simpleTestNode{
		X:      5,
		Y:      5,
		Height: 100,
		Width:  100,
	}
}
func BenchmarkRenderSingleTestNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSingleTestNode()
	}
}

func renderSingleTestNodeWithChildren() Node {
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
func BenchmarkRenderSingleTestNodeWithChildren(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSingleTestNodeWithChildren()
	}
}

func renderSimpleNodeTree() Node {
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
func BenchmarkRenderSimpleNodeTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSimpleNodeTree()
	}
}

func renderSimpleNodeDeepTree() Node {
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
func BenchmarkRenderSimpleNodeDeepTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootNode = renderSimpleNodeDeepTree()
	}
}
