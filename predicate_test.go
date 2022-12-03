package streams

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewPredicate(t *testing.T) {
	called := false
	fp := NewPredicate[string](func(v string) bool {
		called = true
		return false
	})
	r := fp.Test("a")
	require.False(t, r)
	require.True(t, called)

	fp = NewPredicate[string](nil)
	r = fp.Test("a")
	require.True(t, r)
}

func TestPredicate_Negate(t *testing.T) {
	called := false
	fp := NewPredicate[string](func(v string) bool {
		called = true
		return false
	})
	nfp := fp.Negate()
	r := nfp.Test("a")
	require.True(t, r)
	require.True(t, called)
}

func TestPredicate_And(t *testing.T) {
	firstCalled := false
	secondCalled := false
	tp2 := NewPredicate[string](func(v string) bool {
		secondCalled = true
		return true
	})
	tp := NewPredicate[string](func(v string) bool {
		firstCalled = true
		return true
	}).And(tp2)
	r := tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.True(t, secondCalled)

	firstCalled = false
	secondCalled = false
	tp = NewPredicate[string](func(v string) bool {
		firstCalled = true
		return false
	}).And(tp2)
	r = tp.Test("a")
	require.False(t, r)
	require.True(t, firstCalled)
	require.False(t, secondCalled)
}

func TestPredicate_Or(t *testing.T) {
	firstCalled := false
	secondCalled := false
	tp2 := NewPredicate[string](func(v string) bool {
		secondCalled = true
		return true
	})
	tp := NewPredicate[string](func(v string) bool {
		firstCalled = true
		return false
	}).Or(tp2)
	r := tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.True(t, secondCalled)

	firstCalled = false
	secondCalled = false
	tp = NewPredicate[string](func(v string) bool {
		firstCalled = true
		return true
	}).Or(tp2)
	r = tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.False(t, secondCalled)
}

func TestPredicate_AndOr(t *testing.T) {
	callCount := 0
	fp := NewPredicate[string](func(v string) bool {
		callCount++
		return false
	})
	tp := NewPredicate[string](func(v string) bool {
		callCount++
		return true
	})
	p := fp.And(fp).Or(tp)
	r := p.Test("a")
	require.True(t, r)
	require.Equal(t, 2, callCount)

	p = p.Negate()
	r = p.Test("a")
	require.False(t, r)
	require.Equal(t, 4, callCount)
}

func TestPredicate_OrAnd(t *testing.T) {
	callCount := 0
	fp := NewPredicate[string](func(v string) bool {
		callCount++
		return false
	})
	tp := NewPredicate[string](func(v string) bool {
		callCount++
		return true
	})
	p := fp.Or(tp).And(tp)
	r := p.Test("a")
	require.True(t, r)
	require.Equal(t, 3, callCount)

	p = p.Negate()
	r = p.Test("a")
	require.False(t, r)
	require.Equal(t, 6, callCount)
}

func TestPredicateFunc(t *testing.T) {
	p := PredicateFunc[string](func(v string) bool {
		return strings.ToUpper(v) == v
	})
	require.True(t, p.Test("A"))
	require.False(t, p.Test("a"))
}

func TestPredicateFunc_Negate(t *testing.T) {
	called := false
	fp := PredicateFunc[string](func(v string) bool {
		called = true
		return false
	})
	nfp := fp.Negate()
	r := nfp.Test("a")
	require.True(t, r)
	require.True(t, called)
}

func TestPredicateFunc_And(t *testing.T) {
	firstCalled := false
	secondCalled := false
	tp2 := PredicateFunc[string](func(v string) bool {
		secondCalled = true
		return true
	})
	tp := PredicateFunc[string](func(v string) bool {
		firstCalled = true
		return true
	}).And(tp2)
	r := tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.True(t, secondCalled)

	firstCalled = false
	secondCalled = false
	tp = PredicateFunc[string](func(v string) bool {
		firstCalled = true
		return false
	}).And(tp2)
	r = tp.Test("a")
	require.False(t, r)
	require.True(t, firstCalled)
	require.False(t, secondCalled)
}

func TestPredicateFunc_Or(t *testing.T) {
	firstCalled := false
	secondCalled := false
	tp2 := PredicateFunc[string](func(v string) bool {
		secondCalled = true
		return true
	})
	tp := PredicateFunc[string](func(v string) bool {
		firstCalled = true
		return false
	}).Or(tp2)
	r := tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.True(t, secondCalled)

	firstCalled = false
	secondCalled = false
	tp = PredicateFunc[string](func(v string) bool {
		firstCalled = true
		return true
	}).Or(tp2)
	r = tp.Test("a")
	require.True(t, r)
	require.True(t, firstCalled)
	require.False(t, secondCalled)
}

func TestPredicateFunc_AndOr(t *testing.T) {
	callCount := 0
	fp := PredicateFunc[string](func(v string) bool {
		callCount++
		return false
	})
	tp := PredicateFunc[string](func(v string) bool {
		callCount++
		return true
	})
	p := fp.And(fp).Or(tp)
	r := p.Test("a")
	require.True(t, r)
	require.Equal(t, 2, callCount)

	p = p.Negate()
	r = p.Test("a")
	require.False(t, r)
	require.Equal(t, 4, callCount)
}

func TestPredicateFunc_OrAnd(t *testing.T) {
	callCount := 0
	fp := PredicateFunc[string](func(v string) bool {
		callCount++
		return false
	})
	tp := PredicateFunc[string](func(v string) bool {
		callCount++
		return true
	})
	p := fp.Or(tp).And(tp)
	r := p.Test("a")
	require.True(t, r)
	require.Equal(t, 3, callCount)

	p = p.Negate()
	r = p.Test("a")
	require.False(t, r)
	require.Equal(t, 6, callCount)
}
