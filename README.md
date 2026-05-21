# Perfstat
## A simple and useful helper to provide performance statistics for Go programs

## Dependencies

Go

## Motivation

This was created because the profiles generated with `go tool pprof` measures
CPU time, and when I/O is the bottleneck, CPU time is useless.

Imagine that your application spends 80% of the time doing I/O, then you run a
CPU profile and sees that a hashing procedure is using almost 100% of the CPU
time, then, you remove this hashing procedure and the result is that now your 
program is only ~20% faster. This kind of thing happens a lot, because the time
that the goroutine spends on I/O is not measured by the CPU profiler.

This tool fixes this problem, because now we can measure the wall clock time
spent in sections of the code.

Example output:
```
rootFunc totalCost 1001702111 iterCost 1001702111 relativeCost 1.0000
        f3 totalCost 800447272 iterCost 800447272 relativeCost 0.7991
        f2 totalCost 150688036 iterCost 150688036 relativeCost 0.1504
        f1 totalCost 50513786 iterCost 50513786 relativeCost 0.0504
```

As we can see, rootFunc uses 100% of the wall clock time, then, inside this
function, f3 uses ~80% of the time. We can easily see that we can reduce up to
~80% of the elapsed time just by removing f3, in other words, we want to make f3 
faster.

All the values are stored as uint64, I usually set `cost` to the elapsed time in
nanoseconds, and `count` to number of iterations in case of loops, or 1
otherwise.

## Usage

Check the examples in the file `perfstat_example_test.go`

Note that this library will only work when compiled with the tag `perfstat`,
because of this, if you want, you can keep your code instrumented without
runtime cost, because without the tag, all procedures will be replaced with
NOP ones that the compiler can optimize out.
