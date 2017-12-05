package animation

import (
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

type ColorAnimation struct {
	From     sg.Color
	To       sg.Color
	Duration time.Duration
	Easing   EasingFunc

	initialized   bool
	vec4animation Vec4Animation
}

func (this *ColorAnimation) Advance(frameTime time.Duration) {
	if !this.initialized {
		this.initialized = true
		this.Restart()
	}
	this.vec4animation.Advance(frameTime)
}

func (this *ColorAnimation) Get() sg.Color {
	return sg.Color(this.vec4animation.Get())
}

func (this *ColorAnimation) Restart() {
	this.vec4animation.From = sg.Vec4(this.From)
	this.vec4animation.To = sg.Vec4(this.To)
	this.vec4animation.Duration = this.Duration
	this.vec4animation.Easing = this.Easing
	this.vec4animation.Restart()
}
