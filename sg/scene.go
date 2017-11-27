package main

import (
	"fmt"
)

// A Scene is a controller containing a tree of Nodeable and Renderable, as well
// as the logic to tie that and the Renderer together.
type Scene struct {
	// A list of Renderables  that need updating on the next sync pass.
	dirtyList []Renderable
}

// Mark this renderable as dirty. It will be cleaned on the next sync pass.
func (this *Scene) MarkDirty(renderable Renderable) {
	fmt.Printf("Marking dirty %p\n", renderable)
	this.dirtyList = append(this.dirtyList, renderable)
}

// Renders a single Renderable, and its children (if it is also a Nodeable).
func (this *Scene) renderItem(item Renderable) TreeNode {
	fmt.Printf("Rendering %p %+v\n", item, item)
	node := item.Render()
	fmt.Printf("Got node %p %+v\n", node, node)

	// If the child is also a tree item, render those.
	if treeChild, ok := item.(Nodeable); ok {
		for _, citem := range treeChild.GetChildren() {
			if renderableChild, ok := citem.(Renderable); ok {
				this.renderItem(renderableChild)

				// ### and append the nodes of the children to that of the
				// parent, right?
			}
		}
	}

	return node
}

// Prepare the scene for rendering.
func (this *Scene) Sync() {
	// For each dirty item, find out what it wants to render
	for _, item := range this.dirtyList {
		// ### cache the subtree of this renderItem calls for use by a
		// renderer..?
		this.renderItem(item)

		// ### and now actually render the content to screen somehow.
	}
}
