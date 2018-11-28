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

package layouts

import "github.com/CrimsonAS/goggle/sg"

type Box struct {
	// Layout is a layout function for the box; see BoxFunc for details
	Layout BoxFunc
	// Props is an opaque type for passing properties to the Layout function
	Props interface{}
	// ParentProps is an opaque type for passing properties to a parent Box during layout
	ParentProps interface{}
	// Children defines child nodes of the Box
	Children []sg.Node
	// Child is a convenient alternative to Children for boxes with only one child
	Child sg.Node

	// Transform is a transformation matrix applied for this box and all of its children.
	// It could be moved to dynamically determine from Layout at some point, but for now
	// this seems fine.
	//
	// It's critical that transforms are part of Box to preserve the mandate that Box is
	// the only source of geometry in the tree; otherwise, it becomes complex to track
	// changes to transforms in the descendant tree of a box to appropriately transform
	// positions and sizes.
	//
	// As a special case, zero-value transforms will be ignored. It is not necessary to
	// initialize Transform to an identity matrix when not applying transformations.
	Transform sg.Mat4
}

// BoxChild is used as a parameter to BoxFunc.
type BoxChild interface {
	// Props is taken from the ParentProps of the child Box. These are properties for
	// the parent Box's layout, of an undefined layout-specific type.
	Props() interface{}
	// Render resolves and executes layout for this child Box with the given constraints,
	// and returns the actual new size of the child Box.
	//
	// Render may be called multiple times during a layout if constraints change, although
	// this is arbitrarily expensive and should be avoided when possible. If render is called
	// more than once, the resulting scene will be identical to rendering with only the final
	// constraints. State is not carried between multiple render passes within a layout.
	//
	// If render is not called for a child Box, the Box and its descendants are omitted
	// from the scene and have no render cost. This is not a way to control visibility for
	// child boxes; it is a way to make them temporarily not exist. If a Box is not rendered
	// in a scene, state will be lost for that tree.
	Render(c sg.Constraints) sg.Size
	// Geometry returns the child's geometry, as determined by prior calls to Render and
	// SetPosition. The origin and size are undefined if Render or SetPosition hasn't been
	// called, respectively. This is provided as a convenience for layouts that need to refer
	// to child geometry for multiple passes.
	Geometry() sg.Geometry
	// SetPosition sets the top-left position of the child box in relative coordinates.
	SetPosition(pos sg.Position)
}

type BoxFunc func(c sg.Constraints, children []BoxChild, props interface{}) sg.Size

var _ sg.Parentable = Box{}

func (b Box) GetChildren() []sg.Node {
	if b.Child != nil {
		return append([]sg.Node{b.Child}, b.Children...)
	} else {
		return b.Children
	}
}
