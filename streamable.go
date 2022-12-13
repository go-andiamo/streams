package streams

import (
	"github.com/go-andiamo/gopt"
	"sort"
)

// Streamable is a type alias that provides a Stream interface around a slice
type Streamable[T any] []T

// AllMatch returns whether all elements of this stream match the provided predicate
//
// if the provided predicate is nil or the stream is empty, always returns false
func (s Streamable[T]) AllMatch(p Predicate[T]) bool {
	if p == nil || len(s) == 0 {
		return false
	}
	for _, v := range s {
		if !p.Test(v) {
			return false
		}
	}
	return true
}

// AnyMatch returns whether any elements of this stream match the provided predicate
//
// if the provided predicate is nil or the stream is empty, always returns false
func (s Streamable[T]) AnyMatch(p Predicate[T]) bool {
	if p != nil {
		for _, v := range s {
			if p.Test(v) {
				return true
			}
		}
	}
	return false
}

// Append creates a new stream with all the elements of this stream followed by the specified elements
func (s Streamable[T]) Append(items ...T) Stream[T] {
	return &stream[T]{
		elements: append(s, items...),
	}
}

// Concat creates a new stream with all the elements of this stream followed by all the elements of the added stream
func (s Streamable[T]) Concat(add Stream[T]) Stream[T] {
	r := &stream[T]{
		elements: make([]T, 0, len(s)+add.Len()),
	}
	r.elements = append(r.elements, s...)
	if as, ok := add.(*stream[T]); ok {
		r.elements = append(r.elements, as.elements...)
	} else if sas, ok := add.(Streamable[T]); ok {
		r.elements = append(r.elements, sas...)
	} else if ssl, ok := add.(*streamableSlice[T]); ok {
		r.elements = append(r.elements, *ssl.elements...)
	} else {
		_ = add.ForEach(NewConsumer(func(v T) error {
			r.elements = append(r.elements, v)
			return nil
		}))
	}
	return r
}

// Count returns the count of elements that match the provided predicate
//
// If the predicate is nil, returns the count of all elements
func (s Streamable[T]) Count(p Predicate[T]) int {
	if p == nil {
		return len(s)
	}
	c := 0
	for _, v := range s {
		if p.Test(v) {
			c++
		}
	}
	return c
}

// Difference creates a new stream that is the set difference between this and the supplied other stream
//
// equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)
func (s Streamable[T]) Difference(other Stream[T], c Comparator[T]) Stream[T] {
	if c == nil {
		return &stream[T]{}
	}
	p := NewPredicate[T](func(v T) bool {
		return !other.Has(v, c)
	})
	return s.Filter(p)
}

// Distinct creates a new stream of distinct elements in this stream
func (s Streamable[T]) Distinct() Stream[T] {
	dvs := map[any]bool{}
	r := &stream[T]{}
	for _, v := range s {
		if !dvs[v] {
			dvs[v] = true
			r.elements = append(r.elements, v)
		}
	}
	return r
}

// Filter creates a new stream of elements in this stream that match the provided predicate
//
// if the provided predicate is nil, all elements in this stream are returned
func (s Streamable[T]) Filter(p Predicate[T]) Stream[T] {
	r := &stream[T]{}
	for _, v := range s {
		if p == nil || p.Test(v) {
			r.elements = append(r.elements, v)
		}
	}
	return r
}

// FirstMatch returns an optional of the first element that matches the provided predicate
//
// if no elements match the provided predicate, an empty (not present) optional is returned
//
// if the provided predicate is nil, the first element in this stream is returned
func (s Streamable[T]) FirstMatch(p Predicate[T]) *gopt.Optional[T] {
	for _, v := range s {
		if p == nil || p.Test(v) {
			return gopt.Of[T](v)
		}
	}
	return gopt.Empty[T]()
}

// ForEach performs an action on each element of this stream
//
// the action to be performed is defined by the provided consumer
//
// if the provided consumer is nil, nothing is performed
func (s Streamable[T]) ForEach(c Consumer[T]) error {
	if c != nil {
		for _, v := range s {
			if err := c.Accept(v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Has returns whether this stream contains an element that is equal to the element value provided
//
// equality is determined using the provided comparator
//
// if the provided comparator is nil, always returns false
func (s Streamable[T]) Has(v T, c Comparator[T]) bool {
	if c != nil {
		for _, v2 := range s {
			if c.Compare(v, v2) == 0 {
				return true
			}
		}
	}
	return false
}

// Intersection creates a new stream that is the set intersection of this and the supplied other stream
//
// equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)
func (s Streamable[T]) Intersection(other Stream[T], c Comparator[T]) Stream[T] {
	if c == nil {
		return &stream[T]{}
	}
	p := NewPredicate[T](func(v T) bool {
		return other.Has(v, c)
	})
	return s.Filter(p)
}

// Iterator returns an iterator (pull) function
//
// the iterator function can be used in for loops, for example
//  next := strm.Iterator()
//  for v, ok := next(); ok; v, ok = next() {
//      fmt.Println(v)
//  }
//
// Iterator can also optionally accept varargs of Predicate - which, if specified, are logically OR-ed on each pull to ensure
// that pulled elements match
func (s Streamable[T]) Iterator(ps ...Predicate[T]) func() (T, bool) {
	curr := 0
	if p := joinPredicates[T](ps...); p != nil {
		return func() (T, bool) {
			var r T
			ok := false
			for i := curr; i < len(s); i++ {
				if p.Test(s[i]) {
					ok = true
					r = s[i]
					curr = i + 1
					break
				}
			}
			return r, ok
		}
	} else {
		return func() (T, bool) {
			var r T
			ok := false
			if curr < len(s) {
				r, ok = s[curr], true
				curr++
			}
			return r, ok
		}
	}
}

// LastMatch returns an optional of the last element that matches the provided predicate
//
// if no elements match the provided predicate, an empty (not present) optional is returned
//
// if the provided predicate is nil, the last element in this stream is returned
func (s Streamable[T]) LastMatch(p Predicate[T]) *gopt.Optional[T] {
	for i := len(s) - 1; i >= 0; i-- {
		if p == nil || p.Test(s[i]) {
			return gopt.Of[T](s[i])
		}
	}
	return gopt.Empty[T]()
}

// Len returns the length (number of elements) of this stream
func (s Streamable[T]) Len() int {
	return len(s)
}

// Limit creates a new stream whose number of elements is limited to the value provided
//
// if the maximum size is greater than the length of this stream, all elements are returned
func (s Streamable[T]) Limit(maxSize int) Stream[T] {
	max := absZero(maxSize)
	if l := len(s); l < max {
		max = l
	}
	return &stream[T]{
		elements: s[0:max],
	}
}

// Max returns the maximum element of this stream according to the provided comparator
//
// if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned
func (s Streamable[T]) Max(c Comparator[T]) *gopt.Optional[T] {
	if l := len(s); l > 0 && c != nil {
		r := s[0]
		for i := 1; i < l; i++ {
			if c.Compare(s[i], r) > 0 {
				r = s[i]
			}
		}
		return gopt.Of(r)
	}
	return gopt.Empty[T]()
}

// Min returns the minimum element of this stream according to the provided comparator
//
// if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned
func (s Streamable[T]) Min(c Comparator[T]) *gopt.Optional[T] {
	if l := len(s); l > 0 && c != nil {
		r := s[0]
		for i := 1; i < l; i++ {
			if c.Compare(s[i], r) < 0 {
				r = s[i]
			}
		}
		return gopt.Of(r)
	}
	return gopt.Empty[T]()
}

// MinMax returns the minimum and maximum element of this stream according to the provided comparator
//
// if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned for both
func (s Streamable[T]) MinMax(c Comparator[T]) (*gopt.Optional[T], *gopt.Optional[T]) {
	if l := len(s); l > 0 && c != nil {
		mn := s[0]
		mx := mn
		for i := 1; i < l; i++ {
			if c.Compare(s[i], mn) < 0 {
				mn = s[i]
			} else if c.Compare(s[i], mx) > 0 {
				mx = s[i]
			}
		}
		return gopt.Of(mn), gopt.Of(mx)
	}
	return gopt.Empty[T](), gopt.Empty[T]()
}

// NoneMatch returns whether none of the elements of this stream match the provided predicate
//
// if the provided predicate is nil or the stream is empty, always returns true
func (s Streamable[T]) NoneMatch(p Predicate[T]) bool {
	if p != nil {
		for _, v := range s {
			if p.Test(v) {
				return false
			}
		}
	}
	return true
}

// NthMatch returns an optional of the nth matching element (1 based) according to the provided predicate
//
// if the nth argument is negative, the nth is taken as relative to the last
//
// if the provided predicate is nil, any element is taken as matching
//
// if no elements match in the specified position, an empty (not present) optional is returned
func (s Streamable[T]) NthMatch(p Predicate[T], nth int) *gopt.Optional[T] {
	absn := absInt(nth)
	if absn > len(s) {
		return gopt.Empty[T]()
	}
	c := 0
	if p == nil && nth < 0 {
		return gopt.Of[T](s[len(s)-absn])
	} else if p == nil && nth > 0 {
		return gopt.Of[T](s[nth-1])
	} else if nth < 0 {
		nth = absn
		for i := len(s) - 1; i >= 0; i-- {
			if p.Test(s[i]) {
				c++
				if c == nth {
					return gopt.Of[T](s[i])
				}
			}
		}
	} else if nth > 0 {
		for _, v := range s {
			if p.Test(v) {
				c++
				if c == nth {
					return gopt.Of[T](v)
				}
			}
		}
	}
	return gopt.Empty[T]()
}

// Reverse creates a new stream composed of elements from this stream but in reverse order
func (s Streamable[T]) Reverse() Stream[T] {
	l := len(s)
	r := &stream[T]{
		elements: make([]T, l),
	}
	for i := 0; i < l; i++ {
		r.elements[i] = s[l-i-1]
	}
	return r
}

// Skip creates a new stream consisting of this stream after discarding the first n elements
//
// if the specified n to skip is equal to or greater than the number of elements in this stream,
// an empty stream is returned
func (s Streamable[T]) Skip(n int) Stream[T] {
	skip := absZero(n)
	if l := len(s); skip >= l {
		skip = l
	}
	return &stream[T]{
		elements: s[skip:],
	}
}

// Slice creates a new stream composed of elements from this stream starting at the specified start and including
// the specified count (or to the end)
//
// the start is zero based (and less than zero is ignored)
//
// if the specified count is negative, items are selected from the start and then backwards by the count
func (s Streamable[T]) Slice(start int, count int) Stream[T] {
	start = absZero(start)
	end := start + count
	if count < 0 {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	return &stream[T]{
		elements: s[start:end],
	}
}

// Sorted creates a new stream consisting of the elements of this stream, sorted according to the provided comparator
//
// if the provided comparator is nil, the elements are not sorted
func (s Streamable[T]) Sorted(c Comparator[T]) Stream[T] {
	es := make([]T, 0, len(s))
	es = append(es, s...)
	r := &stream[T]{
		elements: es,
	}
	if c != nil {
		sort.Slice(r.elements, func(i, j int) bool {
			return c.Less(r.elements[i], r.elements[j])
		})
	}
	return r
}

// SymmetricDifference creates a new stream that is the set symmetric difference between this and the supplied other stream
//
// equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)
func (s Streamable[T]) SymmetricDifference(other Stream[T], c Comparator[T]) Stream[T] {
	if c == nil {
		return &stream[T]{}
	}
	i := s.Intersection(other, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Filter(p).Concat(other.Filter(p))
}

// Union creates a new stream that is the set union of this and the supplied other stream
//
// equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)
func (s Streamable[T]) Union(other Stream[T], c Comparator[T]) Stream[T] {
	if c == nil {
		return &stream[T]{}
	}
	i := s.Intersection(other, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Concat(other.Filter(p))
}

// Unique creates a new stream of unique elements in this stream
//
// uniqueness is determined using the provided comparator
//
// if provided comparator is nil but the value type of elements in this stream are directly mappable (i.e. primitive or non-pointer types) then
// Distinct is used as the result, otherwise returns an empty stream
func (s Streamable[T]) Unique(c Comparator[T]) Stream[T] {
	r := &stream[T]{}
	var vt T
	if c != nil {
		pres := make([]bool, len(s))
		for i, v := range s {
			if !pres[i] {
				for j := i + 1; j < len(s); j++ {
					if !pres[j] && c.Compare(v, s[j]) == 0 {
						pres[j] = true
					}
				}
				pres[i] = true
				r.elements = append(r.elements, v)
			}
		}
	} else if isDistinctable(vt) {
		return s.Distinct()
	}
	return r
}
