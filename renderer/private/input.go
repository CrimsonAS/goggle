package private

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg/nodes"
)

type InputHelper struct {
	hoveredNodes    map[sg.Node]bool
	oldHoveredNodes map[sg.Node]bool
	MousePos        sg.Position
	ButtonUp        bool
	ButtonDown      bool

	grabNodeState *nodes.InputState
	grabNodeSeen  bool
}

func NewInputHelper() InputHelper {
	return InputHelper{
		hoveredNodes:    make(map[sg.Node]bool),
		oldHoveredNodes: make(map[sg.Node]bool),
		MousePos:        sg.Position{-1, -1},
	}
}

func mouseDebug(fstr string, vals ...interface{}) {
	const debug = false

	if debug {
		log.Printf(fstr, vals...)
	}
}
func mouseMoveDebug(fstr string, vals ...interface{}) {
	const debug = false

	if debug {
		log.Printf(fstr, vals...)
	}
}

func (this *InputHelper) ResetFrameState() {
	this.oldHoveredNodes = this.hoveredNodes
	this.hoveredNodes = make(map[sg.Node]bool)

	this.ButtonDown = false
	this.ButtonUp = false
}

// Process pointer events for an item.
func (this *InputHelper) ProcessPointerEvents(in *nodes.Input, transform sg.Mat4, sz sg.Size, state *nodes.InputState) bool {
	// ### This is wrong for non-trivial transforms, but I don't want to mess with SDL
	// enough to draw complex shapes for now.
	geom := sg.Geometry{Size: sz}.TransformedBounds(transform)

	handledEvents := false

	// Copy previous state for comparison
	oldState := *state

	containsMouse := geom.Contains(sg.Position{this.MousePos.X, this.MousePos.Y})
	// Calculate relative position and store in the InputState.
	// ### Oh boy this is not fast.. there must be a neater way, or maybe we'll just
	// need some shortcuts for simpler transforms.
	var garbage bool
	state.SceneMousePos = this.MousePos
	state.MousePos = sg.PositionV2(transform.Inverted(&garbage).MulV2(this.MousePos.Vec2()))

	// BUG: ### unsolved problems: we should also probably block propagation of hover.
	// We could have a return code to block hover propagating further down the tree,
	// letting someone write code like:
	//
	// Root UI node
	//     Sidebar PointerEnter() { return true; /* block */ }
	//         Button Hoverable // to highlight as need be
	//     UI page

	if containsMouse {
		if !state.IsHovered {
			mouseDebug("Pointer entering: %+v at %s %s", in, this.MousePos, geom)
			state.IsHovered = true
			if in.OnEnter != nil {
				in.OnEnter(*state)
				handledEvents = true
			}
		}
	} else {
		if state.IsHovered {
			mouseDebug("Pointer leaving: %+v at %s %s", in, this.MousePos, geom)
			state.IsHovered = false
			if in.OnLeave != nil {
				in.OnLeave(*state)
				handledEvents = true
			}
		}
	}

	if state.IsGrabbed {
		if in.OnMove != nil && state.MousePos != oldState.MousePos {
			in.OnMove(*state)
			handledEvents = true
		}
	}

	if containsMouse && this.ButtonDown && this.grabNodeState == nil {
		if in.OnPress != nil {
			mouseDebug("Pointer pressed (and grabbed): %+v at %s %s", in, this.MousePos, geom)
			this.grabNodeState = state
			state.IsGrabbed, state.IsPressed = true, true
			in.OnPress(*state)
			handledEvents = true
		}
	} else if this.ButtonUp && state.IsPressed {
		mouseDebug("Pointer released (ungrabbed): %+v at %s %s", in, this.MousePos, geom)
		if state.IsGrabbed {
			this.grabNodeState = nil
		}
		state.IsGrabbed, state.IsPressed = false, false
		if in.OnRelease != nil {
			in.OnRelease(*state)
			handledEvents = true
		}
	}

	// Flag when the grabNode is seen during event processing so it won't be cleaned
	// up during post-processing. Also, sync up the IsGrabbed flag here for paranoia.
	if state == this.grabNodeState {
		this.grabNodeSeen = true
		state.IsGrabbed = true
	} else {
		state.IsGrabbed = false
	}

	return handledEvents
}

func (this *InputHelper) EndPointerEvents() {
	if this.grabNodeState != nil && !this.grabNodeSeen {
		mouseDebug("Grab node disappeared, ungrabbing: %+v", this.grabNodeState)
		this.grabNodeState = nil
		// ### Other cleanup, handling of ongoing events, ..?
	}
	// Reset for next pass
	this.grabNodeSeen = false
}
