package sg

// Positioner is a node that can reposition its children during rendering.
// A node implementing Positioner will have a call to PositionChildren
// during rendering, during which it can call SetPosition on these nodes.
type Positioner interface {
	Parentable

	// PositionChildren is called during rendering, before moving down the tree
	// to any child nodes. The Positioner may call SetPosition on these nodes.
	//
	// The nodes array is equal in size and corresponds in index to GetChildren().
	// For each child, the array will contain:
	//
	//   - If the child node is Geometryable, the child node itself. In this case,
	//     the child's Render() will not have been called yet.
	//   - If the child node is Renderable and the Render() function returned a
	//     Geometryable node, the rendered node. Render() will not be called again
	//     on the child after positioning.
	//   - nil, indicating that this child is not applicable to the positioner.
	//
	PositionChildren(nodes []Geometryable)
}

// Row is a positioner which places all of its Geometryable children
// horizontally adjacent to eachother, without regard for the total width.
type Row struct {
	Children []Node
	Padding  float32
}

// ### Crazy idea: could there be a Row that is actually a []Node with
// functions, so it's possible to construct as just &Row{item, item}?

func (row *Row) GetChildren() []Node {
	return row.Children
}

func (row *Row) PositionChildren(nodes []Geometryable) {
	x := float32(0)
	for _, node := range nodes {
		if node == nil {
			continue
		}
		childPos := node.Position()
		if x != childPos.X {
			childPos.X = x
			node.SetPosition(childPos)
		}
		sz := node.Size()
		x += sz.X + row.Padding
	}
}
