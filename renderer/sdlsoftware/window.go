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
	sdl.Delay(2500)
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

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	this.window.UpdateSurface()

	this.renderItem(scene)
	fmt.Printf("Done rendering\n")
}

func (this *Window) renderItem(item sg.TreeNode) {
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

	// At this point, rootNode should be something we can draw.
	fmt.Printf("Can draw %+v\n", rootNode)

	// ### this is a wee bit ugly, but we need to check if either rootNode (the
	// non-Renderable, decomposed node) or the original Item are of type
	// Nodeable.
	if treeChild, ok := item.(sg.Nodeable); ok {
		fmt.Printf("%+v: Nodable.\n", rootNode)
		for _, citem := range treeChild.GetChildren() {
			fmt.Printf("Examining child %+v\n", citem)
			this.renderItem(citem)
		}
	} else if treeChild, ok := rootNode.(sg.Nodeable); ok {
		fmt.Printf("%+v: Nodable.\n", rootNode)
		for _, citem := range treeChild.GetChildren() {
			fmt.Printf("Examining child %+v\n", citem)
			this.renderItem(citem)
		}
	} else {
		panic(fmt.Sprintf("Bad node type (not a pointer?) returned: %+v, %+v", item, rootNode))
	}
}
