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
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/CrimsonAS/goggle/renderer/private"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/layouts"
	"github.com/CrimsonAS/goggle/sg/nodes"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Window struct {
	window      *sdl.Window
	sdlRenderer *sdl.Renderer
	ourRenderer *Renderer
	inputHelper InputHelper

	endLastFrame  time.Time
	frameDuration time.Duration
	blendMode     sdl.BlendMode

	id uint32

	sceneRenderer SceneRenderer
}

func (this *Window) GetSize() sg.Size {
	ww, hh := this.window.GetSize()
	return sg.Size{float32(ww), float32(hh)}
}

// Returns the time between frames. Used to advance animations.
func (this *Window) FrameTime() time.Duration {
	return this.frameDuration
}

// Destroy a window
func (this *Window) Destroy() {
	delete(this.ourRenderer.windows, this.id)
	this.window.Destroy()
}

func (this *Window) CreateSDLTexture(w, h float32, alpha bool) (*sdl.Texture, error) {
	fmt := sdl.PIXELFORMAT_RGBX8888
	if alpha {
		fmt = sdl.PIXELFORMAT_RGBA8888
	}

	return this.sdlRenderer.CreateTexture(uint32(fmt), sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
}

const renderDebug = false       // this is expensive..
const headlessRendering = false // turn off all rendering (for use in benchmarking algorithms w/o SDL intereference)

// Render a scene onto the window
func (this *Window) Render(scene sg.Node) {
	if renderDebug {
		log.Printf("Rendering")
	}

	windowSize := this.GetSize()
	windowConstraints := sg.Constraints{Max: windowSize}
	windowBox := layouts.Box{
		Layout: func(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
			for _, child := range children {
				child.Render(windowConstraints)
				child.SetPosition(sg.Position{0, 0})
			}
			return windowSize
		},
		Child: scene,
	}

	this.frameDuration = time.Since(this.endLastFrame)
	this.endLastFrame = time.Now()

	this.sdlRenderer.Clear()
	this.sceneRenderer.DeliverEvents()
	this.sceneRenderer.Render(windowBox)
	this.sceneRenderer.Draw(this.drawNode)
	this.sdlRenderer.Present()

	elapsed := time.Since(this.ourRenderer.start) / time.Millisecond
	sleepyTime := (1000/60 - elapsed) * time.Millisecond
	const fpsDebug = true
	if fpsDebug {

		div := elapsed
		if div == 0 {
			div = 1
		}
		log.Printf("Done rendering in %s @ %d FPS, sleeping %s", time.Since(this.ourRenderer.start), 1000/div, sleepyTime)
	}

	if !headlessRendering {
		time.Sleep(sleepyTime) // cap rendering
	}
	this.inputHelper.ResetFrameState()
}

func sdlGeometry(geo sg.Geometry) sdl.Rect {
	return sdl.Rect{
		int32(geo.Origin.X),
		int32(geo.Origin.Y),
		int32(geo.Size.Width),
		int32(geo.Size.Height),
	}
}

func (this *Window) drawRectangle(node nodes.Rectangle, transform sg.Mat4, size sg.Size) {
	// ### This is wrong for non-trivial transforms, but I don't want to mess with SDL
	// enough to draw complex shapes for now.
	geo := sg.Geometry{Size: size}.TransformedBounds(transform)
	if headlessRendering {
		return
	}
	rect := sdlGeometry(geo)
	if renderDebug {
		log.Printf("Filling rect %v with color %v", geo, node.Color)
	}
	// argb -> rgba
	this.sdlRenderer.SetDrawColor(uint8(255.0*node.Color.Y), uint8(255.0*node.Color.Z), uint8(255.0*node.Color.W), uint8(255.0*node.Color.X))
	if node.Color.X == 1 {
		this.setBlendMode(sdl.BLENDMODE_NONE)
	} else {
		this.setBlendMode(sdl.BLENDMODE_BLEND)
	}
	this.sdlRenderer.FillRect(&rect)
}

func (this *Window) drawImage(node nodes.Image, transform sg.Mat4, size sg.Size) {
	geo := sg.Geometry{Size: size}.TransformedBounds(transform)
	var texture *sdl.Texture

	if renderDebug {
		log.Printf("Drawing image from %T at %v", node.Texture, geo)
	}

	if headlessRendering {
		return
	}

	switch t := node.Texture.(type) {
	case nodes.FileTexture:
		image, err := img.LoadTexture(this.sdlRenderer, string(t))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load source: %s (%s)\n", string(t), err.Error())
			return
		}
		defer image.Destroy()
		texture = image

	case *sdl.Texture:
		texture = t

	default:
		panic("unknown texture")
	}

	rect := sdlGeometry(geo)
	this.sdlRenderer.Copy(texture, nil, &rect)
}

func (this *Window) drawText(node nodes.Text, transform sg.Mat4, size sg.Size) {
	geo := sg.Geometry{Size: size}.TransformedBounds(transform)

	if headlessRendering {
		return
	}

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
		defer texture.Destroy()
		rect := sdlGeometry(geo)
		if node.Color.X == 1 {
			this.setBlendMode(sdl.BLENDMODE_NONE)
		} else {
			this.setBlendMode(sdl.BLENDMODE_BLEND)
		}
		this.sdlRenderer.Copy(texture, nil, &rect)
	}
}

func (this *Window) setBlendMode(bm sdl.BlendMode) {
	// this is cheaper than hitting cgo
	if this.blendMode != bm {
		this.sdlRenderer.SetDrawBlendMode(bm)
		this.blendMode = bm
	}
}

func (this *Window) drawNode(node sg.Node, transform sg.Mat4, size sg.Size) {
	if renderDebug {
		log.Printf("drawing node %s at %v: %+v transform:[%+v]", sg.NodeName(node), size, node, transform)
	}
	switch cnode := node.(type) {
	case nodes.Rectangle:
		this.drawRectangle(cnode, transform, size)
	case nodes.Image:
		this.drawImage(cnode, transform, size)
	case nodes.Text:
		this.drawText(cnode, transform, size)

	default:
		panic(fmt.Sprintf("unknown drawable %T %+v", node, node))
	}
}
