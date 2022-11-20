package streams

type Reducer[T any, R any] interface {
	Reduce(s Stream[T]) R
}

func NewReducer[T any, R any](accumulator Accumulator[T, R]) Reducer[T, R] {
	return &reducer[T, R]{
		accumulator: accumulator,
	}
}

type reducer[T any, R any] struct {
	accumulator Accumulator[T, R]
}

func (r reducer[T, R]) Reduce(s Stream[T]) R {
	var result R
	s.ForEach(NewConsumer[T](func(v T) {
		result = r.accumulator.Apply(v, result)
	}))
	return result
}
