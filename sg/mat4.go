package sg

import "fmt"

type Mat4Type int

const (
	IdentityType         Mat4Type = 0x00
	Translation2DType             = 0x01
	Scale2DType                   = 0x02
	Rotation2DType                = 0x04
	ScaleAndRotate2DType          = 0x07 // all of the above
	GenericType                   = 0xff
)

type Mat4 struct {
	M    [16]float32
	Type Mat4Type
}

func NewIdentity() Mat4 {
	return Mat4{
		M: [16]float32{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
		Type: IdentityType,
	}
}

func NewMat4(
	m11 float32, m12 float32, m13 float32, m14 float32,
	m21 float32, m22 float32, m23 float32, m24 float32,
	m31 float32, m32 float32, m33 float32, m34 float32,
	m41 float32, m42 float32, m43 float32, m44 float32,
	mtype Mat4Type) Mat4 {
	return Mat4{
		M: [16]float32{
			m11, m12, m13, m14,
			m21, m22, m23, m24,
			m31, m32, m33, m34,
			m41, m42, m43, m44,
		},
		Type: mtype,
	}
}

func (this Mat4) Equals(other Mat4) bool {
	for i, v := range this.M {
		if v != other.M[i] {
			return false
		}
	}

	return true
}

func (this Mat4) MulM4(o Mat4) Mat4 {
	m := this.M
	if this.Type == Translation2DType && o.Type == Translation2DType {
		return NewMat4(1, 0, 0, m[3]+o.M[3],
			0, 1, 0, m[7]+o.M[7],
			0, 0, 1, 0,
			0, 0, 0, 1,
			Translation2DType)

	} else if this.Type == Translation2DType {
		return NewMat4(
			o.M[0]+m[3]*o.M[12],
			o.M[1]+m[3]*o.M[13],
			o.M[2]+m[3]*o.M[14],
			o.M[3]+m[3]*o.M[15],
			o.M[4]+m[7]*o.M[12],
			o.M[5]+m[7]*o.M[13],
			o.M[6]+m[7]*o.M[14],
			o.M[7]+m[7]*o.M[15],
			o.M[8],
			o.M[9],
			o.M[10],
			o.M[11],
			o.M[12],
			o.M[13],
			o.M[14],
			o.M[15],
			this.Type|o.Type)

	} else if o.Type == Translation2DType {
		return NewMat4(
			m[0], m[1], m[2], m[0]*o.M[3]+m[1]*o.M[7]+m[3],
			m[4], m[5], m[6], m[4]*o.M[3]+m[5]*o.M[7]+m[7],
			m[8], m[9], m[10], m[8]*o.M[3]+m[9]*o.M[7]+m[11],
			m[12], m[13], m[14], m[12]*o.M[3]+m[13]*o.M[7]+m[15],
			Mat4Type(this.Type|o.Type))

	} else if this.Type <= ScaleAndRotate2DType && o.Type <= ScaleAndRotate2DType {
		return NewMat4(
			m[0]*o.M[0]+m[1]*o.M[4],
			m[0]*o.M[1]+m[1]*o.M[5],
			0,
			m[0]*o.M[3]+m[1]*o.M[7]+m[3],

			m[4]*o.M[0]+m[5]*o.M[4],
			m[4]*o.M[1]+m[5]*o.M[5],
			0,
			m[4]*o.M[3]+m[5]*o.M[7]+m[7],

			0, 0, 1, 0,
			0, 0, 0, 1,
			Mat4Type(this.Type|o.Type))
	}

	// Genereic full multiplication
	return NewMat4(
		m[0]*o.M[0]+m[1]*o.M[4]+m[2]*o.M[8]+m[3]*o.M[12],
		m[0]*o.M[1]+m[1]*o.M[5]+m[2]*o.M[9]+m[3]*o.M[13],
		m[0]*o.M[2]+m[1]*o.M[6]+m[2]*o.M[10]+m[3]*o.M[14],
		m[0]*o.M[3]+m[1]*o.M[7]+m[2]*o.M[11]+m[3]*o.M[15],

		m[4]*o.M[0]+m[5]*o.M[4]+m[6]*o.M[8]+m[7]*o.M[12],
		m[4]*o.M[1]+m[5]*o.M[5]+m[6]*o.M[9]+m[7]*o.M[13],
		m[4]*o.M[2]+m[5]*o.M[6]+m[6]*o.M[10]+m[7]*o.M[14],
		m[4]*o.M[3]+m[5]*o.M[7]+m[6]*o.M[11]+m[7]*o.M[15],

		m[8]*o.M[0]+m[9]*o.M[4]+m[10]*o.M[8]+m[11]*o.M[12],
		m[8]*o.M[1]+m[9]*o.M[5]+m[10]*o.M[9]+m[11]*o.M[13],
		m[8]*o.M[2]+m[9]*o.M[6]+m[10]*o.M[10]+m[11]*o.M[14],
		m[8]*o.M[3]+m[9]*o.M[7]+m[10]*o.M[11]+m[11]*o.M[15],

		m[12]*o.M[0]+m[13]*o.M[4]+m[14]*o.M[8]+m[15]*o.M[12],
		m[12]*o.M[1]+m[13]*o.M[5]+m[14]*o.M[9]+m[15]*o.M[13],
		m[12]*o.M[2]+m[13]*o.M[6]+m[14]*o.M[10]+m[15]*o.M[14],
		m[12]*o.M[3]+m[13]*o.M[7]+m[14]*o.M[11]+m[15]*o.M[15],
		Mat4Type(this.Type|o.Type))
}

func (this Mat4) MulV2(v Vec2) Vec2 {
	m := this.M
	return Vec2{
		m[0]*v.X + m[1]*v.Y + m[3],
		m[4]*v.X + m[5]*v.Y + m[7]}
}
func (this Mat4) MulV3(v Vec3) Vec3 {
	m := this.M
	return Vec3{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3],
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7],
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]}
}
func (this Mat4) MulV4(v Vec4) Vec4 {
	m := this.M
	return Vec4{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W,
		m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W}
}
func (this Mat4) Transposed() Mat4 {
	m := this.M
	return NewMat4(
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[12],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
		IdentityType)
}
func (this Mat4) Inverted(invertible *bool) Mat4 {
	m := this.M
	var inv Mat4
	inv.M[0] = m[5]*m[10]*m[15] -
		m[5]*m[11]*m[14] -
		m[9]*m[6]*m[15] +
		m[9]*m[7]*m[14] +
		m[13]*m[6]*m[11] -
		m[13]*m[7]*m[10]
	inv.M[4] = -m[4]*m[10]*m[15] +
		m[4]*m[11]*m[14] +
		m[8]*m[6]*m[15] -
		m[8]*m[7]*m[14] -
		m[12]*m[6]*m[11] +
		m[12]*m[7]*m[10]
	inv.M[8] = m[4]*m[9]*m[15] -
		m[4]*m[11]*m[13] -
		m[8]*m[5]*m[15] +
		m[8]*m[7]*m[13] +
		m[12]*m[5]*m[11] -
		m[12]*m[7]*m[9]
	inv.M[12] = -m[4]*m[9]*m[14] +
		m[4]*m[10]*m[13] +
		m[8]*m[5]*m[14] -
		m[8]*m[6]*m[13] -
		m[12]*m[5]*m[10] +
		m[12]*m[6]*m[9]
	inv.M[1] = -m[1]*m[10]*m[15] +
		m[1]*m[11]*m[14] +
		m[9]*m[2]*m[15] -
		m[9]*m[3]*m[14] -
		m[13]*m[2]*m[11] +
		m[13]*m[3]*m[10]
	inv.M[5] = m[0]*m[10]*m[15] -
		m[0]*m[11]*m[14] -
		m[8]*m[2]*m[15] +
		m[8]*m[3]*m[14] +
		m[12]*m[2]*m[11] -
		m[12]*m[3]*m[10]
	inv.M[9] = -m[0]*m[9]*m[15] +
		m[0]*m[11]*m[13] +
		m[8]*m[1]*m[15] -
		m[8]*m[3]*m[13] -
		m[12]*m[1]*m[11] +
		m[12]*m[3]*m[9]
	inv.M[13] = m[0]*m[9]*m[14] -
		m[0]*m[10]*m[13] -
		m[8]*m[1]*m[14] +
		m[8]*m[2]*m[13] +
		m[12]*m[1]*m[10] -
		m[12]*m[2]*m[9]
	inv.M[2] = m[1]*m[6]*m[15] -
		m[1]*m[7]*m[14] -
		m[5]*m[2]*m[15] +
		m[5]*m[3]*m[14] +
		m[13]*m[2]*m[7] -
		m[13]*m[3]*m[6]
	inv.M[6] = -m[0]*m[6]*m[15] +
		m[0]*m[7]*m[14] +
		m[4]*m[2]*m[15] -
		m[4]*m[3]*m[14] -
		m[12]*m[2]*m[7] +
		m[12]*m[3]*m[6]
	inv.M[10] = m[0]*m[5]*m[15] -
		m[0]*m[7]*m[13] -
		m[4]*m[1]*m[15] +
		m[4]*m[3]*m[13] +
		m[12]*m[1]*m[7] -
		m[12]*m[3]*m[5]
	inv.M[14] = -m[0]*m[5]*m[14] +
		m[0]*m[6]*m[13] +
		m[4]*m[1]*m[14] -
		m[4]*m[2]*m[13] -
		m[12]*m[1]*m[6] +
		m[12]*m[2]*m[5]
	inv.M[3] = -m[1]*m[6]*m[11] +
		m[1]*m[7]*m[10] +
		m[5]*m[2]*m[11] -
		m[5]*m[3]*m[10] -
		m[9]*m[2]*m[7] +
		m[9]*m[3]*m[6]
	inv.M[7] = m[0]*m[6]*m[11] -
		m[0]*m[7]*m[10] -
		m[4]*m[2]*m[11] +
		m[4]*m[3]*m[10] +
		m[8]*m[2]*m[7] -
		m[8]*m[3]*m[6]
	inv.M[11] = -m[0]*m[5]*m[11] +
		m[0]*m[7]*m[9] +
		m[4]*m[1]*m[11] -
		m[4]*m[3]*m[9] -
		m[8]*m[1]*m[7] +
		m[8]*m[3]*m[5]
	inv.M[15] = m[0]*m[5]*m[10] -
		m[0]*m[6]*m[9] -
		m[4]*m[1]*m[10] +
		m[4]*m[2]*m[9] +
		m[8]*m[1]*m[6] -
		m[8]*m[2]*m[5]

	var det float32 = m[0]*inv.M[0] + m[1]*inv.M[4] + m[2]*inv.M[8] + m[3]*inv.M[12]

	if det == 0 {
		if invertible != nil {
			*invertible = false
		}
		return NewIdentity()
	}

	det = 1.0 / det
	for i := 0; i < 16; i++ {
		inv.M[i] *= det
	}

	if invertible != nil {
		*invertible = true
	}

	return inv
}
func (this Mat4) String() string {
	return fmt.Sprintf("%g %g %g %g\n%g %g %g %g\n%g %g %g %g\n%g %g %g %g",
		this.M[0], this.M[1], this.M[2], this.M[3],
		this.M[4], this.M[5], this.M[6], this.M[7],
		this.M[8], this.M[9], this.M[10], this.M[11],
		this.M[12], this.M[13], this.M[14], this.M[15])
}
