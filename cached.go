package fasttime

import (
	"context"
	"sync/atomic"
	"time"
)

// Clock is an instance of a clock whose time increments roughly at a
// configured granularity, but lookups are effectively free relative to
// normal [time.Now].
type Clock struct {
	ctx    context.Context
	cancel context.CancelFunc
	now    atomic.Int64
}

// NewClock creates a new Clock configured to tick at approximately
// granularity intervals. Clock is running when created, and may be stopped
// by calling Stop. A stopped Clock cannot be resumed.
func NewClock(granularity time.Duration) *Clock {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Clock{ctx: ctx, cancel: cancel}
	c.now.Store(int64(Now()))
	go c.run(granularity)
	return c
}

// Now returns an Instant that represents the current cached time.
// The Instant returned will never be in the future, but will always be
// less than or equal to the actual current time.
func (c *Clock) Now() Instant {
	return Instant(c.now.Load())
}

// Stop stops the Clock ticker and cannot be resumed.
func (c *Clock) Stop() {
	c.cancel()
}

func (c *Clock) run(granularity time.Duration) {
	t := time.NewTicker(granularity)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			c.now.Store(int64(Now()))
		case <-c.ctx.Done():
			return
		}
	}
}
