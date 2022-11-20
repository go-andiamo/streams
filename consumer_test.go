package streams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	collected := ""
	called := false
	c := NewConsumer[string](func(v string) {
		called = true
		collected = v
	})
	c.Accept("a")
	require.Equal(t, "a", collected)
	require.True(t, called)
}

func TestConsumer_AndThen(t *testing.T) {
	collected := ""
	calledCount := 0
	c := NewConsumer[string](func(v string) {
		calledCount++
		collected = v
	})
	c = c.AndThen(c).AndThen(c)
	c.Accept("a")
	require.Equal(t, "a", collected)
	require.Equal(t, 3, calledCount)
}
