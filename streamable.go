package streams

import (
	"github.com/go-andiamo/gopt"
	"sort"
)

type Streamable[T any] []T

func StreamableOf[T any](sl []T) Stream[T] {
	return Of(sl...)
}

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

func (s Streamable[T]) Concat(add Stream[T]) Stream[T] {
	r := &stream[T]{
		elements: make([]T, 0, len(s)+add.Len()),
	}
	r.elements = append(r.elements, s...)
	if as, ok := add.(*stream[T]); ok {
		r.elements = append(r.elements, as.elements...)
	} else if sas, ok := add.(Streamable[T]); ok {
		r.elements = append(r.elements, sas...)
	} else {
		add.ForEach(NewConsumer(func(v T) {
			r.elements = append(r.elements, v)
		}))
	}
	return r
}

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

func (s Streamable[T]) Difference(with Stream[T], c Comparator[T]) Stream[T] {
	p := NewPredicate[T](func(v T) bool {
		return !with.Has(v, c)
	})
	return s.Filter(p)
}

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

func (s Streamable[T]) Filter(p Predicate[T]) Stream[T] {
	r := &stream[T]{}
	for _, v := range s {
		if p == nil || p.Test(v) {
			r.elements = append(r.elements, v)
		}
	}
	return r
}

func (s Streamable[T]) FirstMatch(p Predicate[T]) gopt.Optional[T] {
	for _, v := range s {
		if p == nil || p.Test(v) {
			return gopt.Of[T](v)
		}
	}
	return gopt.Empty[T]()
}

func (s Streamable[T]) ForEach(c Consumer[T]) {
	if c != nil {
		for _, v := range s {
			c.Accept(v)
		}
	}
}

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

func (s Streamable[T]) Intersection(with Stream[T], c Comparator[T]) Stream[T] {
	p := NewPredicate[T](func(v T) bool {
		return with.Has(v, c)
	})
	return s.Filter(p)
}

func (s Streamable[T]) LastMatch(p Predicate[T]) gopt.Optional[T] {
	for i := len(s) - 1; i >= 0; i-- {
		if p == nil || p.Test(s[i]) {
			return gopt.Of[T](s[i])
		}
	}
	return gopt.Empty[T]()
}

func (s Streamable[T]) Len() int {
	return len(s)
}

func (s Streamable[T]) Limit(maxSize int) Stream[T] {
	max := maxSize
	if l := len(s); l < max {
		max = l
	}
	return &stream[T]{
		elements: s[0:max],
	}
}

func (s Streamable[T]) Max(c Comparator[T]) gopt.Optional[T] {
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

func (s Streamable[T]) Min(c Comparator[T]) gopt.Optional[T] {
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

func (s Streamable[T]) NthMatch(p Predicate[T], nth int) gopt.Optional[T] {
	c := 0
	if nth < 0 {
		nth = 0 - nth
		for i := len(s) - 1; i >= 0; i-- {
			if p == nil || p.Test(s[i]) {
				c++
				if c == nth {
					return gopt.Of[T](s[i])
				}
			}
		}
	} else if nth > 0 {
		for _, v := range s {
			if p == nil || p.Test(v) {
				c++
				if c == nth {
					return gopt.Of[T](v)
				}
			}
		}
	}
	return gopt.Empty[T]()
}

func (s Streamable[T]) Skip(n int) Stream[T] {
	skip := n
	if l := len(s); skip >= l {
		skip = l
	}
	return &stream[T]{
		elements: s[skip:],
	}
}

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

func (s Streamable[T]) SymmetricDifference(with Stream[T], c Comparator[T]) Stream[T] {
	i := s.Intersection(with, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Filter(p).Concat(with.Filter(p))
}

func (s Streamable[T]) Union(with Stream[T], c Comparator[T]) Stream[T] {
	i := s.Intersection(with, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Concat(with.Filter(p))
}

func (s Streamable[T]) Unique(c Comparator[T]) Stream[T] {
	r := &stream[T]{}
	var vt T
	if isDistinctable(vt) {
		return s.Distinct()
	} else if c != nil {
		for i, v := range s {
			add := true
			for j := i + 1; j < len(s); j++ {
				if c.Compare(v, s[j]) == 0 {
					add = false
					break
				}
			}
			if add {
				r.elements = append(r.elements, v)
			}
		}
	}
	return r
}
