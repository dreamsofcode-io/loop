package crayon

// Batch is used to turn any slice into an iterator of batches, with the size
// of each batch being the second parameter.
//
// This is useful if you want to perform batch operations on a large slice
// of elements, for example, breaking up a large request into multiple
// smaller ones.
func Batch[E any](xs []E, size uint) func(func(int, []E) bool) {
	return func(yield func(int, []E) bool) {
		if size == 0 {
			return
		}

		index := 0

		for i := uint(0); i < uint(len(xs)); i += size {
			top := min(uint(len(xs)), uint(i)+size)
			batch := xs[i:top]

			if !yield(index, batch) {
				return
			}

			index += 1
		}
	}
}
