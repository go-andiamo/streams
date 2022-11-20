package streams

type Predicate[T any] interface {
	Test(v T) bool
	And(other Predicate[T]) Predicate[T]
	Or(other Predicate[T]) Predicate[T]
	Negate() Predicate[T]
}

type PredicateFunc[T any] func(v T) bool

func NewPredicate[T any](f PredicateFunc[T]) Predicate[T] {
	return predicate[T]{
		f: f,
	}
}

type predicate[T any] struct {
	f       PredicateFunc[T]
	inner   Predicate[T]
	negated bool
	or      Predicate[T]
	and     Predicate[T]
}

func (p predicate[T]) Test(v T) bool {
	var r bool
	if p.f != nil {
		r = p.f(v)
	} else {
		r = p.inner.Test(v)
	}
	if r && p.and != nil {
		r = r && p.and.Test(v)
	}
	if !r && p.or != nil {
		r = p.or.Test(v)
	}
	if p.negated {
		return !r
	}
	return r
}

func (p predicate[T]) And(other Predicate[T]) Predicate[T] {
	return predicate[T]{
		inner: p,
		and:   other,
	}
}

func (p predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return predicate[T]{
		inner: p,
		or:    other,
	}
}

func (p predicate[T]) Negate() Predicate[T] {
	return predicate[T]{
		inner:   p,
		negated: true,
	}
}
