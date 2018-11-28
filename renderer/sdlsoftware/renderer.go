/*
 * Copyright 2017 Crimson AS <info@crimson.no>
 * Author: Robin Burchell <robin.burchell@crimson.no>
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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
			win.inputHelper.MousePos = sg.Position{X: float32(t.X), Y: float32(t.Y)}
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
				win.inputHelper.MousePos = sg.Position{-1, -1} // ### this initial state isn't really acceptable, items may have negative coords.
			}
		}
	}
}

// Create (and show, for now) a window
// ### params needed
func (this *Renderer) CreateWindow() (*Window, error) {
	w := &Window{
		ourRenderer:  this,
		inputHelper:  private.NewInputHelper(),
		endLastFrame: time.Now(),
	}

	// Construct the scene renderer and resolve the scenegraph, including
	// delivery of input events. The result is a list of DrawableNode.
	w.sceneRenderer = private.SceneRenderer{
		Window:      w,
		InputHelper: &w.inputHelper,
	}
	var err error
	w.window, w.sdlRenderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	// ### a 'clear color' on the Window might make sense
	w.sdlRenderer.SetDrawColor(0, 0, 0, 0)
	w.id, err = w.window.GetID()
	if err != nil {
		return nil, err
	}
	this.windows[w.id] = w
	if err != nil {
		return nil, err
	}
	return w, nil
}
