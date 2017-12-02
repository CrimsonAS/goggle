package animation

import (
	"time"
)

// Animate From a value To another value over a given Duration, and reverse
// direction after that Duration.
type FloatAnimation struct {
	From     float32       // from value
	To       float32       // to value
	Duration time.Duration // how long transition takes

	initialized       bool // initialized?
	goingDown         bool
	remainingDuration time.Duration // remaining Duration left before swap
	lastTick          time.Time     // when was Get last called
}

func (this *FloatAnimation) Get() float32 {
	tickTime := time.Since(this.lastTick)
	this.lastTick = time.Now()

	if !this.initialized {
		this.initialized = true
		this.Restart()
		return this.From
	}

	if this.goingDown {
		this.remainingDuration -= tickTime
		if this.remainingDuration <= 0 { // ### handle underflow gracefully
			this.remainingDuration = 0
			this.goingDown = false
		}
	} else {
		this.remainingDuration += tickTime
		if this.remainingDuration >= this.Duration {
			this.remainingDuration = this.Duration
			this.goingDown = true
		}
	}

	percentage := 1.0 - float32(this.remainingDuration)/float32(this.Duration)
	return (this.To - this.From) * percentage
}

func (this *FloatAnimation) Restart() {
	this.remainingDuration = this.Duration
	this.lastTick = time.Now()
}
