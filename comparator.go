package streams

type Comparator[T any] interface {
	Compare(v1, v2 T) int
	Less(v1, v2 T) bool
	LessOrEqual(v1, v2 T) bool
	Greater(v1, v2 T) bool
	GreaterOrEqual(v1, v2 T) bool
	Equal(v1, v2 T) bool
	NotEqual(v1, v2 T) bool
	Reversed() Comparator[T]
	Then(other Comparator[T]) Comparator[T]
}

type ComparatorFunc[T any] func(v1, v2 T) int

func NewComparator[T any](f ComparatorFunc[T]) Comparator[T] {
	return comparator[T]{
		f: f,
	}
}

type comparator[T any] struct {
	f        ComparatorFunc[T]
	inner    Comparator[T]
	then     Comparator[T]
	reversed bool
}

func (c comparator[T]) Compare(v1, v2 T) int {
	result := 0
	if c.f != nil {
		result = c.f(v1, v2)
	} else {
		result = c.inner.Compare(v1, v2)
	}
	if c.reversed {
		result = 0 - result
	}
	if result == 0 && c.then != nil {
		result = c.then.Compare(v1, v2)
	}
	return result
}

func (c comparator[T]) Less(v1, v2 T) bool {
	return c.Compare(v1, v2) < 0
}

func (c comparator[T]) LessOrEqual(v1, v2 T) bool {
	return c.Compare(v1, v2) <= 0
}

func (c comparator[T]) Greater(v1, v2 T) bool {
	return c.Compare(v1, v2) > 0
}

func (c comparator[T]) GreaterOrEqual(v1, v2 T) bool {
	return c.Compare(v1, v2) >= 0
}

func (c comparator[T]) Equal(v1, v2 T) bool {
	if r := c.Compare(v1, v2); c.reversed {
		return r != 0
	} else {
		return r == 0
	}
}

func (c comparator[T]) NotEqual(v1, v2 T) bool {
	if r := c.Compare(v1, v2); c.reversed {
		return r == 0
	} else {
		return r != 0
	}
}

func (c comparator[T]) Reversed() Comparator[T] {
	return comparator[T]{
		f:        c.f,
		reversed: !c.reversed,
	}
}

func (c comparator[T]) Then(other Comparator[T]) Comparator[T] {
	return comparator[T]{
		inner: c,
		then:  other,
	}
}
