package streams

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestIsDistinctable(t *testing.T) {
	testCases := []struct {
		value  any
		expect bool
	}{
		{
			"",
			true,
		},
		{
			1,
			true,
		},
		{
			int8(1),
			true,
		},
		{
			int16(1),
			true,
		},
		{
			int32(1),
			true,
		},
		{
			int64(1),
			true,
		},
		{
			uint(1),
			true,
		},
		{
			uint8(1),
			true,
		},
		{
			uint16(1),
			true,
		},
		{
			uint32(1),
			true,
		},
		{
			uint64(1),
			true,
		},
		{
			true,
			true,
		},
		{
			float32(1.23),
			true,
		},
		{
			float64(1.23),
			true,
		},
		{
			struct{}{},
			true,
		},
		{
			&struct{}{},
			false,
		},
		{
			func() {},
			true,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			require.Equal(t, tc.expect, isDistinctable(tc.value))
		})
	}
}

func TestAbsInt(t *testing.T) {
	require.Equal(t, 0, absInt(0))
	require.Equal(t, 1, absInt(1))
	require.Equal(t, 1, absInt(-1))
}

func TestAbsZero(t *testing.T) {
	require.Equal(t, 0, absZero(0))
	require.Equal(t, 1, absZero(1))
	require.Equal(t, 0, absZero(-1))
}

func TestJoinPredicates(t *testing.T) {
	ps := []Predicate[string]{
		nil,
		NewPredicate(func(v string) bool {
			return v == "a"
		}),
		NewPredicate(func(v string) bool {
			return v == "b"
		}),
		NewPredicate(func(v string) bool {
			return v == "c"
		}),
		nil,
	}
	p := joinPredicates(ps...)
	require.NotNil(t, p)
	require.True(t, p.Test("a"))
	require.True(t, p.Test("b"))
	require.True(t, p.Test("c"))
	require.False(t, p.Test("d"))

	p = joinPredicates[string]()
	require.Nil(t, p)
}

func TestSliceIterator(t *testing.T) {
	next := SliceIterator([]string{"a", "b", "c", "d"})
	result := ""
	count := 0
	for v, ok := next(); ok; v, ok = next() {
		result += v
		count++
	}
	require.Equal(t, 4, count)
	require.Equal(t, "abcd", result)
}

func TestSliceIteratorWithPredicates(t *testing.T) {
	upper := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == v
	})
	isA := NewPredicate(func(v string) bool {
		return strings.ToUpper(v) == "A"
	})
	next := SliceIterator([]string{"a", "B", "c", "D", "e", "F"}, upper, isA)
	result := ""
	count := 0
	for v, ok := next(); ok; v, ok = next() {
		result += v
		count++
	}
	require.Equal(t, 4, count)
	require.Equal(t, "aBDF", result)
}

func TestStringComparator(t *testing.T) {
	c := StringComparator
	require.Equal(t, 0, c.Compare("a", "a"))
	require.Equal(t, -1, c.Compare("a", "b"))
	require.Equal(t, 1, c.Compare("b", "a"))
}

func TestStringInsensitiveComparator(t *testing.T) {
	c := StringInsensitiveComparator
	require.Equal(t, 0, c.Compare("a", "a"))
	require.Equal(t, -1, c.Compare("a", "b"))
	require.Equal(t, 1, c.Compare("b", "a"))
	require.Equal(t, 0, c.Compare("a", "A"))
	require.Equal(t, -1, c.Compare("a", "B"))
	require.Equal(t, 1, c.Compare("B", "a"))
	require.Equal(t, -1, c.Compare("A", "b"))
	require.Equal(t, 1, c.Compare("b", "A"))
}

func TestIntComparator(t *testing.T) {
	c := IntComparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestInt8Comparator(t *testing.T) {
	c := Int8Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestInt16Comparator(t *testing.T) {
	c := Int16Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestInt32Comparator(t *testing.T) {
	c := Int32Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestInt64Comparator(t *testing.T) {
	c := Int64Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestUintComparator(t *testing.T) {
	c := UintComparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestUint8Comparator(t *testing.T) {
	c := Uint8Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestUint16Comparator(t *testing.T) {
	c := Uint16Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestUint32Comparator(t *testing.T) {
	c := Uint32Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestUint64Comparator(t *testing.T) {
	c := Uint64Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestFloat32Comparator(t *testing.T) {
	c := Float32Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}

func TestFloat64Comparator(t *testing.T) {
	c := Float64Comparator
	require.Equal(t, 0, c.Compare(1, 1))
	require.Equal(t, -1, c.Compare(1, 2))
	require.Equal(t, 1, c.Compare(2, 1))
}
