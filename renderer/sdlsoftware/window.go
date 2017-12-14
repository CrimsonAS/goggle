package sdlsoftware

import (
	"fmt"
	"log"
	"time"

	. "github.com/CrimsonAS/goggle/renderer/private"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	window      *sdl.Window
	sdlRenderer *sdl.Renderer
	ourRenderer *Renderer
	inputHelper InputHelper

	endLastFrame  time.Time
	frameDuration time.Duration
	blendMode     sdl.BlendMode

	sceneRenderer SceneRenderer
}

func (this *Window) GetSize() sg.Vec2 {
	ww, hh := this.window.GetSize()
	return sg.Vec2{float32(ww), float32(hh)}
}

// Returns the time between frames. Used to advance animations.
func (this *Window) FrameTime() time.Duration {
	return this.frameDuration
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

const renderDebug = false       // this is expensive..
const headlessRendering = false // turn off all rendering (for use in benchmarking algorithms w/o SDL intereference)

// Render a scene onto the window
func (this *Window) Render(scene sg.Node) {
	if renderDebug {
		log.Printf("Rendering")
	}

	this.frameDuration = time.Since(this.endLastFrame)
	this.endLastFrame = time.Now()

	this.sdlRenderer.Clear()
	this.sceneRenderer.DeliverEvents()
	this.sceneRenderer.Render(scene)
	this.sceneRenderer.Draw(this.drawNode)
	this.sdlRenderer.Present()

	elapsed := time.Since(this.ourRenderer.start) / time.Millisecond
	sleepyTime := (1000/60 - elapsed) * time.Millisecond
	const fpsDebug = false
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

func (this *Window) drawRectangle(node sg.SimpleRectangleNode, transform sg.Mat4) {
	// ### This is wrong for non-trivial transforms, but I don't want to mess with SDL
	// enough to draw complex shapes for now.
	geo := sg.Geometry{0, 0, node.Size.X, node.Size.Y}.TransformedBounds(transform)
	if headlessRendering {
		return
	}
	rect := sdl.Rect{int32(geo.X), int32(geo.Y), int32(geo.Width), int32(geo.Height)}
	if renderDebug {
		log.Printf("Filling rect xy %gx%g wh %gx%g with color %v", geo.X, geo.Y, geo.Width, geo.Height, node.Color)
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

/*
func (this *Window) drawImage(node *sg.ImageNode, transform sg.Transform) {
	geo := transform.Geometry(sg.Vec4{0, 0, node.Width, node.Height})
	var fileTexture *sg.FileTexture
	var err error
	var ok bool

	if fileTexture, ok = node.Texture.(*sg.FileTexture); !ok {
		panic("unknown texture")
	}

	if headlessRendering {
		return
	}

	// ### file caching
	image, err := img.LoadTexture(this.sdlRenderer, fileTexture.Source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load source: %s (%s)\n", fileTexture.Source, err.Error())
		return
	}

	defer image.Destroy()
	rect := sdl.Rect{int32(geo.X), int32(geo.Y), int32(geo.Z), int32(geo.W)}
	this.sdlRenderer.Copy(image, nil, &rect)
}

func (this *Window) drawText(node *sg.TextNode, transform sg.Transform) {
	geo := transform.Geometry(sg.Vec4{0, 0, node.Width, node.Height})

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
		rect := sdl.Rect{int32(geo.X), int32(geo.Y), int32(geo.Z), int32(geo.W)}
		if node.Color.X == 1 {
			this.setBlendMode(sdl.BLENDMODE_NONE)
		} else {
			this.setBlendMode(sdl.BLENDMODE_BLEND)
		}
		this.sdlRenderer.Copy(texture, nil, &rect)
	}
}
*/

func (this *Window) setBlendMode(bm sdl.BlendMode) {
	// this is cheaper than hitting cgo
	if this.blendMode != bm {
		this.sdlRenderer.SetDrawBlendMode(bm)
		this.blendMode = bm
	}
}

func (this *Window) drawNode(node sg.Node, transform sg.Mat4) {
	if renderDebug {
		log.Printf("drawing node %s: %+v transform:[%+v]", sg.NodeName(node), node, transform)
	}
	switch cnode := node.(type) {
	case sg.SimpleRectangleNode:
		this.drawRectangle(cnode, transform)
	case sg.TransformNode:
	case sg.InputNode:
	//case *sg.ImageNode:
	//	this.drawImage(node, draw.Transform)
	//case *sg.TextNode:
	//	this.drawText(node, draw.Transform)

	default:
		panic(fmt.Sprintf("unknown drawable %T %+v", node, node))
	}
}
