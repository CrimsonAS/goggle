package sg

import (
	"fmt"
	"strconv"
)

// A R G B
type Color Vec4

func twoDigitHex(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%s", strconv.FormatInt(int64(n), 16))
	} else {
		return fmt.Sprintf("%s", strconv.FormatInt(int64(n), 16))
	}
}

func (this Color) String() string {
	return fmt.Sprintf("#%s%s%s%s",
		twoDigitHex(this.Aint()),
		twoDigitHex(this.Rint()),
		twoDigitHex(this.Gint()),
		twoDigitHex(this.Bint()))
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
