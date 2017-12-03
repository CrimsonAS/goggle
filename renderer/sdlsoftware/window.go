package sdlsoftware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CrimsonAS/goggle/renderer/private"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Window struct {
	window      *sdl.Window
	sdlRenderer *sdl.Renderer
	ourRenderer *Renderer
	inputHelper private.InputHelper

	endLastFrame  time.Time
	frameDuration time.Duration
}

func (this *Window) GetSize() sg.Vec2 {
	ww, hh := this.window.GetSize()
	return sg.Vec2{float32(ww), float32(hh)}
}

// Returns the time between frames. Used to advance animations.
func (this *Window) FrameTime() time.Duration {
	return this.frameDuration
}

// Destroy a window
func (this *Window) Destroy() {
	id, err := this.window.GetID()
	if err != nil {
		panic("No window ID!")
	}
	delete(this.ourRenderer.windows, id)
	this.window.Destroy()
}

// Render a scene onto the window
func (this *Window) Render(scene sg.Node) {
	debugOut("Rendering\n")

	this.frameDuration = time.Since(this.endLastFrame)
	this.endLastFrame = time.Now()

	// ### a 'clear color' on the Window might make sense
	this.sdlRenderer.SetDrawColor(0, 0, 0, 0)
	this.sdlRenderer.Clear()

	// The strategy here is to render in two stages:
	// first walk the scene and reduce the tree to a list of
	// drawable primitives; then, iterate drawables and draw
	// them into the window. These could be split up further
	// in the future, and potentially happen across goroutines.
	drawables := this.renderItem(scene, nil, sg.Vec2{X: 0, Y: 0}, 1.0, 1.0)

	debugOut("scene rendered to %d drawables\n", len(drawables))
	for _, node := range drawables {
		debugOut("drawing node %s: %+v\n", sg.NodeName(node), node)
		scale := float32(1.0)
		rotation := float32(1.0)
		switch node := node.(type) {
		case sg.Scaleable:
			scale = node.GetScale()
		case sg.Rotateable:
			rotation = node.GetRotation()
		}
		this.drawNode(node, scale, rotation)
	}

	this.sdlRenderer.Present()

	elapsed := time.Since(this.ourRenderer.start) / time.Millisecond
	sleepyTime := (1000/60 - elapsed) * time.Millisecond
	const fpsDebug = false
	if fpsDebug {

		div := elapsed
		if div == 0 {
			div = 1
		}
		fmt.Printf("Done rendering in %s @ %d FPS, sleeping %s\n", time.Since(this.ourRenderer.start), 1000/div, sleepyTime)
	}

	time.Sleep(sleepyTime) // cap rendering
	this.inputHelper.ResetFrameState()
}

// renderItem walks a tree of nodes and reduces them to a list of drawable nodes.
// origin translate the item's coordinates, such that origin.X + item.X
// is the left side of the item in window coordinates.
func (this *Window) renderItem(item, itemRendered sg.Node, origin sg.Vec2, scale, rotation float32) []sg.Node {
	var drawables []sg.Node

	debugOut("rendering node %s (%s) to origin (%s): %+v\n",
		sg.NodeName(item),
		strings.Join(sg.NodeInterfaces(item), " "),
		origin,
		item)

	// Drawable stacks lowest for a node (below Render and any children)
	if draw, ok := item.(sg.Drawable); ok {
		// Copy instance for safe modification & independent draw
		draw = draw.CopyDrawable()

		if positionable, ok := draw.(sg.Positionable); ok {
			// Offset position with origin
			pos := positionable.Position()
			pos = pos.Add(origin)
			positionable.SetPosition(pos)
		}

		if sizeable, ok := draw.(sg.Sizeable); ok {
			sz := sizeable.Size()
			sz.X *= scale
			sz.Y *= scale
			sizeable.SetSize(sz)
		}

		if rotateable, ok := draw.(sg.Rotateable); ok {
			// ### wrong wrong wrong?
			childRotation := rotateable.GetRotation()
			childRotation *= rotation
			rotateable.SetRotation(childRotation)
		}

		drawables = append(drawables, draw)
	}

	// If node is a Positionable, adjust origin for relative
	// coordinates in the rendered tree and in children
	if geo, ok := item.(sg.Positionable); ok {
		cpos := geo.Position()
		origin = origin.Add(cpos)
	}

	if scaleable, ok := item.(sg.Scaleable); ok {
		scale *= scaleable.GetScale()
	}

	if rotateable, ok := item.(sg.Rotateable); ok {
		rotation *= rotateable.GetRotation()
	}

	if sizeable, ok := item.(sg.Sizeable); ok {
		childSize := sizeable.Size()
		// ### this isn't really right. I think we should traverse the tree of
		// renderables twice: once to deliver input events (and this must be
		// done in paint order, so deepest children first), recursing up to
		// parents.
		this.inputHelper.ProcessPointerEvents(origin, childSize.X, childSize.Y, item)
	}

	// Render stacks next, below children
	if renderableNode, ok := item.(sg.Renderable); ok {
		if itemRendered == nil {
			itemRendered = renderableNode.Render(this)
		}
		drawables = append(drawables, this.renderItem(itemRendered, nil, origin, scale, rotation)...)
	}

	// If this item is a positioner, iterate children to process geometry. For some
	// children, this may also require calling Render, in which case we need to save
	// the rendered node for renderItem on that child.
	if positioner, ok := item.(sg.Positioner); ok {
		children := positioner.GetChildren()
		geoNodes := make([]sg.Geometryable, len(children))
		renderNodes := make([]sg.Node, len(children))

		for i, cNode := range children {
			if cGeo, ok := cNode.(sg.Geometryable); ok {
				geoNodes[i] = cGeo
			} else if renderableNode, ok := cNode.(sg.Renderable); ok {
				renderNodes[i] = renderableNode.Render(this)
				if cGeo, ok := renderNodes[i].(sg.Geometryable); ok {
					geoNodes[i] = cGeo
				}
			}
		}

		positioner.PositionChildren(geoNodes)

		for i, cNode := range children {
			drawables = append(drawables, this.renderItem(cNode, renderNodes[i], origin, scale, rotation)...)
		}
	} else if parentNode, ok := item.(sg.Parentable); ok {
		// Children stack in listed order from bottom to top
		for _, cNode := range parentNode.GetChildren() {
			drawables = append(drawables, this.renderItem(cNode, nil, origin, scale, rotation)...)
		}
	}

	return drawables
}

func (this *Window) drawRectangle(node *sg.RectangleNode, scale, rotation float32) {
	w := node.Width * scale
	h := node.Height * scale
	rect := sdl.Rect{int32(node.X), int32(node.Y), int32(w), int32(h)}
	debugOut("Filling rect xy %gx%g wh %gx%g with color %v\n", node.X, node.Y, w, h, node.Color)

	// argb -> rgba
	this.sdlRenderer.SetDrawColor(uint8(255.0*node.Color.Y), uint8(255.0*node.Color.Z), uint8(255.0*node.Color.W), uint8(255.0*node.Color.X))
	if node.Color.X == 1 {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	} else {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	}
	this.sdlRenderer.FillRect(&rect)
}

func (this *Window) drawImage(node *sg.ImageNode, scale, rotation float32) {
	w := node.Width * scale
	h := node.Height * scale
	var fileTexture *sg.FileTexture
	var err error
	var ok bool

	if fileTexture, ok = node.Texture.(*sg.FileTexture); !ok {
		panic("unknown texture")
	}

	// ### file caching
	image, err := img.LoadTexture(this.sdlRenderer, fileTexture.Source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load source: %s (%s)\n", fileTexture.Source, err.Error())
		return
	}

	// ###? defer image.Free()
	rect := sdl.Rect{int32(node.X), int32(node.Y), int32(w), int32(h)}
	this.sdlRenderer.Copy(image, nil, &rect)
}

func (this *Window) drawText(node *sg.TextNode, scale, rotation float32) {
	w := node.Width * scale
	h := node.Height * scale
	// ### font caching (and database)
	var font *ttf.Font
	var err error

	if font, err = ttf.OpenFont(node.FontFamily, node.PixelSize); err != nil {
		fmt.Fprint(os.Stderr, "Failed to open font %s (%s)\n", node.FontFamily, err)
		return
	}
	defer font.Close()

	sdlColor := sdl.Color{uint8(255.0 * node.Color.Y), uint8(255.0 * node.Color.Z), uint8(255.0 * node.Color.W), uint8(255.0 * node.Color.X)}

	var renderedText *sdl.Surface
	if renderedText, err = font.RenderUTF8Blended(node.Text, sdlColor); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		return
	}
	defer renderedText.Free()

	texture, err := this.sdlRenderer.CreateTextureFromSurface(renderedText)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to get texture for text: %s\n", err)
	} else {
		// ###? defer texture.Free()
		rect := sdl.Rect{int32(node.X), int32(node.Y), int32(w), int32(h)}
		if node.Color.X == 1 {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		} else {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		}
		this.sdlRenderer.Copy(texture, nil, &rect)
	}

}

func (this *Window) drawNode(baseNode sg.Node, scale, rotation float32) {
	switch node := baseNode.(type) {
	case *sg.RectangleNode:
		this.drawRectangle(node, scale, rotation)
	case *sg.ImageNode:
		this.drawImage(node, scale, rotation)
	case *sg.TextNode:
		this.drawText(node, scale, rotation)
	case *DrawNode:
		debugOut("Calling custom draw function %+v\n", node.Draw)
		node.Draw(this.sdlRenderer, node)

	default:
		panic("unknown drawable")
	}
}
