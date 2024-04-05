package loop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dreamsofcode-io/loop"
)

func TestBatchIncrementor(t *testing.T) {
	expected := 0

	for i, _ := range loop.Batch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2) {
		assert.Equal(t, expected, i)
		expected += 1
	}
}

func TestBatch(t *testing.T) {
	type input struct {
		items []int
		size  uint
	}

	testCases := []struct {
		name  string
		input input
		wants [][]int
	}{
		{
			name: "happy path",
			input: input{
				items: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				},
				size: 3,
			},
			wants: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10},
			},
		},
		{
			name: "batch size of 1",
			input: input{
				items: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				},
				size: 1,
			},
			wants: [][]int{
				{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10},
			},
		},
		{
			name: "batch size of 0",
			input: input{
				items: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				},
				size: 0,
			},
			wants: [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := [][]int{}
			for _, x := range loop.Batch(tc.input.items, tc.input.size) {
				res = append(res, x)
			}

			assert.Equal(t, res, tc.wants)
		})
	}
}
