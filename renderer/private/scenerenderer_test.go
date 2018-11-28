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

package private

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

type fakeWindow struct {
}

var _ sg.Windowable = (*fakeWindow)(nil)

func (this *fakeWindow) GetSize() sg.Vec2 {
	return sg.Vec2{10, 10}
}

func (this *fakeWindow) FrameTime() time.Duration {
	return 1
}

type rendererTestTypeTree struct {
	typeName string
	children []rendererTestTypeTree
}

type rendererTest struct {
	root      sg.Node
	drawables []DrawableNode
}

func TestSingleDrawable(t *testing.T) {
	root := &sg.RectangleNode{
		X:      0,
		Y:      0,
		Width:  100,
		Height: 100,
		Color:  sg.Color{1, 1, 0, 0},
	}
	rt := rendererTest{
		root: root,
		drawables: []DrawableNode{
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     1.0,
				},
				Node: root,
			},
		},
	}
	renderTest(t, rt)
}

func TestSingleDrawableWithChild(t *testing.T) {
	child := &sg.RectangleNode{
		X:      0,
		Y:      0,
		Width:  10,
		Height: 10,
		Color:  sg.Color{0, 0, 1, 1},
	}
	root := &sg.RectangleNode{
		X:      0,
		Y:      0,
		Width:  100,
		Height: 100,
		Color:  sg.Color{1, 1, 0, 0},
		Children: []sg.Node{
			child,
		},
	}
	rt := rendererTest{
		root: root,
		drawables: []DrawableNode{
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     1.0,
				},
				Node: root,
			},
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     1.0,
				},
				Node: child,
			},
		},
	}
	renderTest(t, rt)
}

func TestTranslationPropagation(t *testing.T) {
	child2 := &sg.RectangleNode{
		X:      30,
		Y:      30,
		Width:  10,
		Height: 10,
		Color:  sg.Color{0, 0, 1, 1},
	}
	child := &sg.RectangleNode{
		X:        10,
		Y:        10,
		Width:    10,
		Height:   10,
		Color:    sg.Color{0, 0, 1, 1},
		Children: []sg.Node{child2},
	}
	root := &sg.RectangleNode{
		X:      10,
		Y:      10,
		Width:  100,
		Height: 100,
		Color:  sg.Color{1, 1, 0, 0},
		Children: []sg.Node{
			child,
		},
	}
	rt := rendererTest{
		root: root,
		drawables: []DrawableNode{
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{10, 10},
					Scale:     1.0,
				},
				Node: root,
			},
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{20, 20},
					Scale:     1.0,
				},
				Node: child,
			},
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{50, 50},
					Scale:     1.0,
				},
				Node: child2,
			},
		},
	}
	renderTest(t, rt)
}

func TestScalePropagation(t *testing.T) {
	child2 := &sg.RectangleNode{
		X:      0,
		Y:      0,
		Width:  10,
		Height: 10,
		Color:  sg.Color{0, 0, 1, 1},
	}
	scale2 := &sg.ScaleNode{
		Scale:    2.0,
		Children: []sg.Node{child2},
	}
	child := &sg.RectangleNode{
		X:        0,
		Y:        0,
		Width:    10,
		Height:   10,
		Color:    sg.Color{0, 0, 1, 1},
		Children: []sg.Node{scale2},
	}
	scale := &sg.ScaleNode{
		Scale:    2.0,
		Children: []sg.Node{child},
	}
	root := &sg.RectangleNode{
		X:      0,
		Y:      0,
		Width:  100,
		Height: 100,
		Color:  sg.Color{1, 1, 0, 0},
		Children: []sg.Node{
			scale,
		},
	}
	rt := rendererTest{
		root: root,
		drawables: []DrawableNode{
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     1.0,
				},
				Node: root,
			},
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     2.0,
				},
				Node: child,
			},
			DrawableNode{
				Transform: sg.Transform{
					Translate: sg.Vec2{0, 0},
					Scale:     4.0,
				},
				Node: child2,
			},
		},
	}
	renderTest(t, rt)
}

func renderTest(t *testing.T, rt rendererTest) {
	fw := &fakeWindow{}
	ih := &InputHelper{}
	sr1 := SceneRenderer{
		Window:          fw,
		InputHelper:     ih,
		DisableParallel: true,
	}
	drawables1 := sr1.Render(rt.root)

	sr2 := SceneRenderer{
		Window:          fw,
		InputHelper:     ih,
		DisableParallel: false,
	}
	drawables2 := sr2.Render(rt.root)

	if len(drawables1) != len(drawables2) {
		t.Fatalf("Drawables differed between parallel and not (%d and %d)", len(drawables1), len(drawables2))
	}

	if !reflect.DeepEqual(drawables1, rt.drawables) {
		log.Printf("Wanted: %+v", rt.drawables)
		log.Printf("Got: %+v", drawables1)
		t.Fatalf("Parallel run did not produce expected drawables.")
	}
	if !reflect.DeepEqual(drawables2, rt.drawables) {
		log.Printf("Wanted: %+v", rt.drawables)
		log.Printf("Got: %+v", drawables2)
		t.Fatalf("Serial run did not produce expected drawables")
	}
}
