package fasttime_test

import (
	"testing"
	"time"

	"go.withmatt.com/fasttime"
)

var (
	timeSink     time.Time
	durationSink time.Duration
	instantSink  fasttime.Instant
)

func BenchmarkTime(b *testing.B) {
	b.Run("time.Now", func(b *testing.B) {
		for range b.N {
			timeSink = time.Now()
		}
	})

	b.Run("fasttime.Now", func(b *testing.B) {
		for range b.N {
			instantSink = fasttime.Now()
		}
	})

	b.Run("fasttime.Clock.Now", func(b *testing.B) {
		clock := fasttime.NewClock(100 * time.Millisecond)
		defer clock.Stop()
		for range b.N {
			instantSink = clock.Now()
		}
	})

	b.Run("time.Since", func(b *testing.B) {
		a := time.Now()
		for range b.N {
			durationSink = time.Since(a)
		}
	})

	b.Run("fasttime.Since", func(b *testing.B) {
		a := fasttime.Now()
		for range b.N {
			durationSink = fasttime.Since(a)
		}
	})

	b.Run("fasttime.Clock.Since", func(b *testing.B) {
		clock := fasttime.NewClock(100 * time.Millisecond)
		defer clock.Stop()
		a := clock.Now()
		for range b.N {
			durationSink = clock.Since(a)
		}
	})

	b.Run("time.Now.Since", func(b *testing.B) {
		for range b.N {
			durationSink = time.Since(time.Now())
		}
	})

	b.Run("fasttime.Now.Since", func(b *testing.B) {
		for range b.N {
			durationSink = fasttime.Since(fasttime.Now())
		}
	})

	b.Run("fasttime.Instant.ToTime", func(b *testing.B) {
		a := fasttime.Now()
		for range b.N {
			timeSink = a.ToTime()
		}
	})
}
