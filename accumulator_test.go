package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAccumulator(t *testing.T) {
	a := NewAccumulator[int, int](func(t int, u int) int {
		return t + u
	})
	require.NotNil(t, a)

	a = NewAccumulator[int, int](nil)
	require.Nil(t, a)
}

func TestAccumulator_Apply(t *testing.T) {
	a := NewAccumulator[int, int](func(t int, u int) int {
		return t + u
	})
	r := a.Apply(1, 2)
	require.Equal(t, 3, r)
}

func TestAccumulatorFunc_Apply(t *testing.T) {
	a := AccumulatorFunc[int, int](func(t int, u int) int {
		return t + u
	})
	r := a.Apply(1, 2)
	require.Equal(t, 3, r)
}
