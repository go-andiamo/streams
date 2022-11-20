package streams

type Consumer[T any] interface {
	Accept(v T)
	AndThen(after Consumer[T]) Consumer[T]
}

type ConsumerFunc[T any] func(v T)

func NewConsumer[T any](f ConsumerFunc[T]) Consumer[T] {
	return consumer[T]{
		f: f,
	}
}

type consumer[T any] struct {
	f       ConsumerFunc[T]
	inner   Consumer[T]
	andThen Consumer[T]
}

func (c consumer[T]) Accept(v T) {
	if c.f != nil {
		c.f(v)
	} else {
		c.inner.Accept(v)
	}
	if c.andThen != nil {
		c.andThen.Accept(v)
	}
}

func (c consumer[T]) AndThen(after Consumer[T]) Consumer[T] {
	return consumer[T]{
		inner:   c,
		andThen: after,
	}
}
