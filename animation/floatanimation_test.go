package animation

import (
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
