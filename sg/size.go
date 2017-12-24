package sg

import "fmt"

type Size struct {
	Width, Height float32
}

func SizeV2(v Vec2) Size {
	return Size{v.X, v.Y}
}

func (s Size) Min(o Size) Size {
	if o.Width < s.Width {
		s.Width = o.Width
	}
	if o.Height < s.Height {
		s.Height = o.Height
	}
	return s
}

func (s Size) Max(o Size) Size {
	if o.Width > s.Width {
		s.Width = o.Width
	}
	if o.Height > s.Height {
		s.Height = o.Height
	}
	return s
}

func (s Size) IsNil() bool {
	return s.Width == 0 && s.Width == 0
}

func (s Size) IsEmpty() bool {
	return s.Width <= 0 || s.Height <= 0
}

func (s Size) ToPosition() Position {
	return Position{s.Width, s.Height}
}

func (s Size) Sub(o Size) Size {
	return Size{s.Width - o.Width, s.Height - o.Height}
}

func (s Size) SubPosition(p Position) Size {
	return Size{s.Width - p.X, s.Height - p.Y}
}

func (s Size) String() string {
	return fmt.Sprintf("%gx%g", s.Width, s.Height)
}
