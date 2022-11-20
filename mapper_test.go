package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type outStruct struct {
	value string
}

func TestMapper_Map(t *testing.T) {
	s := Of("D", "j", "F", "g", "H", "i", "E", "a", "B", "c")
	c := NewConverter[string, outStruct](func(v string) outStruct {
		return outStruct{value: v}
	})
	m := NewMapper(c)
	out := m.Map(s)
	require.Equal(t, 10, out.Len())
	first := out.FirstMatch(nil)
	require.True(t, first.IsPresent())
	v, ok := first.GetOk()
	require.True(t, ok)
	require.Equal(t, "D", v.value)
}
