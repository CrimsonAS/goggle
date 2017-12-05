package sg

import (
	"fmt"
	"strconv"
)

// A R G B
type Color Vec4

func (this Color) String() string {
	if this.X == 0 {
		return fmt.Sprintf("#%s%s%s", strconv.FormatInt(int64(this.Rint()), 16), strconv.FormatInt(int64(this.Gint()), 16), strconv.FormatInt(int64(this.Bint()), 16))
	} else {
		return fmt.Sprintf("#%s%s%s%s", strconv.FormatInt(int64(this.Aint()), 16), strconv.FormatInt(int64(this.Rint()), 16), strconv.FormatInt(int64(this.Gint()), 16), strconv.FormatInt(int64(this.Bint()), 16))
	}
}

func (this Color) A() float32 {
	return this.X
}
func (this Color) Aint() int {
	return int(this.X * 255.0)
}
func (this Color) R() float32 {
	return this.Y
}
func (this Color) Rint() int {
	return int(this.Y * 255.0)
}
func (this Color) G() float32 {
	return this.Z
}
func (this Color) Gint() int {
	return int(this.Z * 255.0)
}
func (this Color) B() float32 {
	return this.W
}
func (this Color) Bint() int {
	return int(this.W * 255.0)
}
