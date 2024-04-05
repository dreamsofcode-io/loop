package loop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dreamsofcode-io/loop"
)

func TestRange(t *testing.T) {
	type input struct {
		start int
		stop  int
	}

	testCases := []struct {
		name  string
		input input
		wants []int
	}{
		{
			name: "happy path",
			input: input{
				start: 5,
				stop:  10,
			},
			wants: []int{5, 6, 7, 8, 9},
		},
		{
			name: "single item",
			input: input{
				start: 0,
				stop:  0,
			},
			wants: nil,
		},
		{
			name: "reverse item",
			input: input{
				start: 10,
				stop:  0,
			},
			wants: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res []int
			for i := range loop.Range(tc.input.start, tc.input.stop) {
				res = append(res, i)
			}

			assert.Equal(t, res, tc.wants)
		})
	}
}

func TestRangeWithStep(t *testing.T) {
	type input struct {
		start int
		stop  int
		step  int
	}

	testCases := []struct {
		name  string
		input input
		wants []int
	}{
		{
			name: "happy path",
			input: input{
				start: 0,
				stop:  10,
				step:  2,
			},
			wants: []int{0, 2, 4, 6, 8},
		},
		{
			name: "happy path #2",
			input: input{
				start: 10,
				stop:  21,
				step:  5,
			},
			wants: []int{10, 15, 20},
		},
		{
			name: "single item",
			input: input{
				start: 0,
				stop:  0,
				step:  10,
			},
			wants: nil,
		},
		{
			name: "reverse step",
			input: input{
				start: 10,
				stop:  0,
				step:  -2,
			},
			wants: []int{10, 8, 6, 4, 2},
		},
		{
			name: "reverse step #2",
			input: input{
				start: 12,
				stop:  0,
				step:  -3,
			},
			wants: []int{12, 9, 6, 3},
		},
		{
			name: "zero step",
			input: input{
				start: 0,
				stop:  100,
				step:  0,
			},
			wants: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res []int
			for i := range loop.RangeWithStep(tc.input.start, tc.input.stop, tc.input.step) {
				res = append(res, i)
			}

			assert.Equal(t, res, tc.wants)
		})
	}
}
