package streams

type ConverterFunc[T any, R any] func(v T) R

type Converter[T any, R any] interface {
	Convert(v T) R
}

func NewConverter[T any, R any](f ConverterFunc[T, R]) Converter[T, R] {
	return converter[T, R]{
		f: f,
	}
}

type converter[T any, R any] struct {
	f ConverterFunc[T, R]
}

func (c converter[T, R]) Convert(v T) R {
	return c.f(v)
}
