package loop

import (
	"context"
	"sync"
)

func Pool[E any](xs []E, size int) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		type iteration struct {
			val E
			i   int
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan iteration, size)

		var wg sync.WaitGroup
		wg.Add(size)

		for range size {
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
