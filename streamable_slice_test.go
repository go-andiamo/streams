package streams

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestStreamableSlice(t *testing.T) {
	sl := &[]string{"a", "b", "c"}
	s := NewStreamableSlice(sl)
	require.Equal(t, 3, len(*sl))
	require.Equal(t, 3, s.Len())

	*sl = append(*sl, "d")
	require.Equal(t, 4, len(*sl))
	require.Equal(t, 4, s.Len())
}

func TestStreamableSlice_AllMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"D", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.AllMatch(p)
	require.False(t, m)
	s = NewStreamableSlice(&[]string{"A", "B", "C"})
	m = s.AllMatch(p)
	require.True(t, m)
	m = s.AllMatch(nil)
	require.False(t, m)
	s = NewStreamableSlice(&[]string{})
	m = s.AllMatch(p)
	require.False(t, m)
}

func TestStreamableSlice_AnyMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"D", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.AnyMatch(p)
	require.True(t, m)
	m = s.AnyMatch(nil)
	require.False(t, m)
	s = NewStreamableSlice(&[]string{"a", "b", "c"})
	m = s.AnyMatch(p)
	require.False(t, m)
}

func TestStreamableSlice_Append(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "b", "c"})
	s2 := s.Append("d", "e", "f")
	require.Equal(t, 6, s2.Len())
}

func TestStreamableSlice_AsSlice(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "b", "c"})
	sl := s.AsSlice()
	require.Equal(t, 3, len(sl))
}

func TestStreamableSlice_Concat(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "b", "c"})

	add := Of("d", "e", "f")
	s2 := s.Concat(add)
	require.Equal(t, 6, s2.Len())

	sl2 := []string{"g", "h", "i"}
	add2 := Streamable[string](sl2)
	s2 = s.Concat(add2)
	require.Equal(t, 6, s2.Len())

	add3 := NewStreamableSlice(&[]string{"j", "k", "l"})
	s3 := s.Concat(add3)
	require.Equal(t, 6, s3.Len())

	ts := &testStream[string]{
		elements: []string{"j", "k", "l"},
	}
	s2 = s.Concat(ts)
	require.Equal(t, 6, s2.Len())
}

func TestStreamableSlice_Count(t *testing.T) {
	s := NewStreamableSlice(&[]string{"D", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
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

func TestStreamableSlice_Difference(t *testing.T) {
	s1 := NewStreamableSlice(&[]string{"a", "b", "c"})
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

	s3 := s1.Difference(s2, nil)
	require.Equal(t, 0, s3.Len())
}

func TestStreamableSlice_Distinct(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	s2 := s.Distinct()
	require.Equal(t, 10, s2.Len())

	s = NewStreamableSlice(&[]string{"d", "d", "d", "b", "b", "b", "c", "c", "c", "a"})
	s2 = s.Distinct()
	require.Equal(t, 4, s2.Len())
}

func TestStreamableSlice_Filter(t *testing.T) {
	s := NewStreamableSlice(&[]string{"D", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
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

func TestStreamableSlice_FirstMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	o := s.FirstMatch(p)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "F", v)

	s = NewStreamableSlice(&[]string{"a", "b", "c"})
	o = s.FirstMatch(p)
	require.False(t, o.IsPresent())
}

func TestStreamableSlice_ForEach(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	s2 := make([]string, 0)
	c := NewConsumer(func(v string) error {
		s2 = append(s2, v)
		return nil
	})
	require.Equal(t, 0, len(s2))
	err := s.ForEach(c)
	require.NoError(t, err)
	require.Equal(t, 10, len(s2))

	err = s.ForEach(nil)
	require.NoError(t, err)

	c = NewConsumer(func(v string) error {
		return errors.New("whoops")
	})
	err = s.ForEach(c)
	require.Error(t, err)
	require.Equal(t, "whoops", err.Error())
}

func TestStreamableSlice_Has(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	h := s.Has("a", StringComparator)
	require.True(t, h)
	h = s.Has("A", StringInsensitiveComparator)
	require.True(t, h)
	h = s.Has("z", StringComparator)
	require.False(t, h)
}

func TestStreamableSlice_Intersection(t *testing.T) {
	s1 := NewStreamableSlice(&[]string{"a", "b", "c"})
	s2 := NewStreamableSlice(&[]string{"b", "c", "d"})
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

	s3 := s1.Intersection(s2, nil)
	require.Equal(t, 0, s3.Len())
}

func TestStreamableSlice_Iterator(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "a", "b", "c", "c"})
	count := 0
	str := ""
	iter := s.Iterator()
	for v, ok := iter(); ok; v, ok = iter() {
		count++
		str = str + v
	}
	require.Equal(t, 5, count)
	require.Equal(t, "aabcc", str)

	count = 0
	str = ""
	iter = s.Iterator(NewPredicate(func(v string) bool {
		return v == "a"
	}), NewPredicate(func(v string) bool {
		return v == "b"
	}))
	for v, ok := iter(); ok; v, ok = iter() {
		count++
		str = str + v
	}
	require.Equal(t, 3, count)
	require.Equal(t, "aab", str)
}

func TestStreamableSlice_LastMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
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

	s = NewStreamableSlice(&[]string{"a", "b", "c"})
	o = s.LastMatch(p)
	require.False(t, o.IsPresent())
}

func TestStreamableSlice_Limit(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "b", "c"})
	s2 := s.Limit(5)
	require.Equal(t, 3, s2.Len())

	s2 = s.Limit(1)
	require.Equal(t, 1, s2.Len())
}

func TestStreamableSlice_Max(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	o := s.Max(c)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "j", v)

	o = s.Max(nil)
	require.False(t, o.IsPresent())

	s = NewStreamableSlice(&[]string{})
	o = s.Max(c)
	require.False(t, o.IsPresent())
}

func TestStreamableSlice_Min(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	o := s.Min(c)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "a", v)

	o = s.Min(nil)
	require.False(t, o.IsPresent())

	s = NewStreamableSlice(&[]string{})
	o = s.Min(c)
	require.False(t, o.IsPresent())
}

func TestStreamableSlice_MinMax(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
	c := NewComparator(func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	mn, mx := s.MinMax(c)
	require.True(t, mn.IsPresent())
	v, _ := mn.GetOk()
	require.Equal(t, "a", v)
	require.True(t, mx.IsPresent())
	v, _ = mx.GetOk()
	require.Equal(t, "j", v)

	mn, mx = s.MinMax(nil)
	require.False(t, mn.IsPresent())
	require.False(t, mx.IsPresent())

	s = NewStreamableSlice(&[]string{})
	mn, mx = s.MinMax(c)
	require.False(t, mn.IsPresent())
	require.False(t, mx.IsPresent())
}

func TestStreamableSlice_NoneMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"D", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
	p := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	m := s.NoneMatch(p)
	require.False(t, m)
	m = s.NoneMatch(nil)
	require.True(t, m)
	s = NewStreamableSlice(&[]string{"a", "b", "c"})
	m = s.NoneMatch(p)
	require.True(t, m)
}

func TestStreamableSlice_NthMatch(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "F", "g", "H", "i", "E", "a", "B", "c"})
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

	o = s.NthMatch(nil, 10)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "c", v)
	o = s.NthMatch(nil, 11)
	require.False(t, o.IsPresent())
	o = s.NthMatch(nil, -10)
	v, _ = o.GetOk()
	require.Equal(t, "d", v)
	require.True(t, o.IsPresent())
	o = s.NthMatch(nil, -11)
	require.False(t, o.IsPresent())

	s = NewStreamableSlice(&[]string{"a", "b", "c"})
	o = s.NthMatch(p, 1)
	require.False(t, o.IsPresent())
}

func TestStreamableSlice_Reverse(t *testing.T) {
	s := NewStreamableSlice(&[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	s2 := s.Reverse()
	require.Equal(t, 10, s2.Len())
	o := s2.NthMatch(nil, 1)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "10", v)
	o = s2.NthMatch(nil, 10)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "1", v)
}

func TestStreamableSlice_Skip(t *testing.T) {
	s := NewStreamableSlice(&[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	s2 := s.Skip(5)
	require.Equal(t, 5, s2.Len())

	s2 = s.Skip(10)
	require.Equal(t, 0, s2.Len())

	s2 = s.Skip(-1)
	require.Equal(t, 10, s2.Len())
}

func TestStreamableSlice_SkipAndLimit(t *testing.T) {
	s := NewStreamableSlice(&[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	s2 := s.Skip(5).Limit(10)
	require.Equal(t, 5, s2.Len())
	s2 = s.Skip(5).Limit(2)
	require.Equal(t, 2, s2.Len())
}

func TestStreamableSlice_Slice(t *testing.T) {
	s := NewStreamableSlice(&[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	s2 := s.Slice(5, 3)
	require.Equal(t, 3, s2.Len())
	o := s2.NthMatch(nil, 1)
	require.True(t, o.IsPresent())
	v, _ := o.GetOk()
	require.Equal(t, "5", v)
	o = s2.NthMatch(nil, -1)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "7", v)

	s2 = s.Slice(5, -3)
	require.Equal(t, 3, s2.Len())
	o = s2.NthMatch(nil, 1)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "2", v)
	o = s2.NthMatch(nil, -1)
	require.True(t, o.IsPresent())
	v, _ = o.GetOk()
	require.Equal(t, "4", v)

	s2 = s.Slice(-10, -3)
	require.Equal(t, 0, s2.Len())

	s2 = s.Slice(20, -10)
	require.Equal(t, 0, s2.Len())

	s2 = s.Slice(10, -10)
	require.Equal(t, 10, s2.Len())
}

func TestStreamableSlice_Sorted(t *testing.T) {
	s := NewStreamableSlice(&[]string{"d", "j", "f", "g", "h", "i", "e", "a", "b", "c"})
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

func TestStreamableSlice_SymmetricDifference(t *testing.T) {
	s1 := NewStreamableSlice(&[]string{"a", "b", "c"})
	s2 := NewStreamableSlice(&[]string{"b", "c", "d"})
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

	s3 := s1.SymmetricDifference(s2, nil)
	require.Equal(t, 0, s3.Len())
}

func TestStreamableSlice_Union(t *testing.T) {
	s1 := NewStreamableSlice(&[]string{"a", "b", "c"})
	s2 := NewStreamableSlice(&[]string{"b", "c", "d"})
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

	s3 := s1.Union(s2, nil)
	require.Equal(t, 0, s3.Len())
}

func TestStreamableSlice_Unique(t *testing.T) {
	s := NewStreamableSlice(&[]string{"a", "a", "b", "c", "c"})
	s2 := s.Unique(StringComparator)
	require.Equal(t, 3, s2.Len())

	s3 := NewStreamableSlice(&[]instruct{{1}, {1}, {2}, {3}, {3}})
	require.Equal(t, 3, s3.Distinct().Len())
	require.Equal(t, 3, s3.Unique(nil).Len())

	s4 := NewStreamableSlice(&[]*instruct{{1}, {1}, {2}, {3}, {3}, {1}})
	require.Equal(t, 6, s4.Distinct().Len())
	s5 := s4.Unique(NewComparator(func(v1, v2 *instruct) int {
		if v1.value < v2.value {
			return -1
		} else if v1.value > v2.value {
			return 1
		}
		return 0
	}))
	require.Equal(t, 3, s5.Len())
}
