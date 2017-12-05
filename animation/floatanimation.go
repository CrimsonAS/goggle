package animation

import (
	"time"
)

type EasingFunc func(float64) float64

// Animate From a value To another value over a given Duration, and reverse
// direction after that Duration.
type FloatAnimation struct {
	From     float32       // from value
	To       float32       // to value
	Duration time.Duration // how long transition takes
	Easing   EasingFunc

	initialized       bool // initialized?
	goingDown         bool
	remainingDuration time.Duration // remaining Duration left before swap
	currentValue      float32
}

func (this *FloatAnimation) Advance(frameTime time.Duration) {
	if !this.initialized {
		this.initialized = true
		this.Restart()
	}

	if this.goingDown {
		this.remainingDuration -= frameTime
		if this.remainingDuration <= 0 { // ### handle underflow gracefully
			this.remainingDuration = 0
			this.goingDown = false
		}
	} else {
		this.remainingDuration += frameTime
		if this.remainingDuration >= this.Duration {
			this.remainingDuration = this.Duration
			this.goingDown = true
		}
	}

	percentage := float64(float64(this.remainingDuration) / float64(this.Duration))
	if this.Easing != nil {
		percentage = this.Easing(percentage)
	}
	this.currentValue = (this.To * (1.0 - float32(percentage))) + (this.From * float32(percentage))
}

func (this *FloatAnimation) Get() float32 {
	return this.currentValue
}

func (this *FloatAnimation) Restart() {
	this.remainingDuration = this.Duration
	this.currentValue = this.From
}
