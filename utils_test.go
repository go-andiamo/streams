package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
