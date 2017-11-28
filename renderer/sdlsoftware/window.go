package sdlsoftware

import (
	"fmt"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

// ### consider an interface when/if we want multiple renderers
type Renderer struct {
}

// Perform any initialization needed
func NewRenderer() (*Renderer, error) {
	return &Renderer{}, sdl.Init(sdl.INIT_EVERYTHING)
}

// Shut down
func (this *Renderer) Quit() {
	sdl.Quit()
}

// Spin the event loop
func (this *Renderer) ProcessEvents() {
	sdl.Delay(5)
}

// Create (and show, for now) a window
// ### params needed
func (this *Renderer) CreateWindow() (*Window, error) {
	w := &Window{}
	var err error
	w.window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type Window struct {
	window  *sdl.Window
	surface *sdl.Surface
}

// Destroy a window
func (this *Window) Destroy() {
	this.window.Destroy()
}

// Render a scene onto the window
func (this *Window) Render(scene sg.TreeNode) {
	fmt.Printf("Rendering\n")
	surface, err := this.window.GetSurface()
	if err != nil {
		panic(err)
	}

	// The strategy here is to render in two stages:
	// first walk the scene and reduce the tree to a list of
	// drawable primitives; then, iterate drawables and draw
	// them into the window. These could be split up further
	// in the future, and potentially happen across goroutines.
	drawables := this.renderItem(scene)
	fmt.Printf("Scene reduced to drawable: %+v\n", drawables)
	for _, node := range drawables {
		this.drawNode(surface, node)
	}

	this.window.UpdateSurface()
	fmt.Printf("Done rendering\n")
}

// renderItem walks a tree of nodes and reduces them to a list of drawable nodes
func (this *Window) renderItem(item sg.TreeNode) []sg.TreeNode {
	var drawables []sg.TreeNode

	// ### Need a proper test for what is actually drawable
	// Drawable stacks lowest for a node (below Render and any children)
	if _, ok := item.(*sg.Rectangle); ok {
		fmt.Printf("Found drawable: %+v\n", item)
		drawables = append(drawables, item)
	}

	// Render stacks next, below children
	if renderableNode, ok := item.(sg.Renderable); ok {
		fmt.Printf("Renderable. Going deeper.\n")
		rendered := renderableNode.Render()
		drawables = append(drawables, this.renderItem(rendered)...)
	}

	// Children stack in listed order from bottom to top
	for _, cNode := range item.GetChildren() {
		fmt.Printf("Examining child %+v\n", cNode)
		drawables = append(drawables, this.renderItem(cNode)...)
	}

	return drawables
}

func (this *Window) drawNode(surface *sdl.Surface, node sg.TreeNode) {
	if rectangle, ok := node.(*sg.Rectangle); ok {
		rect := sdl.Rect{int32(rectangle.X), int32(rectangle.Y), int32(rectangle.Width), int32(rectangle.Height)}
		fmt.Printf("Filling rect xy %fx%f wh %fx%f with color %s\n", rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height, rectangle.Color)

		// argb -> rgba
		var sdlColor uint32 = sdl.MapRGBA(surface.Format, uint8(255.0*rectangle.Color[1]), uint8(255.0*rectangle.Color[2]), uint8(255.0*rectangle.Color[3]), uint8(255.0*rectangle.Color[0]))
		surface.FillRect(&rect, sdlColor)
	} else {
		panic("unknown drawable")
	}
}
