package loop_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dreamsofcode-io/loop"
)

var count uint64 = 1000

func TestParallelShouldNotPanic(t *testing.T) {
	xs := []int{}

	for range 10000 {
		xs = append(xs)
	}

	for i, _ := range loop.Parallel(xs) {
		if i > 300 {
			break
		}
	}
}

// Parallel
func TestParallel(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		wants int
	}{
		{
			name:  "",
			input: []int{1, 2, 3, 4, 5},
			wants: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := 0
			for _, x := range loop.Parallel(tc.input) {
				sum += x
			}

			assert.Equal(t, sum, tc.wants)
		})
	}
}

func BenchmarkParallel(b *testing.B) {
	xs := make([]uint64, 0, count)
	for i := range count {
		xs = append(xs, i)
	}

	for range b.N {
		for _, _ = range loop.Parallel(xs) {
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
