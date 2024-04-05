package loop

import (
	"context"
	"sync"
)

// Pool is used to perform bounded concurrency when iterating over the
// elements in a slice.
//
// The workers parameter specifies the size of the concurrency pool
// for iteration. For example, if a factor of 2 is given, then there
// will only even be 2 iterations running at once.
// 1 would effectively be a serial iteration.
//
// Bounded concurrency is useful in cases where the user may wish
// to perform concurrency but in a reduced rate, so as to avoid
// rate limits or running out of file descriptors.
func Pool[E any](xs []E, workers int) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		type iteration struct {
			val E
			i   int
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan iteration, workers)

		var wg sync.WaitGroup
		wg.Add(workers)

		for range workers {
			go func() {
				defer wg.Done()

				for x := range ch {
					select {
					case <-ctx.Done():
						return
					default:
						if !yield(x.i, x.val) {
							cancel()
						}
					}
				}
			}()
		}

		for i, x := range xs {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- iteration{
					val: x,
					i:   i,
				}
			}
		}

		close(ch)

		wg.Wait()
	}
}
