package streams

import (
	"github.com/go-andiamo/gopt"
	"sort"
)

var myTestStream Stream[any] = &testStream[any]{}

type testStream[T any] struct {
	elements []T
}

func (s *testStream[T]) AllMatch(p Predicate[T]) bool {
	if p == nil || len(s.elements) == 0 {
		return false
	}
	for _, v := range s.elements {
		if !p.Test(v) {
			return false
		}
	}
	return true
}

func (s *testStream[T]) AnyMatch(p Predicate[T]) bool {
	if p != nil {
		for _, v := range s.elements {
			if p.Test(v) {
				return true
			}
		}
	}
	return false
}

func (s *testStream[T]) Concat(add Stream[T]) Stream[T] {
	r := &stream[T]{
		elements: make([]T, 0, len(s.elements)+add.Len()),
	}
	r.elements = append(r.elements, s.elements...)
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

func (s *testStream[T]) Count(p Predicate[T]) int {
	if p == nil {
		return len(s.elements)
	}
	c := 0
	for _, v := range s.elements {
		if p.Test(v) {
			c++
		}
	}
	return c
}

func (s *testStream[T]) Difference(with Stream[T], c Comparator[T]) Stream[T] {
	p := NewPredicate[T](func(v T) bool {
		return !with.Has(v, c)
	})
	return s.Filter(p)
}

func (s *testStream[T]) Distinct() Stream[T] {
	dvs := map[any]bool{}
	r := &stream[T]{}
	for _, v := range s.elements {
		if !dvs[v] {
			dvs[v] = true
			r.elements = append(r.elements, v)
		}
	}
	return r
}

func (s *testStream[T]) Filter(p Predicate[T]) Stream[T] {
	r := &stream[T]{}
	for _, v := range s.elements {
		if p == nil || p.Test(v) {
			r.elements = append(r.elements, v)
		}
	}
	return r
}

func (s *testStream[T]) FirstMatch(p Predicate[T]) gopt.Optional[T] {
	for _, v := range s.elements {
		if p == nil || p.Test(v) {
			return gopt.Of[T](v)
		}
	}
	return gopt.Empty[T]()
}

func (s *testStream[T]) ForEach(c Consumer[T]) {
	if c != nil {
		for _, v := range s.elements {
			c.Accept(v)
		}
	}
}

func (s *testStream[T]) Has(v T, c Comparator[T]) bool {
	if c != nil {
		for _, v2 := range s.elements {
			if c.Compare(v, v2) == 0 {
				return true
			}
		}
	}
	return false
}

func (s *testStream[T]) Intersection(with Stream[T], c Comparator[T]) Stream[T] {
	p := NewPredicate[T](func(v T) bool {
		return with.Has(v, c)
	})
	return s.Filter(p)
}

func (s *testStream[T]) LastMatch(p Predicate[T]) gopt.Optional[T] {
	for i := len(s.elements) - 1; i >= 0; i-- {
		if p == nil || p.Test(s.elements[i]) {
			return gopt.Of[T](s.elements[i])
		}
	}
	return gopt.Empty[T]()
}

func (s *testStream[T]) Len() int {
	return len(s.elements)
}

func (s *testStream[T]) Limit(maxSize int) Stream[T] {
	max := maxSize
	if l := len(s.elements); l < max {
		max = l
	}
	return &stream[T]{
		elements: s.elements[0:max],
	}
}

func (s *testStream[T]) Max(c Comparator[T]) gopt.Optional[T] {
	if l := len(s.elements); l > 0 && c != nil {
		r := s.elements[0]
		for i := 1; i < l; i++ {
			if c.Compare(s.elements[i], r) > 0 {
				r = s.elements[i]
			}
		}
		return gopt.Of(r)
	}
	return gopt.Empty[T]()
}

func (s *testStream[T]) Min(c Comparator[T]) gopt.Optional[T] {
	if l := len(s.elements); l > 0 && c != nil {
		r := s.elements[0]
		for i := 1; i < l; i++ {
			if c.Compare(s.elements[i], r) < 0 {
				r = s.elements[i]
			}
		}
		return gopt.Of(r)
	}
	return gopt.Empty[T]()
}

func (s *testStream[T]) NoneMatch(p Predicate[T]) bool {
	if p != nil {
		for _, v := range s.elements {
			if p.Test(v) {
				return false
			}
		}
	}
	return true
}

func (s *testStream[T]) NthMatch(p Predicate[T], nth int) gopt.Optional[T] {
	c := 0
	if nth < 0 {
		nth = 0 - nth
		for i := len(s.elements) - 1; i >= 0; i-- {
			if p == nil || p.Test(s.elements[i]) {
				c++
				if c == nth {
					return gopt.Of[T](s.elements[i])
				}
			}
		}
	} else if nth > 0 {
		for _, v := range s.elements {
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

func (s *testStream[T]) Skip(n int) Stream[T] {
	skip := n
	if l := len(s.elements); skip >= l {
		skip = l
	}
	return &stream[T]{
		elements: s.elements[skip:],
	}
}

func (s *testStream[T]) Sorted(c Comparator[T]) Stream[T] {
	es := make([]T, 0, len(s.elements))
	es = append(es, s.elements...)
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

func (s *testStream[T]) SymmetricDifference(with Stream[T], c Comparator[T]) Stream[T] {
	i := s.Intersection(with, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Filter(p).Concat(with.Filter(p))
}

func (s *testStream[T]) Union(with Stream[T], c Comparator[T]) Stream[T] {
	i := s.Intersection(with, c)
	p := NewPredicate[T](func(v T) bool {
		return !i.Has(v, c)
	})
	return s.Concat(with.Filter(p))
}

func (s *testStream[T]) Unique(c Comparator[T]) Stream[T] {
	r := &stream[T]{}
	var vt T
	if isDistinctable(vt) {
		return s.Distinct()
	} else if c != nil {
		prevs := make(map[int]bool, len(s.elements))
		for i, v := range s.elements {
			if !prevs[i] {
				for j := i + 1; j < len(s.elements); j++ {
					if !prevs[j] && c.Compare(v, s.elements[j]) == 0 {
						prevs[j] = true
					}
				}
				prevs[i] = true
				r.elements = append(r.elements, v)
			}
		}
	}
	return r
}
