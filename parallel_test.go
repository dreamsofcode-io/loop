package crayon_test

import (
	"testing"
	"time"

	"github.com/dreamsofcode-io/crayon"
)

var count uint64 = 1000

// Parallel
func TestParallel(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint64
		wants uint64
	}{
		{
			name:  "",
			input: []uint64{1, 2, 3, 4, 5},
			wants: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := uint64(0)
			for _, x := range crayon.Parallel(tc.input) {
				sum += x
			}

			if sum != tc.wants {
				t.Errorf("Reverse: %q, want %q", sum, tc.wants)
			}
		})
	}
}

func BenchmarkParallel(b *testing.B) {
	xs := make([]uint64, 0, count)
	for i := range count {
		xs = append(xs, i)
	}

	for range b.N {
		for _, _ = range crayon.Parallel(xs) {
			time.Sleep(time.Microsecond)
		}
	}
}

func BenchmarkParallelExisting(b *testing.B) {
	xs := make([]uint64, 0, count)
	for i := range count {
		xs = append(xs, i)
	}

	for range b.N {
		for range xs {
			time.Sleep(time.Microsecond)
		}
	}
}
