package streams

// Converter is the interface used by Mapper to convert one value type to another
type Converter[T any, R any] interface {
	// Convert converts a value of type T and returns a value of type R
	Convert(v T) (R, error)
}

// ConverterFunc is the function signature used to create a new Converter
type ConverterFunc[T any, R any] func(v T) (R, error)

// NewConverter creates a new Converter from the function provided
func NewConverter[T any, R any](f ConverterFunc[T, R]) Converter[T, R] {
	return converter[T, R]{
		f: f,
	}
}

type converter[T any, R any] struct {
	f ConverterFunc[T, R]
}

// Convert converts a value of type T and returns a value of type R
func (c converter[T, R]) Convert(v T) (R, error) {
	return c.f(v)
}
