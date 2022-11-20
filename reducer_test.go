package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type instruct struct {
	value int
}

func TestReducer(t *testing.T) {
	a := NewAccumulator[instruct, int](func(t instruct, r int) int {
		return r + t.value
	})
	r := NewReducer(a)
	s := Of(instruct{1}, instruct{2}, instruct{3}, instruct{4})
	tot := r.Reduce(s)
	require.Equal(t, 10, tot)
}
