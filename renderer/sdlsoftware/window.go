package sdlsoftware

import (
	"fmt"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
	"time"
)

// ### consider an interface when/if we want multiple renderers
type Renderer struct {
	isRunning bool
}

// Perform any initialization needed
func NewRenderer() (*Renderer, error) {
	return &Renderer{isRunning: true}, sdl.Init(sdl.INIT_EVERYTHING)
}

func (this *Renderer) IsRunning() bool {
	return this.isRunning
}

// Shut down
func (this *Renderer) Quit() {
	sdl.Quit()
}

// Spin the event loop
func (this *Renderer) ProcessEvents() {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			this.isRunning = false
		case *sdl.MouseMotionEvent:
			fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
		}
	}
}

// Create (and show, for now) a window
// ### params needed
func (this *Renderer) CreateWindow() (*Window, error) {
	w := &Window{}
	var err error
	w.window, w.renderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type Window struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// Destroy a window
func (this *Window) Destroy() {
	this.window.Destroy()
}

func debugOut(fstr string, vals ...interface{}) {
	const debug = false

	if debug {
		fmt.Printf(fstr, vals...)
	}
}

// Render a scene onto the window
func (this *Window) Render(scene sg.Node) {
	start := time.Now()

	debugOut("Rendering\n")

	// ### a 'clear color' on the Window might make sense
	this.renderer.SetDrawColor(0, 0, 0, 0)
	this.renderer.Clear()

	// The strategy here is to render in two stages:
	// first walk the scene and reduce the tree to a list of
	// drawable primitives; then, iterate drawables and draw
	// them into the window. These could be split up further
	// in the future, and potentially happen across goroutines.
	drawables := this.renderItem(scene, 0, 0)

	debugOut("scene rendered to %d drawables\n", len(drawables))
	for _, node := range drawables {
		debugOut("drawing node %s: %+v\n", sg.NodeName(node), node)
		this.drawNode(node)
	}

	this.renderer.Present()

	elapsed := time.Since(start) / time.Millisecond
	if elapsed == 0 {
		elapsed = 1
	}

	const fpsDebug = false
	if fpsDebug {
		fmt.Printf("Done rendering in %d ms, %d FPS\n", elapsed, 1000/elapsed)
	}
}

// renderItem walks a tree of nodes and reduces them to a list of drawable nodes.
// originX and originY translate the item's coordinates, such that originX + item.X
// is the left side of the item in window coordinates.
func (this *Window) renderItem(item sg.Node, originX, originY float32) []sg.Node {
	var drawables []sg.Node

	debugOut("rendering node %s (%s) to origin (%g,%g): %+v\n",
		sg.NodeName(item),
		strings.Join(sg.NodeInterfaces(item), " "),
		originX, originY,
		item)

	// Drawable stacks lowest for a node (below Render and any children)
	if draw, ok := item.(sg.Drawable); ok {
		// Copy instance for safe modification & independent draw
		draw = draw.CopyDrawable()

		if geo, ok := draw.(sg.GeometryNode); ok {
			// Offset position with originX/originY
			x, y, w, h := geo.Geometry()
			x += originX
			y += originY
			geo.SetGeometry(x, y, w, h)
		}

		drawables = append(drawables, draw)
	}

	// If node is a GeometryNode, adjust originX/originY for relative
	// coordinates in the rendered tree and in children
	if geo, ok := item.(sg.GeometryNode); ok {
		childX, childY, _, _ := geo.Geometry()
		originX += childX
		originY += childY
	}

	// Render stacks next, below children
	if renderableNode, ok := item.(sg.Renderable); ok {
		rendered := renderableNode.Render()
		drawables = append(drawables, this.renderItem(rendered, originX, originY)...)
	}

	// Children stack in listed order from bottom to top
	if parentNode, ok := item.(sg.Parentable); ok {
		for _, cNode := range parentNode.GetChildren() {
			drawables = append(drawables, this.renderItem(cNode, originX, originY)...)
		}
	}

	// ### A node that is not geometry, renderable, drawable, or parentable
	// has no effect on rendering and is plausibly a bug. Is it worth erroring?

	return drawables
}

func (this *Window) drawNode(baseNode sg.Node) {
	switch node := baseNode.(type) {
	case *sg.Rectangle:
		rect := sdl.Rect{int32(node.X), int32(node.Y), int32(node.Width), int32(node.Height)}
		debugOut("Filling rect xy %gx%g wh %gx%g with color %v\n", node.X, node.Y, node.Width, node.Height, node.Color)

		// argb -> rgba
		this.renderer.SetDrawColor(uint8(255.0*node.Color[1]), uint8(255.0*node.Color[2]), uint8(255.0*node.Color[3]), uint8(255.0*node.Color[0]))
		this.renderer.FillRect(&rect)
	case *sg.Image:
		if fileTexture, ok := node.Texture.(*sg.FileTexture); ok {
			image, err := img.LoadTexture(this.renderer, fileTexture.Source)
			rect := sdl.Rect{int32(node.X), int32(node.Y), int32(node.Width), int32(node.Height)}
			if err != nil {
				debugOut("Failed to load source: %s (%s)\n", fileTexture.Source, err.Error())
			} else {
				this.renderer.Copy(image, nil, &rect)
			}
		} else {
			panic("unknown texture")
		}

	case *DrawNode:
		debugOut("Calling custom draw function %+v\n", node.Draw)
		node.Draw(this.renderer, node)

	default:
		panic("unknown drawable")
	}
}
