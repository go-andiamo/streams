package streams

type Mapper[T any, R any] interface {
	Map(in Stream[T]) Stream[R]
}

func NewMapper[T any, R any](c Converter[T, R]) Mapper[T, R] {
	return mapper[T, R]{
		c: c,
	}
}

type mapper[T any, R any] struct {
	c Converter[T, R]
}

func (m mapper[T, R]) Map(in Stream[T]) Stream[R] {
	r := make([]R, 0, in.Len())
	in.ForEach(NewConsumer[T](func(v T) {
		r = append(r, m.c.Convert(v))
	}))
	return Of(r...)
}
