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

package main

import (
	"github.com/CrimsonAS/goggle/renderer/sdlsoftware"
	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/components"
	"github.com/CrimsonAS/goggle/sg/layouts"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

func RowLayout(c sg.Constraints, children []layouts.BoxChild, props interface{}) sg.Size {
	remainingWidth := c.Max.Width
	var x, maxChildY float32

	for i, child := range children {
		childConstraint := sg.Constraints{
			Max: sg.Size{
				Width:  remainingWidth / float32(len(children)-i),
				Height: c.Max.Height,
			},
		}
		childSize := child.Render(childConstraint)
		remainingWidth -= childSize.Width
		child.SetPosition(sg.Position{x, 0})
		x += childSize.Width
		if childSize.Height > maxChildY {
			maxChildY = childSize.Height
		}
	}

	return sg.Size{x, maxChildY}
}

func LayoutWindow(props components.PropType, state *components.RenderState) sg.Node {
	return nodes.Rectangle{
		Color: sg.Color{1, 0, 0, 1},
		Children: []sg.Node{
			layouts.Box{
				Layout: layouts.Flex,
				Props: layouts.FlexLayoutProps{
					Direction: layouts.FlexRow,
				},
				Children: []sg.Node{
					layouts.Box{
						Layout: layouts.Fill,
						ParentProps: layouts.FlexChildProps{
							Basis:  layouts.PixelUnit(400),
							Grow:   2,
							Shrink: 1,
						},
						Child: nodes.Rectangle{
							Color: sg.Color{1, 1, 1, 0},
						},
					},
					layouts.Box{
						Layout: layouts.None,
						ParentProps: layouts.FlexChildProps{
							Basis: layouts.PixelUnit(200),
						},
						Child: components.Component{
							Type: components.Rectangle,
							Props: components.RectangleProps{
								Color: sg.Color{1, 0.5, 0.5, 0},
								Size:  sg.Size{200, 200},
							},
						},
					},
				},
			},
		},
	}
}

func main() {
	r, err := sdlsoftware.NewRenderer()
	if err != nil {
		panic(err)
	}

	w, err := r.CreateWindow()
	defer w.Destroy()

	if err != nil {
		panic(err)
	}

	for r.IsRunning() {
		r.ProcessEvents()
		w.Render(components.Component{Type: LayoutWindow})
	}

	r.Quit()
}
