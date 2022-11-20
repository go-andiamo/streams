package streams

type AccumulatorFunc[T any, R any] func(t T, r R) R

type Accumulator[T any, R any] interface {
	Apply(t T, r R) R
}

func NewAccumulator[T any, R any](f AccumulatorFunc[T, R]) Accumulator[T, R] {
	return accumulator[T, R]{
		f: f,
	}
}

type accumulator[T any, R any] struct {
	f AccumulatorFunc[T, R]
}

func (a accumulator[T, R]) Apply(t T, r R) R {
	return a.f(t, r)
}
