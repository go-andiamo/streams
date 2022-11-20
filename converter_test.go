package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConverter(t *testing.T) {
	c := NewConverter[string, outStruct](func(v string) (outStruct, error) {
		return outStruct{value: v}, nil
	})
	out, err := c.Convert("a")
	require.NoError(t, err)
	require.Equal(t, "a", out.value)
}
