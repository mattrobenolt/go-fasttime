package fasttime

import "time"

var monotonicRoot = time.Now()

// Instant is roughly equivalent to a [time.Time], but backed by a nanosecond
// duration since process start.
type Instant int64

// ToTime converts an [Instant] into a [time.Time].
func (i Instant) ToTime() time.Time {
	return monotonicRoot.Add(time.Duration(i))
}

// Sub subtracts two [Instant]s, similar to [time.Time.Sub].
func (i Instant) Sub(u Instant) time.Duration {
	return time.Duration(i - u)
}

func (i Instant) String() string {
	return i.ToTime().String()
}

// Now is roughly equivalent to [time.Now], but returns an [Instant].
func Now() Instant {
	return Instant(time.Since(monotonicRoot))
}

// Since is roughly equivalent to [time.Since], but operates on [Instant]s.
func Since(i Instant) time.Duration {
	return Now().Sub(i)
}
