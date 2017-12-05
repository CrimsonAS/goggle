package animation

import (
	"testing"
	"time"

	"github.com/CrimsonAS/goggle/sg"
)

var vec4AnimationBench *Vec4Animation
var vec4Bench sg.Vec4

func BenchmarkVec4AnimationConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vec4AnimationBench = &Vec4Animation{
			From:     sg.Vec4{1, 0, 1, 0},
			To:       sg.Vec4{0, 1, 0, 1},
			Duration: 1000 * time.Millisecond,
		}
	}
}

func BenchmarkVec4AnimationGet(b *testing.B) {
	anim := Vec4Animation{
		From:     sg.Vec4{1, 0, 1, 0},
		To:       sg.Vec4{0, 1, 0, 1},
		Duration: 1000 * time.Millisecond,
	}
	for i := 0; i < b.N; i++ {
		anim.Advance(16 * time.Millisecond)
		vec4Bench = anim.Get()
	}
}

type vec4AnimationTestStep struct {
	advanceTime   time.Duration
	expectedValue sg.Vec4
}

type vec4AnimationTestData struct {
	fromValue sg.Vec4
	toValue   sg.Vec4
	duration  time.Duration
	testSteps []vec4AnimationTestStep
}

// We don't test the case of calling Get() on an un-advanced animation. We
// probably should do that.
func TestVec4AnimationGet(t *testing.T) {
	data := vec4AnimationTestData{
		fromValue: sg.Vec4{1, 0, 1, 0},
		toValue:   sg.Vec4{0, 1, 0, 1},
		duration:  1000 * time.Millisecond,
		testSteps: []vec4AnimationTestStep{
			vec4AnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: sg.Vec4{1, 0, 1, 0}},
			vec4AnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: sg.Vec4{0.9, 0.1, 0.9, 0.1}},
			vec4AnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Vec4{0.6, 0.4, 0.6, 0.4}},
			vec4AnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Vec4{0.55, 0.45, 0.55, 0.45}},
			vec4AnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Vec4{0.50, 0.50, 0.50, 0.50}},
			vec4AnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Vec4{0.2, 0.8, 0.2, 0.8}},
			vec4AnimationTestStep{advanceTime: 200 * time.Millisecond, expectedValue: sg.Vec4{0.0, 1.0, 0.0, 1.0}},

			// back down
			vec4AnimationTestStep{advanceTime: 0 * time.Millisecond, expectedValue: sg.Vec4{0.0, 1.0, 0.0, 1.0}},
			vec4AnimationTestStep{advanceTime: 200 * time.Millisecond, expectedValue: sg.Vec4{0.2, 0.8, 0.2, 0.8}},
			vec4AnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Vec4{0.50, 0.50, 0.50, 0.50}},
			vec4AnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Vec4{0.55, 0.45, 0.55, 0.45}},
			vec4AnimationTestStep{advanceTime: 50 * time.Millisecond, expectedValue: sg.Vec4{0.6, 0.4, 0.6, 0.4}},
			vec4AnimationTestStep{advanceTime: 300 * time.Millisecond, expectedValue: sg.Vec4{0.9, 0.1, 0.9, 0.1}},
			vec4AnimationTestStep{advanceTime: 100 * time.Millisecond, expectedValue: sg.Vec4{1, 0, 1, 0}},
		},
	}
	testVec4Animation(t, data)
}

func testVec4Animation(t *testing.T, data vec4AnimationTestData) {
	anim := Vec4Animation{
		From:     data.fromValue,
		To:       data.toValue,
		Duration: data.duration,
	}

	for idx, val := range data.testSteps {
		anim.Advance(val.advanceTime)
		vec4 := anim.Get()
		if !equalEnough(vec4.X, val.expectedValue.X) {
			t.Fatalf("Get for index %d's X advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.X, vec4.X)
		}
		if !equalEnough(vec4.Y, val.expectedValue.Y) {
			t.Fatalf("Get for index %d's Y advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.Y, vec4.Y)
		}
		if !equalEnough(vec4.Z, val.expectedValue.Z) {
			t.Fatalf("Get for index %d's Z advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.Z, vec4.Z)
		}
		if !equalEnough(vec4.W, val.expectedValue.W) {
			t.Fatalf("Get for index %d's W advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue.W, vec4.W)
		}
	}
}
