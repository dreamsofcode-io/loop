package loop

import (
	"context"
	"sync"
)

// Parallel provides the ability to range over a slice concurrently.
// Each element of the slice will be called within it's own goroutine.
//
// This function should not be used in hopes to speed up any pure compute
// operation as there is an associated cost with spawning a new goroutine.
// Instead, it makes sense if there are any long running tasks inside of
// your loop.
//
// The benchmarks in parallel_test.go show a good example of when this
// method will speed up performance. (using time.Sleep)
func Parallel[E any](xs []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := range ParallelTimes(len(xs)) {
			x := xs[i]
			if !yield(i, x) {
				break
			}
		}
	}
}

// ParallelTimes allows you to peform a parallel iteration for a given
// integer type.
//
// This is very similar to the loop.Parallel method, except that instead
// of looping over a slice of elements, it instead will range for
// the number of given
func ParallelTimes[Int intType](num Int) func(func(Int) bool) {
	return func(yield func(Int) bool) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(int(num))

		for i := range uint64(num) {
			go func() {
				defer wg.Done()

				select {
				case <-ctx.Done():
					return
				default:
					if !yield(Int(i)) {
						cancel()
						return
					}
				}
			}()
		}

		wg.Wait()
	}
}
