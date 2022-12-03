package streams

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type outStruct struct {
	value string
}

func TestNewMapperPanics(t *testing.T) {
	require.Panics(t, func() {
		NewMapper[string, string](nil)
	})
}

func TestMapper_Map(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	c := NewConverter[string, outStruct](func(v string) (outStruct, error) {
		return outStruct{value: v}, nil
	})
	m := NewMapper(c)
	out, err := m.Map(s)
	require.NoError(t, err)
	require.Equal(t, 10, out.Len())
	first := out.FirstMatch(nil)
	require.True(t, first.IsPresent())
	v, ok := first.GetOk()
	require.True(t, ok)
	require.Equal(t, "D", v.value)
}

func TestMapper_Map_Error(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	c := NewConverter[string, outStruct](func(v string) (outStruct, error) {
		return outStruct{}, errors.New("whoops")
	})
	m := NewMapper(c)
	_, err := m.Map(s)
	require.Error(t, err)
	require.Equal(t, "whoops", err.Error())
}
