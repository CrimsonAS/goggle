package sdlsoftware

import (
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

type SDLDrawNode struct {
	sg.BasicNode
	Draw func(surface *sdl.Surface, node *SDLDrawNode)
}

func (node *SDLDrawNode) CopyDrawable() sg.Drawable {
	re := &SDLDrawNode{node.BasicNode, node.Draw}
	re.Children = nil
	return re
}
