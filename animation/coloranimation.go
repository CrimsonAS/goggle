package animation

import (
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

// Animate From a value To another value over a given Duration, and reverse
// direction after that Duration.
type ColorAnimation struct {
	From     sg.Color
	To       sg.Color
	Duration time.Duration // how long transition takes

	initialized  bool // initialized?
	aAnim        FloatAnimation
	rAnim        FloatAnimation
	gAnim        FloatAnimation
	bAnim        FloatAnimation
	currentColor sg.Color
}

func (this *ColorAnimation) Advance(frameTime time.Duration) {
	if !this.initialized {
		this.initialized = true
		this.aAnim.From = this.From.X
		this.rAnim.From = this.From.Y
		this.gAnim.From = this.From.Z
		this.bAnim.From = this.From.W

		this.aAnim.To = this.To.X
		this.rAnim.To = this.To.Y
		this.gAnim.To = this.To.Z
		this.bAnim.To = this.To.W

		this.aAnim.Duration = this.Duration
		this.rAnim.Duration = this.Duration
		this.gAnim.Duration = this.Duration
		this.bAnim.Duration = this.Duration
		this.currentColor = this.From
		return
	}

	this.aAnim.Advance(frameTime)
	this.rAnim.Advance(frameTime)
	this.gAnim.Advance(frameTime)
	this.bAnim.Advance(frameTime)
	this.currentColor = sg.Color{this.aAnim.Get(), this.rAnim.Get(), this.gAnim.Get(), this.bAnim.Get()}
}

func (this *ColorAnimation) Get() sg.Color {
	return this.currentColor
}

func (this *ColorAnimation) Restart() {
	this.aAnim.Restart()
	this.rAnim.Restart()
	this.gAnim.Restart()
	this.bAnim.Restart()
}
