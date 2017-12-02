package private

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
)

type InputHelper struct {
	hoveredNodes    map[sg.Node]bool
	oldHoveredNodes map[sg.Node]bool
	MousePos        sg.TouchPoint
	ButtonUp        bool
	ButtonDown      bool
	MouseGrabber    sg.Node
}

func NewInputHelper() InputHelper {
	return InputHelper{
		hoveredNodes:    make(map[sg.Node]bool),
		oldHoveredNodes: make(map[sg.Node]bool),
		MousePos:        sg.TouchPoint{-1, -1},
	}
}

func mouseDebug(fstr string, vals ...interface{}) {
	const debug = true

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

func (this *InputHelper) DispatchMouseMove() {
	if this.MouseGrabber != nil {
		if moveable, ok := this.MouseGrabber.(sg.Moveable); ok {
			mouseMoveDebug("Pointer moved over %+v at %s", this.MouseGrabber, this.MousePos)
			moveable.PointerMoved(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
		}
	}
}

func (this *InputHelper) ResetFrameState() {
	this.oldHoveredNodes = this.hoveredNodes
	this.hoveredNodes = make(map[sg.Node]bool)

	this.ButtonDown = false
	this.ButtonUp = false
}

func pointInside(x, y, w, h float32, tp sg.TouchPoint) bool {
	return (tp.X >= x && tp.X <= x+w) && (tp.Y >= y && tp.Y <= y+h)
}

// Process pointer events for an item.
// ### should scale/rotate affect input events? i'd say yes, personally.
func (this *InputHelper) ProcessPointerEvents(originX, originY, childWidth, childHeight float32, item sg.Node) {
	// BUG: ### unsolved problems: we should also probably block propagation of hover.
	// We could have a return code to block hover propagating further down the tree,
	// letting someone write code like:
	//
	// Root UI node
	//     Sidebar PointerEnter() { return true; /* block */ }
	//         Button Hoverable // to highlight as need be
	//     UI page
	if hoverable, ok := item.(sg.Hoverable); ok {
		if pointInside(originX, originY, childWidth, childHeight, this.MousePos) {
			this.hoveredNodes[item] = true
			if _, ok = this.oldHoveredNodes[item]; !ok {
				mouseDebug("Pointer entering: %+v at %s", hoverable, this.MousePos)
				hoverable.PointerEnter(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
			}
		} else if _, ok = this.oldHoveredNodes[item]; ok {
			mouseDebug("Pointer leaving: %+v at %s", hoverable, this.MousePos)
			hoverable.PointerLeave(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
		}
	}
	if this.ButtonDown || this.ButtonUp {
		if pressable, ok := item.(sg.Pressable); ok {
			if this.ButtonDown {
				if this.MouseGrabber == nil {
					if pointInside(originX, originY, childWidth, childHeight, this.MousePos) {
						this.MouseGrabber = item
						mouseDebug("Pointer pressed (and grabbed): %+v at %s", pressable, this.MousePos)
						pressable.PointerPressed(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
					}
				}
			} else if this.ButtonUp {
				if this.MouseGrabber == item {
					mouseDebug("Pointer released (ungrabbed): %+v at %s", pressable, this.MousePos)
					pressable.PointerReleased(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
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
					if pointInside(originX, originY, childWidth, childHeight, this.MousePos) {
						mouseDebug("Tappable released (ungrabbed): %+v at %s", tappable, this.MousePos)
						tappable.PointerTapped(sg.TouchPoint{this.MousePos.X, this.MousePos.Y})
					}
				}
			}
		}
		if this.ButtonUp && this.MouseGrabber == item {
			this.MouseGrabber = nil
		}
	}
}
