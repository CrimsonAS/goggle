package sg

type Transform struct {
	Translate Vec2
	Scale     float32
	// ### Matrix
}

var IdentityTransform = Transform{
	Translate: Vec2{0, 0},
	Scale:     1,
}

func (t Transform) X() float32 {
	return t.Translate.X
}

func (t Transform) Y() float32 {
	return t.Translate.Y
}

func (t Transform) Pos(pos Vec2) Vec2 {
	return pos.Add(t.Translate)
}

func (t Transform) Size(size Vec2) Vec2 {
	return size.Mul(Vec2{t.Scale, t.Scale})
}

func (t Transform) Geometry(geo Vec4) Vec4 {
	pos := t.Pos(geo.XY())
	size := t.Size(geo.ZW())
	return Vec4{pos.X, pos.Y, size.X, size.Y}
}

func (t Transform) Mat4() Mat4 {
	mat := Translate2DV2(t.Translate)
	if t.Scale != 1 {
		mat = mat.MulM4(Scale2D(t.Scale, t.Scale))
	}
	return mat
}
