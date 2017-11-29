package sdlsoftware

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

type DrawNode struct {
	Children      []sg.Node
	X, Y          float32
	Width, Height float32
	Draw          func(renderer *sdl.Renderer, node *DrawNode)
}

func (node *DrawNode) GetChildren() []sg.Node {
	return node.Children
}

func (node *DrawNode) Geometry() (x, y, w, h float32) {
	return node.X, node.Y, node.Width, node.Height
}

func (node *DrawNode) SetGeometry(x, y, w, h float32) {
	node.X, node.Y, node.Width, node.Height = x, y, w, h
}

func (node *DrawNode) CopyDrawable() sg.Drawable {
	re := *node
	re.Children = nil
	return &re
}
