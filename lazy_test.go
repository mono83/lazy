package lazy

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConst(t *testing.T) {
	assert.Equal(t, "foobar", Const("foobar")())
	assert.Equal(t, 15, Const(15)())
	assert.Nil(t, Const[error](nil)())
}

func TestConstE(t *testing.T) {
	if x, err := ConstE(3)(); assert.NoError(t, err) {
		assert.Equal(t, 3, x)
	}
	if x, err := ConstE("foo")(); assert.NoError(t, err) {
		assert.Equal(t, "foo", x)
	}
	if x, err := ConstE[error](nil)(); assert.NoError(t, err) { // Error as value
		assert.Nil(t, x)
	}
}

func TestError(t *testing.T) {
	if x, err := Error[int](nil)(); assert.NoError(t, err) {
		assert.Equal(t, 0, x)
	}
	if x, err := Error[string](errors.New("foo"))(); assert.Error(t, err) {
		assert.Equal(t, "", x)
		assert.Equal(t, errors.New("foo"), err)
	}
}

func TestNew(t *testing.T) {
	cnt := 0

	lazyInt := New(func() int {
		cnt++
		return 42
	})
	assert.Equal(t, 0, cnt) // No invocations
	assert.Equal(t, 42, lazyInt())
	assert.Equal(t, 1, cnt)
	assert.Equal(t, 42, lazyInt())
	assert.Equal(t, 1, cnt) // Still one
}

func TestNewE(t *testing.T) {
	cnt := 0

	lazyIntOk := NewE(func() (int, error) {
		cnt++
		return 42, nil
	})
	lazyIntErr := NewE(func() (int, error) {
		cnt++
		return 0, errors.New("foo")
	})
	assert.Equal(t, 0, cnt) // No invocations
	if value, err := lazyIntOk(); assert.NoError(t, err) {
		assert.Equal(t, 42, value)
		assert.Equal(t, 1, cnt)
	}
	if value, err := lazyIntOk(); assert.NoError(t, err) {
		assert.Equal(t, 42, value)
		assert.Equal(t, 1, cnt) // Still one
	}
	if _, err := lazyIntErr(); assert.Error(t, err) {
		assert.Equal(t, 2, cnt)
	}
	if _, err := lazyIntErr(); assert.Error(t, err) {
		assert.Equal(t, 2, cnt) // Still two
	}
}
