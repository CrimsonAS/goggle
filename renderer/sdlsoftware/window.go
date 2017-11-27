package sdlsoftware

import "github.com/CrimsonAS/goggle/sg"
import "github.com/veandco/go-sdl2/sdl"

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
func (this *Window) Render(scene *sg.Scene) {
	surface, err := this.window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	this.window.UpdateSurface()
}
