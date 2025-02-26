//go:build go1.24 && goexperiment.synctest

package fasttime_test

import (
	"testing"
	"testing/synctest"
)

func TestMonotonic(t *testing.T) {
	synctest.Run(func() {
		testMonotonic(t, false)
	})
}

func TestCached(t *testing.T) {
	synctest.Run(func() {
		testCached(t, false)
	})
}
