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
	windows   map[uint32]*Window
}

// Perform any initialization needed
func NewRenderer() (*Renderer, error) {
	r := &Renderer{
		isRunning: true,
		windows:   make(map[uint32]*Window),
	}
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
			win := this.windows[t.WindowID]
			win.mousePos = sg.TouchPoint{X: float32(t.X), Y: float32(t.Y)}
		case *sdl.MouseButtonEvent:
			win := this.windows[t.WindowID]
			if t.Type == sdl.MOUSEBUTTONUP {
				win.buttonUp = true
			} else if t.Type == sdl.MOUSEBUTTONDOWN {
				win.buttonDown = true
			}
		case *sdl.WindowEvent:
			win := this.windows[t.WindowID]
			if t.Event == sdl.WINDOWEVENT_LEAVE {
				win.mousePos = sg.TouchPoint{-1, -1} // ### this initial state isn't really acceptable, items may have negative coords.
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
	var err error
	w.window, w.sdlRenderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	id, err := w.window.GetID()
	if err != nil {
		return nil, err
	}
	this.windows[id] = w
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
	id, err := this.window.GetID()
	if err != nil {
		panic("No window ID!")
	}
	delete(this.ourRenderer.windows, id)
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

	if this.mouseGrabber != nil {
		if moveable, ok := this.mouseGrabber.(sg.Moveable); ok {
			moveable.PointerMoved(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
		}
	}

	// ### a 'clear color' on the Window might make sense
	this.sdlRenderer.SetDrawColor(0, 0, 0, 0)
	this.sdlRenderer.Clear()

	// The strategy here is to render in two stages:
	// first walk the scene and reduce the tree to a list of
	// drawable primitives; then, iterate drawables and draw
	// them into the window. These could be split up further
	// in the future, and potentially happen across goroutines.
	drawables := this.renderItem(scene, 0, 0, 1.0, 1.0)

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

	// restore state
	this.oldHoveredNodes = this.hoveredNodes
	this.hoveredNodes = make(map[sg.Node]bool)

	this.buttonDown = false
	this.buttonUp = false
}

// ### should scale/rotate affect input events? i'd say yes, personally.
func (this *Window) processPointerEvents(originX, originY, childWidth, childHeight float32, item sg.Node) {
	// BUG: ### unsolved problems: we should also probably block propagation of hover.
	// We could have a return code to block hover propagating further down the tree,
	// letting someone write code like:
	//
	// Root UI node
	//     Sidebar PointerEnter() { return true; /* block */ }
	//         Button Hoverable // to highlight as need be
	//     UI page
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
		if pressable, ok := item.(sg.Pressable); ok {
			if this.buttonDown {
				if this.mouseGrabber == nil {
					this.mouseGrabber = item
					pressable.PointerPressed(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
				}
			} else if this.buttonUp {
				if this.mouseGrabber == item {
					pressable.PointerReleased(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
				}
			}
		}
		if tappable, ok := item.(sg.Tappable); ok {
			if this.buttonDown {
				if this.mouseGrabber == nil {
					// a Tappable takes an implicit grab
					this.mouseGrabber = item
				}
			} else if this.buttonUp {
				// BUG: right now, PointerTapped is called regardless of whether or not the
				// release happens inside the item boundary.
				if this.mouseGrabber == item {
					tappable.PointerTapped(sg.TouchPoint{this.mousePos.X, this.mousePos.Y})
				}
			}
		}
		if this.buttonUp && this.mouseGrabber == item {
			this.mouseGrabber = nil
		}
	}
}

// renderItem walks a tree of nodes and reduces them to a list of drawable nodes.
// originX and originY translate the item's coordinates, such that originX + item.X
// is the left side of the item in window coordinates.
func (this *Window) renderItem(item sg.Node, originX, originY, scale, rotation float32) []sg.Node {
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

		if scaleable, ok := draw.(sg.Scaleable); ok {
			childScale := scaleable.GetScale()
			childScale *= scale
			scaleable.SetScale(childScale)
		}

		if rotateable, ok := draw.(sg.Rotateable); ok {
			childRotation := rotateable.GetRotation()
			childRotation *= rotation
			rotateable.SetRotation(childRotation)
		}

		drawables = append(drawables, draw)
	}

	// If node is a GeometryNode, adjust originX/originY for relative
	// coordinates in the rendered tree and in children
	if geo, ok := item.(sg.GeometryNode); ok {
		childX, childY, childWidth, childHeight := geo.Geometry()
		originX += childX
		originY += childY

		if scaleable, ok := geo.(sg.Scaleable); ok {
			scale *= scaleable.GetScale()
		}

		if rotateable, ok := geo.(sg.Rotateable); ok {
			rotation *= rotateable.GetRotation()
		}

		// ### this isn't really right. I think we should traverse the tree of
		// renderables twice: once to deliver input events (and this must be
		// done in paint order, so deepest children first), recursing up to
		// parents.
		this.processPointerEvents(originX, originY, childWidth, childHeight, item)
	}

	// Render stacks next, below children
	if renderableNode, ok := item.(sg.Renderable); ok {
		rendered := renderableNode.Render()
		drawables = append(drawables, this.renderItem(rendered, originX, originY, scale, rotation)...)
	}

	// Children stack in listed order from bottom to top
	if parentNode, ok := item.(sg.Parentable); ok {
		for _, cNode := range parentNode.GetChildren() {
			drawables = append(drawables, this.renderItem(cNode, originX, originY, scale, rotation)...)
		}
	}

	// ### A node that is not geometry, renderable, drawable, or parentable
	// has no effect on rendering and is plausibly a bug. Is it worth erroring?

	return drawables
}

func (this *Window) drawRectangle(node *sg.Rectangle, scale, rotation float32) {
	w := node.Width * scale
	h := node.Height * scale
	rect := sdl.Rect{int32(node.X), int32(node.Y), int32(w), int32(h)}
	debugOut("Filling rect xy %gx%g wh %gx%g with color %v\n", node.X, node.Y, w, h, node.Color)

	// argb -> rgba
	this.sdlRenderer.SetDrawColor(uint8(255.0*node.Color[1]), uint8(255.0*node.Color[2]), uint8(255.0*node.Color[3]), uint8(255.0*node.Color[0]))
	if node.Color[0] == 1 {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	} else {
		this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	}
	this.sdlRenderer.FillRect(&rect)
}

func (this *Window) drawImage(node *sg.Image, scale, rotation float32) {
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

func (this *Window) drawText(node *sg.Text, scale, rotation float32) {
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
		rect := sdl.Rect{int32(node.X), int32(node.Y), int32(w), int32(h)}
		if node.Color[0] == 1 {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		} else {
			this.sdlRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		}
		this.sdlRenderer.Copy(texture, nil, &rect)
	}

}

func (this *Window) drawNode(baseNode sg.Node, scale, rotation float32) {
	switch node := baseNode.(type) {
	case *sg.Rectangle:
		this.drawRectangle(node, scale, rotation)
	case *sg.Image:
		this.drawImage(node, scale, rotation)
	case *sg.Text:
		this.drawText(node, scale, rotation)
	case *DrawNode:
		debugOut("Calling custom draw function %+v\n", node.Draw)
		node.Draw(this.sdlRenderer, node)

	default:
		panic("unknown drawable")
	}
}
