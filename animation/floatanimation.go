/*
 * Copyright 2017 Crimson AS <info@crimson.no>
 * Author: Robin Burchell <robin.burchell@crimson.no>
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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
