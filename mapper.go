package streams

// Mapper is an interface for mapping (converting) one element type to another
type Mapper[T any, R any] interface {
	// Map converts the values in the input Stream and produces a Stream of output types
	Map(in Stream[T]) Stream[R]
}

// NewMapper creates a new Mapper that will use the provided Converter
func NewMapper[T any, R any](c Converter[T, R]) Mapper[T, R] {
	return mapper[T, R]{
		c: c,
	}
}

type mapper[T any, R any] struct {
	c Converter[T, R]
}

// Map converts the values in the input stream and produces a stream of output types
func (m mapper[T, R]) Map(in Stream[T]) Stream[R] {
	r := make([]R, 0, in.Len())
	in.ForEach(NewConsumer[T](func(v T) {
		r = append(r, m.c.Convert(v))
	}))
	return Of(r...)
}
