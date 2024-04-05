package loop_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dreamsofcode-io/loop"
)

func TestPool(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := make([]int, 0, 3)

	for _, x := range loop.Pool(xs, 3) {
		res = append(res, x)
		time.Sleep(time.Millisecond)
		break
	}

	assert.Len(t, res, 3)
}
