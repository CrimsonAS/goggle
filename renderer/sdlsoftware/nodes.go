package sdlsoftware

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

type DrawNode struct {
	Children      []sg.Node
	X, Y          float32
	Width, Height float32
	Draw          func(renderer *sdl.Renderer, node *DrawNode, transform sg.Transform)
}

func (node *DrawNode) GetChildren() []sg.Node {
	return node.Children
}

func (node *DrawNode) Size() sg.Vec2 {
	return sg.Vec2{node.Width, node.Height}
}

func (node *DrawNode) SetSize(sz sg.Vec2) {
	node.Width, node.Height = sz.X, sz.Y
}

func (node *DrawNode) Position() sg.Vec2 {
	return sg.Vec2{node.X, node.Y}
}

func (node *DrawNode) SetPosition(sz sg.Vec2) {
	node.X, node.Y = sz.X, sz.Y
}

func (node *DrawNode) CopyDrawable() sg.Drawable {
	re := *node
	re.Children = nil
	return &re
}
