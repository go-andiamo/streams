package streams

// AccumulatorFunc is the function signature used to create a new Accumulator
type AccumulatorFunc[T any, R any] func(t T, r R) R

// Accumulator is the interface used Reducer to reduce to a single resultant
type Accumulator[T any, R any] interface {
	// Apply adds the value of T to R, and returns the new R
	Apply(t T, r R) R
}

// NewAccumulator creates a new Accumulator from the function provided
func NewAccumulator[T any, R any](f AccumulatorFunc[T, R]) Accumulator[T, R] {
	return accumulator[T, R]{
		f: f,
	}
}

type accumulator[T any, R any] struct {
	f AccumulatorFunc[T, R]
}

// Apply adds the value of T to R, and returns the new R
func (a accumulator[T, R]) Apply(t T, r R) R {
	return a.f(t, r)
}
