package private

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
)

type InputHelper struct {
	hoveredNodes    map[sg.Node]bool
	oldHoveredNodes map[sg.Node]bool
	MousePos        sg.Vec2
	ButtonUp        bool
	ButtonDown      bool
	MouseGrabber    sg.Node
}

func NewInputHelper() InputHelper {
	return InputHelper{
		hoveredNodes:    make(map[sg.Node]bool),
		oldHoveredNodes: make(map[sg.Node]bool),
		MousePos:        sg.Vec2{-1, -1},
	}
}

func mouseDebug(fstr string, vals ...interface{}) {
	const debug = true

	if debug {
		log.Printf(fstr, vals...)
	}
}
func mouseMoveDebug(fstr string, vals ...interface{}) {
	const debug = true

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

func pointInside(x, y, w, h float32, tp sg.Vec2) bool {
	return (tp.X >= x && tp.X <= x+w) && (tp.Y >= y && tp.Y <= y+h)
}

func NodeAcceptsInputEvents(node sg.Node) bool {
	switch node.(type) {
	case sg.Hoverable:
		return true
	case sg.Moveable:
		return true
	case sg.Pressable:
		return true
	case sg.Tappable:
		return true
	default:
		return false
	}
}

// Process pointer events for an item.
// ### should scale/rotate affect input events? i'd say yes, personally.
func (this *InputHelper) ProcessPointerEvents(origin sg.Vec2, childWidth, childHeight float32, item sg.Node) bool {
	handledEvents := false

	// BUG: ### unsolved problems: we should also probably block propagation of hover.
	// We could have a return code to block hover propagating further down the tree,
	// letting someone write code like:
	//
	// Root UI node
	//     Sidebar PointerEnter() { return true; /* block */ }
	//         Button Hoverable // to highlight as need be
	//     UI page
	tp := sg.Vec2{X: this.MousePos.X /*- origin.X*/, Y: this.MousePos.Y /*- origin.Y*/}
	if hoverable, ok := item.(sg.Hoverable); ok {
		if pointInside(origin.X, origin.Y, childWidth, childHeight, this.MousePos) {
			this.hoveredNodes[item] = true
			if _, ok = this.oldHoveredNodes[item]; !ok {
				mouseDebug("Pointer entering: %+v at %s %s", hoverable, this.MousePos, tp)
				hoverable.PointerEnter(tp)
				handledEvents = true
			}
		} else if _, ok = this.oldHoveredNodes[item]; ok {
			mouseDebug("Pointer leaving: %+v at %s %s", hoverable, this.MousePos, tp)
			hoverable.PointerLeave(tp)
			handledEvents = true
		}
	}

	// BUG: we should only deliver this if the tp is not the same as the last PointerMoved, I think.
	if this.MouseGrabber != nil {
		if moveable, ok := this.MouseGrabber.(sg.Moveable); ok {
			mouseMoveDebug("Pointer moved over %+v at %s %s", this.MouseGrabber, this.MousePos, tp)
			moveable.PointerMoved(tp)
			handledEvents = true
		}
	}

	if this.ButtonDown || this.ButtonUp {
		if pressable, ok := item.(sg.Pressable); ok {
			if this.ButtonDown {
				if this.MouseGrabber == nil {
					if pointInside(origin.X, origin.Y, childWidth, childHeight, this.MousePos) {
						this.MouseGrabber = item
						mouseDebug("Pointer pressed (and grabbed): %+v at %s %s", pressable, this.MousePos, tp)
						pressable.PointerPressed(tp)
						handledEvents = true
					}
				}
			} else if this.ButtonUp {
				if this.MouseGrabber == item {
					mouseDebug("Pointer released (ungrabbed): %+v at %s %s", pressable, this.MousePos, tp)
					pressable.PointerReleased(tp)
					handledEvents = true
				}
			}
		}
		if tappable, ok := item.(sg.Tappable); ok {
			if this.ButtonDown {
				if this.MouseGrabber == nil {
					// a Tappable takes an implicit grab
					mouseDebug("Tappable pressed (grabbed): %+v at %s", tappable, this.MousePos)
					this.MouseGrabber = item
				}
			} else if this.ButtonUp {
				if this.MouseGrabber == item {
					if pointInside(origin.X, origin.Y, childWidth, childHeight, this.MousePos) {
						mouseDebug("Tappable released (ungrabbed): %+v at %s", tappable, this.MousePos)
						tappable.PointerTapped(tp)
						handledEvents = true
					}
				}
			}
		}
		if this.ButtonUp && this.MouseGrabber == item {
			this.MouseGrabber = nil
		}
	}

	return handledEvents
}
