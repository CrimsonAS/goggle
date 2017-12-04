package animation

import (
	"math"
	"testing"
	"time"
)

var floatAnimationBench *FloatAnimation
var floatBench float32

func BenchmarkConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floatAnimationBench = &FloatAnimation{
			From:     0,
			To:       1000,
			Duration: 1000 * time.Millisecond,
		}
	}
}

func BenchmarkGet(b *testing.B) {
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

type expectedValue struct {
	advanceTime   time.Duration
	expectedValue float32
}

type testData struct {
	fromValue      float32
	toValue        float32
	duration       time.Duration
	expectedValues []expectedValue
}

// We don't test the case of calling Get() on an un-advanced animation. We
// probably should do that.
func TestGet(t *testing.T) {
	data := testData{
		fromValue: 0,
		toValue:   1000,
		duration:  1000 * time.Millisecond,
		expectedValues: []expectedValue{
			expectedValue{advanceTime: 0 * time.Millisecond, expectedValue: 0.0},
			expectedValue{advanceTime: 100 * time.Millisecond, expectedValue: 100.0},
			expectedValue{advanceTime: 100 * time.Millisecond, expectedValue: 200.0},
			expectedValue{advanceTime: 50 * time.Millisecond, expectedValue: 250.0},
			expectedValue{advanceTime: 50 * time.Millisecond, expectedValue: 300.0},
			expectedValue{advanceTime: 300 * time.Millisecond, expectedValue: 600.0},
			expectedValue{advanceTime: 400 * time.Millisecond, expectedValue: 1000.0},

			// back down
			expectedValue{advanceTime: 400 * time.Millisecond, expectedValue: 600.0},
			expectedValue{advanceTime: 300 * time.Millisecond, expectedValue: 300.0},
			expectedValue{advanceTime: 50 * time.Millisecond, expectedValue: 250.0},
			expectedValue{advanceTime: 50 * time.Millisecond, expectedValue: 200.0},
			expectedValue{advanceTime: 100 * time.Millisecond, expectedValue: 100.0},
			expectedValue{advanceTime: 0 * time.Millisecond, expectedValue: 100.0},
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

func testFloatAnimation(t *testing.T, data testData) {
	anim := FloatAnimation{
		From:     data.fromValue,
		To:       data.toValue,
		Duration: data.duration,
	}

	for idx, val := range data.expectedValues {
		anim.Advance(val.advanceTime)
		f := anim.Get()
		if !equalEnough(f, val.expectedValue) {
			t.Fatalf("Get for index %d advanced %s should have given %g, gave %g instead", idx, val.advanceTime, val.expectedValue, f)
		}
	}
}
