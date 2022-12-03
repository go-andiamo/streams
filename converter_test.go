package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConverter(t *testing.T) {
	c := NewConverter[string, string](func(v string) (string, error) {
		return v, nil
	})
	require.NotNil(t, c)

	c = NewConverter[string, string](nil)
	require.Nil(t, c)
}

func TestConverter(t *testing.T) {
	c := NewConverter[string, outStruct](func(v string) (outStruct, error) {
		return outStruct{value: v}, nil
	})
	out, err := c.Convert("a")
	require.NoError(t, err)
	require.Equal(t, "a", out.value)
}

func TestConverterFunc(t *testing.T) {
	c := ConverterFunc[string, outStruct](func(v string) (outStruct, error) {
		return outStruct{value: v}, nil
	})
	out, err := c.Convert("a")
	require.NoError(t, err)
	require.Equal(t, "a", out.value)
}
