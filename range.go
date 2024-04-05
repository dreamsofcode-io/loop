package loop

// Range creates a function iterator to iterate between two given
// integer like values.
//
// The first argument is the starting value, which is included in the
// iteration. The second argument is the stop value, which is when the
// iteration is stopped. This value is not included.
//
// This function is basically loop.RangeWithStep(start, stop, 1)
func Range[Int intType](start Int, stop Int) func(func(Int) bool) {
	return RangeWithStep(start, stop, 1)
}

// RangeWithStep creates a function iterator to iterate between two values
// with a given step incrementor.
//
// The first value is always returned (provided the stop value is value for
// the step amount)
// The stop value is not included.
// The step value can be either either greater than or less than 0. If the
// step is 0 then no iteration will take place.
func RangeWithStep[Int intType](start Int, stop Int, step Int) func(func(Int) bool) {
	return func(yield func(Int) bool) {
		if step > 0 {
			for i := uint64(start); i < uint64(stop); i += uint64(step) {
				if !yield(Int(i)) {
					return
				}
			}
		}

		if step < 0 {
			for i := uint64(start); i > uint64(stop); i += uint64(step) {
				if !yield(Int(i)) {
					return
				}
			}
		}

		if step == 0 {
			return
		}
	}
}
