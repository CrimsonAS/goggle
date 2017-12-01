package sdlsoftware

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// ### consider an interface when/if we want multiple renderers
type Renderer struct {
	isRunning bool
	start     time.Time // when rendering this frame began
	window    *Window   // ### fixme input handling
}

// Perform any initialization needed
func NewRenderer() (*Renderer, error) {
	r := &Renderer{isRunning: true}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	if err := ttf.Init(); err != nil {
		return nil, err
	}

	return r, nil
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
	this.start = time.Now()
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			this.isRunning = false
		case *sdl.MouseMotionEvent:
			// ### windowId to find the window
			this.window.mousePos = sg.TouchPoint{X: float32(t.X), Y: float32(t.Y)}
			//fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
		case *sdl.MouseButtonEvent:
			if t.Type == sdl.MOUSEBUTTONUP {
				this.window.buttonUp = true
			} else if t.Type == sdl.MOUSEBUTTONDOWN {
				this.window.buttonDown = true
			}
		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_LEAVE {
				log.Printf("Cursor left window")
				this.window.mousePos = sg.TouchPoint{-1, -1} // ### this initial state isn't really acceptable, items may have negative coords.
			}
		}
	}
}

// Create (and show, for now) a window
// ### params needed
func (this *Renderer) CreateWindow() (*Window, error) {
	w := &Window{
		ourRenderer:     this,
		hoveredNodes:    make(map[sg.Node]bool),
		oldHoveredNodes: make(map[sg.Node]bool),
		mousePos:        sg.TouchPoint{-1, -1},
	}
	this.window = w
	var err error
	w.window, w.sdlRenderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type Window struct {
	window          *sdl.Window
	sdlRenderer     *sdl.Renderer
	ourRenderer     *Renderer
	hoveredNodes    map[sg.Node]bool
	oldHoveredNodes map[sg.Node]bool
	mousePos        sg.TouchPoint
	buttonUp        bool
	buttonDown      bool
	mouseGrabber    sg.Node
}

// Destroy a window
func (this *Window) Destroy() {
	this.window.Destroy()
}

func debugOut(fstr string, vals ...interface{}) {
	const debug = false

	if debug {
		log.Printf(fstr, vals...)
	}
}

// Render a scene onto the window
func (this *Window) Render(scene sg.Node) {
	debugOut("Rendering\n")

	// ### a 'clear color' on the Window might make sense
	this.sdlRenderer.SetDrawColor(0, 0, 0, 0)
	this.sdlRenderer.Clear()

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

	// restore state
	this.oldHoveredNodes = this.hoveredNodes
	this.hoveredNodes = make(map[sg.Node]bool)

	this.buttonDown = false
	this.buttonUp = false
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
		childX, childY, childWidth, childHeight := geo.Geometry()
		originX += childX
		originY += childY

		if hoverable, ok := item.(sg.Hoverable); ok {
			if this.mousePos.X >= originX &&
				this.mousePos.Y >= originY &&
				this.mousePos.X <= originX+childWidth &&
				this.mousePos.Y <= originY+childHeight {
				this.hoveredNodes[item] = true
				if _, ok = this.oldHoveredNodes[item]; !ok {
					log.Printf("Mouse at %fx%f entering item %+v geom %fx%f %fx%f", this.mousePos.X, this.mousePos.Y, item, originX, originY, originX+childWidth, originY+childWidth)
					hoverable.PointerEnter(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
				}
			} else if _, ok = this.oldHoveredNodes[item]; ok {
				log.Printf("Mouse at %fx%f leaving item %+v geom %fx%f %fx%f", this.mousePos.X, this.mousePos.Y, item, originX, originY, originX+childWidth, originY+childWidth)
				hoverable.PointerLeave(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
			}
		}
		if this.buttonDown || this.buttonUp {
			if tappable, ok := item.(sg.Tappable); ok {
				if this.buttonDown {
					this.mouseGrabber = item
				} else if this.buttonUp {
					if this.mouseGrabber == item {
						tappable.PointerTapped(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
					}
				}
			}
		}
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

func (this *Window) drawRectangle(node *sg.Rectangle) {
	rect := sdl.Rect{int32(node.X), int32(node.Y), int32(node.Width), int32(node.Height)}
	debugOut("Filling rect xy %gx%g wh %gx%g with color %v\n", node.X, node.Y, node.Width, node.Height, node.Color)

	// argb -> rgba
	this.sdlRenderer.SetDrawColor(uint8(255.0*node.Color[1]), uint8(255.0*node.Color[2]), uint8(255.0*node.Color[3]), uint8(255.0*node.Color[0]))
	if node.Color[0] == 1 {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	} else {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	}
	this.sdlRenderer.FillRect(&rect)
}

func (this *Window) drawImage(node *sg.Image) {
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
	rect := sdl.Rect{int32(node.X), int32(node.Y), int32(node.Width), int32(node.Height)}
	this.sdlRenderer.Copy(image, nil, &rect)
}

func (this *Window) drawText(node *sg.Text) {
	// ### font caching (and database)
	var font *ttf.Font
	var err error

	if font, err = ttf.OpenFont(node.FontFamily, node.PixelSize); err != nil {
		fmt.Fprint(os.Stderr, "Failed to open font %s (%s)\n", node.FontFamily, err)
		return
	}
	defer font.Close()

	sdlColor := sdl.Color{uint8(255.0 * node.Color[1]), uint8(255.0 * node.Color[2]), uint8(255.0 * node.Color[3]), uint8(255.0 * node.Color[0])}

	var renderedText *sdl.Surface
	if renderedText, err = font.RenderUTF8Blended("Hello, World!", sdlColor); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		return
	}
	defer renderedText.Free()

	texture, err := this.sdlRenderer.CreateTextureFromSurface(renderedText)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to get texture for text: %s\n", err)
	} else {
		// ###? defer texture.Free()
		rect := sdl.Rect{int32(node.X), int32(node.Y), int32(node.Width), int32(node.Height)}
		if node.Color[0] == 1 {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		} else {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		}
		this.sdlRenderer.Copy(texture, nil, &rect)
	}

}

func (this *Window) drawNode(baseNode sg.Node) {
	switch node := baseNode.(type) {
	case *sg.Rectangle:
		this.drawRectangle(node)
	case *sg.Image:
		this.drawImage(node)
	case *sg.Text:
		this.drawText(node)
	case *DrawNode:
		debugOut("Calling custom draw function %+v\n", node.Draw)
		node.Draw(this.sdlRenderer, node)

	default:
		panic("unknown drawable")
	}
}
