package animation

import (
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

type Vec4Animation struct {
	From     sg.Vec4
	To       sg.Vec4
	Duration time.Duration // how long transition takes
	Easing   EasingFunc

	initialized bool // initialized?
	aAnim       FloatAnimation
	rAnim       FloatAnimation
	gAnim       FloatAnimation
	bAnim       FloatAnimation
	currentVec  sg.Vec4
}

func (this *Vec4Animation) Advance(frameTime time.Duration) {
	if !this.initialized {
		this.initialized = true
		this.Restart()
	}

	this.aAnim.Advance(frameTime)
	this.rAnim.Advance(frameTime)
	this.gAnim.Advance(frameTime)
	this.bAnim.Advance(frameTime)
	this.currentVec = sg.Vec4{this.aAnim.Get(), this.rAnim.Get(), this.gAnim.Get(), this.bAnim.Get()}
}

func (this *Vec4Animation) Get() sg.Vec4 {
	return this.currentVec
}

func (this *Vec4Animation) Restart() {
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

	this.aAnim.Easing = this.Easing
	this.rAnim.Easing = this.Easing
	this.gAnim.Easing = this.Easing
	this.bAnim.Easing = this.Easing

	this.currentVec = this.From

	this.aAnim.Restart()
	this.rAnim.Restart()
	this.gAnim.Restart()
	this.bAnim.Restart()
}
