package streams

// Reducer is the interface used to perform reductions (folds/accumulations)
type Reducer[T any, R any] interface {
	// Reduce performs a reduction of the supplied Stream
	Reduce(s Stream[T]) R
}

// NewReducer creates a new Reducer that will use the supplied Accumulator
//
// NewReducer panics if a nil Accumulator is supplied
func NewReducer[T any, R any](accumulator Accumulator[T, R]) Reducer[T, R] {
	if accumulator == nil {
		panic("accumulator cannot be nil")
	}
	return &reducer[T, R]{
		accumulator: accumulator,
	}
}

type reducer[T any, R any] struct {
	accumulator Accumulator[T, R]
}

// Reduce performs a reduction of the supplied Stream
func (r reducer[T, R]) Reduce(s Stream[T]) R {
	var result R
	_ = s.ForEach(NewConsumer[T](func(v T) error {
		result = r.accumulator.Apply(v, result)
		return nil
	}))
	return result
}
