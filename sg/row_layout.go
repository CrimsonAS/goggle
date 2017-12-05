package sg

// Row is a positioner which places all of its Geometryable children
// horizontally adjacent to eachother, without regard for the total width.
type Row struct {
	Children []Node
	Padding  float32
}

// interface assertions
var _ Layouter = (*Row)(nil)

// ### Crazy idea: could there be a Row that is actually a []Node with
// functions, so it's possible to construct as just &Row{item, item}?

func (row *Row) GetChildren() []Node {
	return row.Children
}

func (row *Row) LayoutChildren(nodes []Geometryable) {
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
