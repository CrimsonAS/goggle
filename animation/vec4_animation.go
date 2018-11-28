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
