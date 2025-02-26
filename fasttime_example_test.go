package fasttime_test

import (
	"fmt"
	"time"

	"go.withmatt.com/fasttime"
)

func ExampleNow() {
	// Get a new Instant representing current time with absolute precision.
	start := fasttime.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Println(fasttime.Since(start))

	// Convert the Instant into a traditional time.Time object if needed.
	start.ToTime()
}

func ExampleClock() {
	// Create a new cached clock with 10ms precision.
	clock := fasttime.NewClock(10 * time.Millisecond)
	defer clock.Stop()

	// Get our new cached time as an Instant.
	start := clock.Now()

	// Sleep 5ms, which is less than our Clock ticker
	// so there should be no difference with Since.
	time.Sleep(5 * time.Millisecond)
	fmt.Println(clock.Since(start))

	// Sleep another 10ms, which should be in the next tick
	// for the Clock, but we should get a duration of about
	// 10ms and not 15ms due to our precision.
	time.Sleep(10 * time.Millisecond)
	fmt.Println(clock.Since(start))
}
