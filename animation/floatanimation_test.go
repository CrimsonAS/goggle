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
	"math"
	"testing"
	"time"
)

var floatAnimationBench *FloatAnimation
var floatBench float32

func BenchmarkFloatAnimationConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floatAnimationBench = &FloatAnimation{
			From:     0,
			To:       1000,
			Duration: 1000 * time.Millisecond,
		}
	}
}

func BenchmarkFloatAnimationGet(b *testing.B) {
	anim := FloatAnimation{
		From:     0,
		To:       1000,
		Duration: 1000 * time.Millisecond,
	}
	for i := 0; i < b.N; i++ {
		anim.Advance(16 * time.Millisecond)
		floatBench = anim.Get()
	}
}

type floatAnimationTestStep struct {
	advanceTime   time.Duration
	expectedValue float32
}

type floatAnimationTestData struct {
	fromValue float32
	toValue   float32
	duration  time.Duration
	testSteps []floatAnimationTestStep
}

// We don't test the case of calling Get() on an un-advanced animation. We
// probably should do that.
func TestFloatAnimationGet(t *testing.T) {
	data := floatAnimationTestData{
		fromValue: 0,
		toValue:   1000,
		duration:  1000 * time.Millisecond,
		testSteps: []floatAnimationTestStep{
			floatAnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: 0.0},
			floatAnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: 100.0},
			floatAnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: 200.0},
			floatAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: 250.0},
			floatAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: 300.0},
			floatAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: 600.0},
			floatAnimationTestStep{advanceTime: 400 * time.Millisecond, expectedValue: 1000.0},

			// back down
			floatAnimationTestStep{advanceTime: 400 * time.Millisecond, expectedValue: 600.0},
			floatAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: 300.0},
			floatAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: 250.0},
			floatAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: 200.0},
			floatAnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: 100.0},
			floatAnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: 100.0},
		},
	}
	testFloatAnimation(t, data)
}

func equalEnough(one, two float32) bool {
	const TOLERANCE = 0.001
	if diff := math.Abs(float64(one - two)); diff < TOLERANCE {
		return true
	} else {
		return false
	}
}

func testFloatAnimation(t *testing.T, data floatAnimationTestData) {
	anim := FloatAnimation{
		From:     data.fromValue,
		To:       data.toValue,
		Duration: data.duration,
	}

	for idx, val := range data.testSteps {
		anim.Advance(val.advanceTime)
		f := anim.Get()
		if !equalEnough(f, val.expectedValue) {
			t.Fatalf("Get for index %d advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue, f)
		}
	}
}
