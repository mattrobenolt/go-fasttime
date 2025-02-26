//go:build !(go1.24 && goexperiment.synctest)

package fasttime_test

import "testing"

func TestMonotonic(t *testing.T) {
	testMonotonic(t, true)
}

func TestCached(t *testing.T) {
	testCached(t, true)
}
