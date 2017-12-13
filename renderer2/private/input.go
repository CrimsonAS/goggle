package private

import (
	"log"

	"github.com/CrimsonAS/goggle/sg"
	"github.com/CrimsonAS/goggle/sg2"
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
func (this *InputHelper) ProcessPointerEvents(transform sg.Mat4, ts *sg2.TouchState) bool {
	handledEvents := false

	// Translate mouse position to node coordinates
	tp := this.MousePos.Sub(transform.MulV2(sg.Vec2{0, 0}))
	tg := ts.TouchGeometry

	// BUG: ### unsolved problems: we should also probably block propagation of hover.
	// We could have a return code to block hover propagating further down the tree,
	// letting someone write code like:
	//
	// Root UI node
	//     Sidebar PointerEnter() { return true; /* block */ }
	//         Button Hoverable // to highlight as need be
	//     UI page

	if pointInside(tg.X, tg.Y, tg.Z, tg.W, tp) {
		if !ts.IsHovered {
			ts.IsHovered = true
			mouseDebug("Pointer entering: %+v at %s %s", ts, this.MousePos, tp)
			if ts.OnEnter != nil {
				ts.OnEnter(ts)
			}
			handledEvents = true
		}
	} else {
		if ts.IsHovered {
			ts.IsHovered = false
			mouseDebug("Pointer leaving: %+v at %s %s", ts, this.MousePos, tp)
			if ts.OnLeave != nil {
				ts.OnLeave(ts)
			}
			handledEvents = true
		}
	}

	/*
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
		}*/

	return handledEvents
}
