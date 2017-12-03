package sdlsoftware

import (
	"time"

	"github.com/CrimsonAS/goggle/renderer/private"
	"github.com/CrimsonAS/goggle/sg"
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
			win.inputHelper.MousePos = sg.Vec2{X: float32(t.X), Y: float32(t.Y)}
		case *sdl.MouseButtonEvent:
			win := this.windows[t.WindowID]
			if t.Type == sdl.MOUSEBUTTONUP {
				win.inputHelper.ButtonUp = true
			} else if t.Type == sdl.MOUSEBUTTONDOWN {
				win.inputHelper.ButtonDown = true
			}
		case *sdl.WindowEvent:
			win := this.windows[t.WindowID]
			if t.Event == sdl.WINDOWEVENT_LEAVE {
				win.inputHelper.MousePos = sg.Vec2{-1, -1} // ### this initial state isn't really acceptable, items may have negative coords.
			}
		}
	}
}

// Create (and show, for now) a window
// ### params needed
func (this *Renderer) CreateWindow() (*Window, error) {
	w := &Window{
		ourRenderer: this,
		inputHelper: private.NewInputHelper(),
	}
	var err error
	w.window, w.sdlRenderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	// ### a 'clear color' on the Window might make sense
	w.sdlRenderer.SetDrawColor(0, 0, 0, 0)
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
