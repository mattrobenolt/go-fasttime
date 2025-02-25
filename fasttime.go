// Package fasttime is a primitive replacement for builtin time, that provides
// a much faster (~approx 2x) alternative to [time.Now], as well as the [Instant]
// is smaller relative to a [time.Time] struct, 8 bytes vs 20 bytes.
package fasttime
