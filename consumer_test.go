package streams

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	collected := ""
	called := false
	c := NewConsumer[string](func(v string) error {
		called = true
		collected = v
		return nil
	})
	err := c.Accept("a")
	require.NoError(t, err)
	require.Equal(t, "a", collected)
	require.True(t, called)
}

func TestConsumer_AndThen(t *testing.T) {
	collected := ""
	calledCount := 0
	c := NewConsumer[string](func(v string) error {
		calledCount++
		collected = v
		return nil
	})
	c = c.AndThen(c).AndThen(c)
	err := c.Accept("a")
	require.NoError(t, err)
	require.Equal(t, "a", collected)
	require.Equal(t, 3, calledCount)
}

func TestConsumer_Error(t *testing.T) {
	calledCount := 0
	c := NewConsumer[string](func(v string) error {
		calledCount++
		return errors.New("whoops")
	})
	err := c.Accept("a")
	require.Error(t, err)
	require.Equal(t, "whoops", err.Error())
	require.Equal(t, 1, calledCount)

	calledCount = 0
	c = NewConsumer[string](func(v string) error {
		calledCount++
		return nil
	}).AndThen(c)
	err = c.Accept("a")
	require.Error(t, err)
	require.Equal(t, "whoops", err.Error())
	require.Equal(t, 2, calledCount)
}
