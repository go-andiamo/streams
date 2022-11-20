package streams

// Comparator is the interface used to compare elements of a Stream
//
// This interface is ued when sorting, when finding min/max of a stream
// and is also used to determine equality during set operations
// (Stream.Difference, Stream.Intersection, Stream.SymmetricDifference and Stream.Union)
type Comparator[T any] interface {
	// Compare compares the two values lexicographically, i.e.:
	//
	// * the result should be 0 if v1 == v2
	//
	// * the result should be -1 if v1 < v2
	//
	// * the result should be 1 if v1 > v2
	Compare(v1, v2 T) int
	// Less returns true if v1 < v2, otherwise false
	Less(v1, v2 T) bool
	// LessOrEqual returns true if v1 <= v2, otherwise false
	LessOrEqual(v1, v2 T) bool
	// Greater returns true if v1 > v2, otherwise false
	Greater(v1, v2 T) bool
	// GreaterOrEqual returns true if v1 >= v2, otherwise false
	GreaterOrEqual(v1, v2 T) bool
	// Equal returns true if v1 == v2, otherwise false
	Equal(v1, v2 T) bool
	// NotEqual returns true if v1 != v2, otherwise false
	NotEqual(v1, v2 T) bool
	// Reversed creates a new comparator that imposes the reverse ordering to this comparator
	//
	// the reversal is against less/greater as well as against equality/non-equality
	Reversed() Comparator[T]
	// Then creates a new comparator from this comparator, with a following then comparator
	// that is used when the initial comparison yields equal
	Then(other Comparator[T]) Comparator[T]
}

// ComparatorFunc is the function signature used to create a new Comparator
//
// the function should compare the two values provided lexicographically, i.e.:
//
// * the result should be 0 if v1 == v2
//
// * the result should be -1 if v1 < v2
//
// * the result should be 1 if v1 > v2
type ComparatorFunc[T any] func(v1, v2 T) int

// NewComparator creates a new Comparator from the function provided
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

// Compare compares the two values lexicographically, i.e.:
//
// * the result should be 0 if v1 == v2
//
// * the result should be -1 if v1 < v2
//
// * the result should be 1 if v1 > v2
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

// Less returns true if v1 < v2, otherwise false
func (c comparator[T]) Less(v1, v2 T) bool {
	return c.Compare(v1, v2) < 0
}

// LessOrEqual returns true if v1 <= v2, otherwise false
func (c comparator[T]) LessOrEqual(v1, v2 T) bool {
	return c.Compare(v1, v2) <= 0
}

// Greater returns true if v1 > v2, otherwise false
func (c comparator[T]) Greater(v1, v2 T) bool {
	return c.Compare(v1, v2) > 0
}

// GreaterOrEqual returns true if v1 >= v2, otherwise false
func (c comparator[T]) GreaterOrEqual(v1, v2 T) bool {
	return c.Compare(v1, v2) >= 0
}

// Equal returns true if v1 == v2, otherwise false
func (c comparator[T]) Equal(v1, v2 T) bool {
	if r := c.Compare(v1, v2); c.reversed {
		return r != 0
	} else {
		return r == 0
	}
}

// NotEqual returns true if v1 != v2, otherwise false
func (c comparator[T]) NotEqual(v1, v2 T) bool {
	if r := c.Compare(v1, v2); c.reversed {
		return r == 0
	} else {
		return r != 0
	}
}

// Reversed creates a new comparator that imposes the reverse ordering to this comparator
//
// the reversal is against less/greater as well as against equality/non-equality
func (c comparator[T]) Reversed() Comparator[T] {
	return comparator[T]{
		f:        c.f,
		reversed: !c.reversed,
	}
}

// Then creates a new comparator from this comparator, with a following then comparator
// that is used when the initial comparison yields equal
func (c comparator[T]) Then(other Comparator[T]) Comparator[T] {
	return comparator[T]{
		inner: c,
		then:  other,
	}
}
