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

// Render a scene onto the window. The scene should have been synced recently;
// ideally once per render.
func (this *Window) Render(scene sg.TreeNode) {
	fmt.Printf("Rendering\n")
	surface, err := this.window.GetSurface()
	if err != nil {
		panic(err)
	}

	this.renderItem(scene, surface)
	this.window.UpdateSurface()
	fmt.Printf("Done rendering\n")
}

func (this *Window) renderItem(item sg.TreeNode, surface *sdl.Surface) {
	rootNode := item

	for {
		fmt.Printf("Rendering %+v\n", rootNode)
		if renderableNode, ok := rootNode.(sg.Renderable); ok {
			// if it's a Renderable, try reduce it to a real node of some kind
			fmt.Printf("Renderable. Going deeper.\n")
			rootNode = renderableNode.Render()
		} else {
			break
		}
	}

	// At this point, rootNode might be something we can draw.
	if rectangle, ok := rootNode.(*sg.Rectangle); ok {
		rect := sdl.Rect{int32(rectangle.X), int32(rectangle.Y), int32(rectangle.Width), int32(rectangle.Height)}
		fmt.Printf("Filling rect xy %fx%f wh %fx%f with color %s\n", rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height, rectangle.Color)

		// argb -> rgba
		var sdlColor uint32 = sdl.MapRGBA(surface.Format, uint8(255.0*rectangle.Color[1]), uint8(255.0*rectangle.Color[2]), uint8(255.0*rectangle.Color[3]), uint8(255.0*rectangle.Color[0]))
		surface.FillRect(&rect, sdlColor)
	}

	// ### this is a wee bit ugly, but we need to check if either rootNode (the
	// non-Renderable, decomposed node) or the original Item are of type
	// Nodeable.
	if treeChild, ok := item.(sg.Nodeable); ok {
		fmt.Printf("%+v: Nodable.\n", rootNode)
		for _, citem := range treeChild.GetChildren() {
			fmt.Printf("Examining child %+v\n", citem)
			this.renderItem(citem, surface)
		}
	} else if treeChild, ok := rootNode.(sg.Nodeable); ok {
		fmt.Printf("%+v: Nodable.\n", rootNode)
		for _, citem := range treeChild.GetChildren() {
			fmt.Printf("Examining child %+v\n", citem)
			this.renderItem(citem, surface)
		}
	} else {
		panic(fmt.Sprintf("Bad node type (not a pointer?) returned: %+v, %+v", item, rootNode))
	}
}
