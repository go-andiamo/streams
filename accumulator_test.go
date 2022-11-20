package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccumulator(t *testing.T) {
	a := NewAccumulator[int, int](func(t int, u int) int {
		return t + u
	})
	r := a.Apply(1, 2)
	require.Equal(t, 3, r)
}
