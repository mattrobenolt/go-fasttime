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
BenchmarkTime/time.Now-10               30490806                38.96 ns/op
BenchmarkTime/fasttime.Now
BenchmarkTime/fasttime.Now-10           61417069                19.75 ns/op
BenchmarkTime/fasttime.Clock.Now
BenchmarkTime/fasttime.Clock.Now-10             1000000000               0.6742 ns/op
BenchmarkTime/time.Since
BenchmarkTime/time.Since-10                     64346902                18.90 ns/op
BenchmarkTime/fasttime.Since
BenchmarkTime/fasttime.Since-10                 61757910                19.50 ns/op
BenchmarkTime/time.Now.Since
BenchmarkTime/time.Now.Since-10                 17892777                67.43 ns/op
BenchmarkTime/fasttime.Now.Since
BenchmarkTime/fasttime.Now.Since-10             32496326                36.88 ns/op
BenchmarkTime/fasttime.Instant.ToTime
BenchmarkTime/fasttime.Instant.ToTime-10        317467594                3.800 ns/op
PASS
ok      go.withmatt.com/fasttime        9.794s
```
