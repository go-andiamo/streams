package streams

// Predicate is the interface used by filtering and matching operations
//
// i.e. is used by: Stream.AllMatch, Stream.AnyMatch, Stream.Count, Stream.Filter, Stream.FirstMatch,
// Stream.LastMatch, Stream.NoneMatch and Stream.NthMatch
type Predicate[T any] interface {
	// Test evaluates this predicate against the supplied value
	Test(v T) bool
	// And creates a composed predicate that represents a short-circuiting logical AND of this predicate and another
	And(other Predicate[T]) Predicate[T]
	// Or creates a composed predicate that represents a short-circuiting logical OR of this predicate and another
	Or(other Predicate[T]) Predicate[T]
	// Negate creates a composed predicate that represents a logical NOT of this predicate
	Negate() Predicate[T]
}

// PredicateFunc is the function signature used to create a new Predicate
type PredicateFunc[T any] func(v T) bool

// NewPredicate creates a new Predicate from the function provided
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

// Test evaluates this predicate against the supplied value
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

// And creates a composed predicate that represents a short-circuiting logical AND of this predicate and another
func (p predicate[T]) And(other Predicate[T]) Predicate[T] {
	return predicate[T]{
		inner: p,
		and:   other,
	}
}

// Or creates a composed predicate that represents a short-circuiting logical OR of this predicate and another
func (p predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return predicate[T]{
		inner: p,
		or:    other,
	}
}

// Negate creates a composed predicate that represents a logical NOT of this predicate
func (p predicate[T]) Negate() Predicate[T] {
	return predicate[T]{
		inner:   p,
		negated: true,
	}
}
