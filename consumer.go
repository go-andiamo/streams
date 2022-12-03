package streams

// Consumer is the interface used by Stream.ForEach
type Consumer[T any] interface {
	// Accept is called by the user of the consumer to supply a value
	Accept(v T) error
	// AndThen creates a new consumer from the current with a subsequent action to be performed
	//
	// multiple consumers can be chained together as one using this method
	AndThen(after Consumer[T]) Consumer[T]
}

// NewConsumer creates a new consumer from the function provided
func NewConsumer[T any](f ConsumerFunc[T]) Consumer[T] {
	if f == nil {
		return nil
	}
	return consumer[T]{
		f: f,
	}
}

type consumer[T any] struct {
	f       ConsumerFunc[T]
	inner   Consumer[T]
	andThen Consumer[T]
}

// Accept is called by the user of the consumer to supply a value
func (c consumer[T]) Accept(v T) (err error) {
	if c.f != nil {
		err = c.f(v)
	} else {
		err = c.inner.Accept(v)
	}
	if err == nil && c.andThen != nil {
		err = c.andThen.Accept(v)
	}
	return
}

// AndThen creates a new consumer from the current with a subsequent action to be performed
//
// multiple consumers can be chained together as one using this method
func (c consumer[T]) AndThen(after Consumer[T]) Consumer[T] {
	return consumer[T]{
		inner:   c,
		andThen: after,
	}
}

// ConsumerFunc is the function signature used to create a new Consumer
type ConsumerFunc[T any] func(v T) error

func (f ConsumerFunc[T]) Accept(v T) error {
	return f(v)
}

func (f ConsumerFunc[T]) AndThen(after Consumer[T]) Consumer[T] {
	return NewConsumer[T](f).AndThen(after)
}
