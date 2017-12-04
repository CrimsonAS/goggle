package animation

import (
	"testing"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

var colorAnimationBench *ColorAnimation
var colorBench sg.Color

func BenchmarkColorAnimationConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colorAnimationBench = &ColorAnimation{
			From:     sg.Color{1, 0, 1, 0},
			To:       sg.Color{0, 1, 0, 1},
			Duration: 1000 * time.Millisecond,
		}
	}
}

func BenchmarkColorAnimationGet(b *testing.B) {
	anim := ColorAnimation{
		From:     sg.Color{1, 0, 1, 0},
		To:       sg.Color{0, 1, 0, 1},
		Duration: 1000 * time.Millisecond,
	}
	for i := 0; i < b.N; i++ {
		anim.Advance(16 * time.Millisecond)
		colorBench = anim.Get()
	}
}

type colorAnimationTestStep struct {
	advanceTime   time.Duration
	expectedValue sg.Color
}

type colorAnimationTestData struct {
	fromValue sg.Color
	toValue   sg.Color
	duration  time.Duration
	testSteps []colorAnimationTestStep
}

// We don't test the case of calling Get() on an un-advanced animation. We
// probably should do that.
func TestColorAnimationGet(t *testing.T) {
	data := colorAnimationTestData{
		fromValue: sg.Color{1, 0, 1, 0},
		toValue:   sg.Color{0, 1, 0, 1},
		duration:  1000 * time.Millisecond,
		testSteps: []colorAnimationTestStep{
			colorAnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: sg.Color{1, 0, 1, 0}},
			colorAnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: sg.Color{0.9, 0.1, 0.9, 0.1}},
			colorAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Color{0.6, 0.4, 0.6, 0.4}},
			colorAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Color{0.55, 0.45, 0.55, 0.45}},
			colorAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Color{0.50, 0.50, 0.50, 0.50}},
			colorAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Color{0.2, 0.8, 0.2, 0.8}},
			colorAnimationTestStep{advanceTime: 200 * time.Millisecond, expectedValue: sg.Color{0.0, 1.0, 0.0, 1.0}},

			// back down
			colorAnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: sg.Color{0.0, 1.0, 0.0, 1.0}},
			colorAnimationTestStep{advanceTime: 200 * time.Millisecond, expectedValue: sg.Color{0.2, 0.8, 0.2, 0.8}},
			colorAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Color{0.50, 0.50, 0.50, 0.50}},
			colorAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Color{0.55, 0.45, 0.55, 0.45}},
			colorAnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Color{0.6, 0.4, 0.6, 0.4}},
			colorAnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Color{0.9, 0.1, 0.9, 0.1}},
			colorAnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: sg.Color{1, 0, 1, 0}},
		},
	}
	testColorAnimation(t, data)
}

func testColorAnimation(t *testing.T, data colorAnimationTestData) {
	anim := ColorAnimation{
		From:     data.fromValue,
		To:       data.toValue,
		Duration: data.duration,
	}

	for idx, val := range data.testSteps {
		anim.Advance(val.advanceTime)
		color := anim.Get()
		if !equalEnough(color.A(), val.expectedValue.A()) {
			t.Fatalf("Get for index %d's A advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.A(), color.A())
		}
		if !equalEnough(color.R(), val.expectedValue.R()) {
			t.Fatalf("Get for index %d's R advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.R(), color.R())
		}
		if !equalEnough(color.G(), val.expectedValue.G()) {
			t.Fatalf("Get for index %d's G advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.G(), color.G())
		}
		if !equalEnough(color.B(), val.expectedValue.B()) {
			t.Fatalf("Get for index %d's B advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.B(), color.B())
		}
	}
}
