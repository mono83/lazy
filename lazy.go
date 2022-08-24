package lazy

import "sync"

// Const returns lazy value supplier with constant value.
func Const[T any](value T) func() T {
	return func() T { return value }
}

// ConstE returns lazy value supplier with constant value and no error.
func ConstE[T any](value T) func() (T, error) {
	return func() (T, error) {
		return value, nil
	}
}

// Error returns lazy value supplier with constant error and no value.
func Error[T any](err error) func() (T, error) {
	return func() (t T, e error) {
		e = err
		return
	}
}

// New constructs lazy value supplier built over given supplier
// function invoked only once.
func New[T any](supply func() T) func() T {
	var value T
	var once sync.Once
	return func() T {
		once.Do(func() { value = supply() })
		return value
	}
}

// NewE constructs lazy value supplier built over given supplier
// function invoked only once.
func NewE[T any](supply func() (T, error)) func() (T, error) {
	var value T
	var err error
	var once sync.Once
	return func() (T, error) {
		once.Do(func() { value, err = supply() })
		return value, err
	}
}
