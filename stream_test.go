package streams

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestOf(t *testing.T) {
	s := Of("a", "b", "c")
	rs, ok := s.(*stream[string])
	require.True(t, ok)
	require.NotNil(t, rs)
	require.Equal(t, 3, len(rs.elements))
	require.Equal(t, 3, s.Len())
}

func TestStream_AllMatch(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.AllMatch(p)
	require.False(t, m)
	s = Of("A", "B", "C")
	m = s.AllMatch(p)
	require.True(t, m)
	m = s.AllMatch(nil)
	require.False(t, m)
	s = Of[string]()
	m = s.AllMatch(p)
	require.False(t, m)
}

func TestStream_AnyMatch(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.AnyMatch(p)
	require.True(t, m)
	m = s.AnyMatch(nil)
	require.False(t, m)
	s = Of("a", "b", "c")
	m = s.AnyMatch(p)
	require.False(t, m)
}

func TestStream_Concat(t *testing.T) {
	s := Of("a", "b", "c")
	add := Of("d", "e", "f")
	s2 := s.Concat(add)
	require.Equal(t, 6, s2.Len())

	sl := []string{"g", "h", "i"}
	add2 := Streamable[string](sl)
	s3 := s2.Concat(add2)
	require.Equal(t, 9, s3.Len())

	ts := &testStream[string]{
		elements: []string{"j", "k", "l"},
	}
	s4 := s3.Concat(ts)
	require.Equal(t, 12, s4.Len())
}

func TestStream_Count(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	}).Or(NewPredicate(func(v string) bool {
		return v == "a"
	}))
	c := s.Count(p)
	require.Equal(t, 6, c)

	c = s.Count(nil)
	require.Equal(t, 10, c)
}

func TestStream_Difference(t *testing.T) {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	s := s1.Difference(s2, StringComparator)
	require.Equal(t, 1, s.Len())
	o := s.NthMatch(nil, 1)
	v, ok := o.GetOk()
	require.True(t, ok)
	require.Equal(t, "a", v)

	s = s2.Difference(s1, StringComparator)
	require.Equal(t, 1, s.Len())
	o = s.NthMatch(nil, 1)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "d", v)
}

func TestStream_Distinct(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	s2 := s.Distinct()
	require.Equal(t, 10, s2.Len())

	s = Of("d", "d", "d", "b", "b", "b", "c", "c", "c", "a")
	s2 = s.Distinct()
	require.Equal(t, 4, s2.Len())
}

func TestStream_Filter(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	}).Or(NewPredicate(func(v string) bool {
		return v == "a"
	}))
	s2 := s.Filter(p)
	require.Equal(t, 6, s2.Len())

	s2 = s.Filter(nil)
	require.Equal(t, 10, s2.Len())

	p = p.Negate()
	s2 = s.Filter(p)
	require.Equal(t, 4, s2.Len())
}

func TestStream_FirstMatch(t *testing.T) {
	s := Of("d", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	o := s.FirstMatch(p)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "F", v)

	o = s.FirstMatch(nil)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "d", v)

	s = Of("a", "b", "c")
	o = s.FirstMatch(p)
	require.False(t, o.IsPresent())
}

func TestStream_ForEach(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	sl := make([]string, 0)
	c := NewConsumer(func(v string) {
		sl = append(sl, v)
	})
	require.Equal(t, 0, len(sl))
	s.ForEach(c)
	require.Equal(t, 10, len(sl))

	s.ForEach(nil)
}

func TestStream_Has(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	h := s.Has("a", StringComparator)
	require.True(t, h)
	h = s.Has("A", StringInsensitiveComparator)
	require.True(t, h)
	h = s.Has("z", StringComparator)
	require.False(t, h)
}

func TestStream_Intersection(t *testing.T) {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	s := s1.Intersection(s2, StringComparator)
	require.Equal(t, 2, s.Len())
	o := s.NthMatch(nil, 1)
	v, ok := o.GetOk()
	require.True(t, ok)
	require.Equal(t, "b", v)
	o = s.NthMatch(nil, 2)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "c", v)
}

func TestStream_LastMatch(t *testing.T) {
	s := Of("d", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	o := s.LastMatch(p)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "B", v)

	o = s.LastMatch(nil)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "c", v)

	s = Of("a", "b", "c")
	o = s.LastMatch(p)
	require.False(t, o.IsPresent())
}

func TestStream_Limit(t *testing.T) {
	s := Of("a", "b", "c")
	s2 := s.Limit(5)
	require.Equal(t, 3, s2.Len())

	s2 = s.Limit(1)
	require.Equal(t, 1, s2.Len())
}

func TestStream_Max(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	o := s.Max(c)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "j", v)

	o = s.Max(nil)
	require.False(t, o.IsPresent())

	s = Of[string]()
	o = s.Max(c)
	require.False(t, o.IsPresent())
}

func TestStream_Min(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	o := s.Min(c)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "a", v)

	o = s.Min(nil)
	require.False(t, o.IsPresent())

	s = Of[string]()
	o = s.Min(c)
	require.False(t, o.IsPresent())
}

func TestStream_NoneMatch(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.NoneMatch(p)
	require.False(t, m)
	m = s.NoneMatch(nil)
	require.True(t, m)
	s = Of("a", "b", "c")
	m = s.NoneMatch(p)
	require.True(t, m)
}

func TestStream_NthMatch(t *testing.T) {
	s := Of("d", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	o := s.NthMatch(p, 2)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "H", v)

	o = s.NthMatch(p, -2)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "E", v)

	o = s.NthMatch(nil, 2)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "j", v)

	o = s.NthMatch(nil, -2)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "B", v)

	s = Of("a", "b", "c")
	o = s.NthMatch(p, 1)
	require.False(t, o.IsPresent())
}

func TestStream_Skip(t *testing.T) {
	s := Of("1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	s2 := s.Skip(5)
	require.Equal(t, 5, s2.Len())

	s2 = s.Skip(10)
	require.Equal(t, 0, s2.Len())
}

func TestStream_SkipAndLimit(t *testing.T) {
	s := Of("1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	s2 := s.Skip(5).Limit(10)
	require.Equal(t, 5, s2.Len())
	s2 = s.Skip(5).Limit(2)
	require.Equal(t, 2, s2.Len())
}

func TestStream_Sorted(t *testing.T) {
	s := Of("d", "j", "f", "g", "h", "i", "e", "a", "b", "c")
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	s2 := s.Sorted(c)
	require.Equal(t, 10, s2.Len())
	rs2, ok := s2.(*stream[string])
	require.True(t, ok)
	require.Equal(t, "a", rs2.elements[0])

	s2 = s.Sorted(nil)
	require.Equal(t, 10, s2.Len())
	rs2, ok = s2.(*stream[string])
	require.True(t, ok)
	require.Equal(t, "d", rs2.elements[0])
}

func TestStream_SymmetricDifference(t *testing.T) {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	s := s1.SymmetricDifference(s2, StringComparator)
	require.Equal(t, 2, s.Len())
	o := s.NthMatch(nil, 1)
	v, ok := o.GetOk()
	require.True(t, ok)
	require.Equal(t, "a", v)
	o = s.NthMatch(nil, 2)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "d", v)

	s = s2.SymmetricDifference(s1, StringComparator)
	require.Equal(t, 2, s.Len())
	o = s.NthMatch(nil, 1)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "d", v)
	o = s.NthMatch(nil, 2)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "a", v)
}

func TestStream_Union(t *testing.T) {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	s := s1.Union(s2, StringComparator)
	require.Equal(t, 4, s.Len())
	o := s.NthMatch(nil, 1)
	v, ok := o.GetOk()
	require.True(t, ok)
	require.Equal(t, "a", v)
	o = s.NthMatch(nil, 2)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "b", v)
	o = s.NthMatch(nil, 3)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "c", v)
	o = s.NthMatch(nil, 4)
	v, ok = o.GetOk()
	require.True(t, ok)
	require.Equal(t, "d", v)
}

func TestStream_Unique(t *testing.T) {
	s := Of("a", "a", "b", "c", "c")
	s = s.Unique(StringComparator)
	require.Equal(t, 3, s.Len())

	s2 := Of(instruct{1}, instruct{1}, instruct{2}, instruct{3}, instruct{3})
	require.Equal(t, 3, s2.Distinct().Len())
	require.Equal(t, 3, s2.Unique(nil).Len())

	s3 := Of(&instruct{1}, &instruct{1}, &instruct{2}, &instruct{3}, &instruct{3})
	require.Equal(t, 5, s3.Distinct().Len())
	s3 = s3.Unique(NewComparator(func(v1, v2 *instruct) int {
		if v1.value < v2.value {
			return -1
		} else if v1.value > v2.value {
			return 1
		}
		return 0
	}))
	require.Equal(t, 3, s3.Len())
}
