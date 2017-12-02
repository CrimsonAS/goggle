package sg

// A Hoverable is a node that will get events when a point's coordintes are
// above the item.
//
// Note that Hoverable must also implement Sizeable for the
// scenegraph to know that the point is inside the item's boundaries.
type Hoverable interface {
	PointerEnter(Vec2)
	PointerLeave(Vec2)
}

// BUG:
// The first Pressable or Tappable to intercept a point automatically grabs further events on the point.
// This isn't ideal, we should have explicit grabbing for Pressable.

// A Pressable is a node that will get events when a point is pressed or
// released in its boundary.
//
// Note that Pressable must also implement Sizeable for the
// scenegraph to know that the point is inside the item's boundaries.
type Pressable interface {
	PointerPressed(Vec2)
	PointerReleased(Vec2)
}

// A Moveable is a node that will get events when a mouse is inside its boundary.
//
// Note that Moveable must also implement Sizeable for the
// scenegraph to know that the point is inside the item's boundaries.
type Moveable interface {
	PointerMoved(Vec2)
}

// A Tappable is a node that will get events when a touch is pressed and released.
//
// Note that Moveable must also implement Sizeable for the
// scenegraph to know that the point is inside the item's boundaries.
type Tappable interface {
	PointerTapped(Vec2)
}
