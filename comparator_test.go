package streams

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestComparator_Compare(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.Equal(t, 0, c.Compare("a", "a"))
	require.Equal(t, -1, c.Compare("a", "b"))
	require.Equal(t, 1, c.Compare("b", "a"))
}

func TestComparator_Reversed(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	}).Reversed()
	require.Equal(t, 0, c.Compare("a", "a"))
	require.Equal(t, 1, c.Compare("a", "b"))
	require.Equal(t, -1, c.Compare("b", "a"))
}

func TestComparator_Less(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.False(t, c.Less("a", "a"))
	require.True(t, c.Less("a", "b"))
	require.False(t, c.Less("b", "a"))

	c = c.Reversed()
	require.False(t, c.Less("a", "a"))
	require.False(t, c.Less("a", "b"))
	require.True(t, c.Less("b", "a"))
}

func TestComparator_LessOrEqual(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.True(t, c.LessOrEqual("a", "a"))
	require.True(t, c.LessOrEqual("a", "b"))
	require.False(t, c.LessOrEqual("b", "a"))

	c = c.Reversed()
	require.True(t, c.LessOrEqual("a", "a"))
	require.False(t, c.LessOrEqual("a", "b"))
	require.True(t, c.LessOrEqual("b", "a"))
}

func TestComparator_Greater(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.False(t, c.Greater("a", "a"))
	require.False(t, c.Greater("a", "b"))
	require.True(t, c.Greater("b", "a"))

	c = c.Reversed()
	require.False(t, c.Greater("a", "a"))
	require.True(t, c.Greater("a", "b"))
	require.False(t, c.Greater("b", "a"))
}

func TestComparator_GreaterOrEqual(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.True(t, c.GreaterOrEqual("a", "a"))
	require.False(t, c.GreaterOrEqual("a", "b"))
	require.True(t, c.GreaterOrEqual("b", "a"))

	c = c.Reversed()
	require.True(t, c.GreaterOrEqual("a", "a"))
	require.True(t, c.GreaterOrEqual("a", "b"))
	require.False(t, c.GreaterOrEqual("b", "a"))
}

func TestComparator_Equal(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.True(t, c.Equal("a", "a"))
	require.False(t, c.Equal("a", "b"))
	require.False(t, c.Equal("b", "a"))

	c = c.Reversed()
	require.False(t, c.Equal("a", "a"))
	require.True(t, c.Equal("a", "b"))
	require.True(t, c.Equal("b", "a"))
}

func TestComparator_NotEqual(t *testing.T) {
	c := NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	require.False(t, c.NotEqual("a", "a"))
	require.True(t, c.NotEqual("a", "b"))
	require.True(t, c.NotEqual("b", "a"))

	c = c.Reversed()
	require.True(t, c.NotEqual("a", "a"))
	require.False(t, c.NotEqual("a", "b"))
	require.False(t, c.NotEqual("b", "a"))
}

type testComparable struct {
	primary   string
	secondary int
}

func TestComparator_Then(t *testing.T) {
	c := NewComparator[testComparable](func(v1, v2 testComparable) int {
		return strings.Compare(v1.primary, v2.primary)
	})
	csub := NewComparator[testComparable](func(v1, v2 testComparable) int {
		if v1.secondary < v2.secondary {
			return -1
		} else if v1.secondary > v2.secondary {
			return 1
		}
		return 0
	})
	ct := c.Then(csub)

	a0 := testComparable{primary: "a", secondary: 0}
	a1 := testComparable{primary: "a", secondary: 1}
	b0 := testComparable{primary: "b", secondary: 0}
	b1 := testComparable{primary: "b", secondary: 1}
	testCases := []struct {
		first  testComparable
		second testComparable
		expect int
	}{
		{
			a0,
			a0,
			0,
		},
		{
			a0,
			a1,
			-1,
		},
		{
			a1,
			a0,
			1,
		},
		{
			b0,
			b0,
			0,
		},
		{
			b0,
			b1,
			-1,
		},
		{
			b1,
			b0,
			1,
		},
		{
			a0,
			b0,
			-1,
		},
		{
			b0,
			a0,
			1,
		},
		{
			a0,
			b1,
			-1,
		},
		{
			b1,
			a0,
			1,
		},
		{
			a1,
			b1,
			-1,
		},
		{
			b1,
			a1,
			1,
		},
		{
			a1,
			b0,
			-1,
		},
		{
			b0,
			a1,
			1,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			require.Equal(t, tc.expect, ct.Compare(tc.first, tc.second))
		})
	}
}
