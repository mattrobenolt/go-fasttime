# fasttime

```go
import "go.withmatt.com/fasttime"

// Get a new Instant representing current time with absolute precision.
start := fasttime.Now()
fasttime.Since(start)

// Convert the Instant into a traditional time.Time object if needed.
start.ToTime()

// Create a new cached clock with 100ms precision.
clock := fasttime.NewClock(100*time.Millisecond)
defer clock.Stop()

// Get our new cached time as an Instant.
clock.Now()
```

`fasttime.Now()` is much faster than builtin `time.Now()`, and works very similarly
without any golink tricks.

A `fasttime.Clock` is useful when you want super fast time, but are willing to
sacrifice precision. Typically a `time.Now()` is up to nanosecond precision if
your system allows it, and for a lot of use cases, that's more precision than
you care about. `fasttime.Clock` allows extremely quick (near free) access to a
a cached time within a configured granularity. This enables accessing
`fasttime.Clock.Now()` in extremely performance sensitive applications when
absolute precision isn't necessary.

## benchmarks

```
goos: darwin
goarch: arm64
pkg: go.withmatt.com/fasttime
cpu: Apple M1 Max
BenchmarkTime
BenchmarkTime/time.Now
BenchmarkTime/time.Now-10               29940246                39.93 ns/op
BenchmarkTime/fasttime.Now
BenchmarkTime/fasttime.Now-10           59820660                19.83 ns/op
BenchmarkTime/fasttime.Clock.Now
BenchmarkTime/fasttime.Clock.Now-10             1000000000               0.6755 ns/op
BenchmarkTime/time.Since
BenchmarkTime/time.Since-10                     63852006                18.77 ns/op
BenchmarkTime/fasttime.Since
BenchmarkTime/fasttime.Since-10                 60776544                19.42 ns/op
BenchmarkTime/fasttime.Clock.Since
BenchmarkTime/fasttime.Clock.Since-10           1000000000               0.6520 ns/op
BenchmarkTime/time.Now.Since
BenchmarkTime/time.Now.Since-10                 18779966                64.79 ns/op
BenchmarkTime/fasttime.Now.Since
BenchmarkTime/fasttime.Now.Since-10             32429546                36.85 ns/op
BenchmarkTime/fasttime.Instant.ToTime
BenchmarkTime/fasttime.Instant.ToTime-10        332593658                3.780 ns/op
PASS
ok      go.withmatt.com/fasttime        11.562s
```
