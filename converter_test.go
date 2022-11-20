package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConverter(t *testing.T) {
	c := NewConverter[string, outStruct](func(v string) outStruct {
		return outStruct{value: v}
	})
	out := c.Convert("a")
	require.Equal(t, "a", out.value)
}
