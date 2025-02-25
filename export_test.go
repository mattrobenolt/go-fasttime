package fasttime

import "time"

// XXX: export for tests
func ReplaceMonotonicRoot(t time.Time) time.Time {
	before := monotonicRoot
	monotonicRoot = t
	return before
}
