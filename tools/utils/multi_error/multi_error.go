package multi_error

import "github.com/aiagt/aiagt/tools/utils/logger"

type RunError struct {
	err error
}

func NewRunError(err error) *RunError {
	if err == nil {
		return nil
	}
	return &RunError{err: err}
}

func (e *RunError) Expect(msg string) {
	if e == nil || e.err == nil {
		return
	}
	logger.Errorf("%s: %v", msg, e.err)
}

type MultiError struct {
	err error
}

func NewMultiError() *MultiError {
	return &MultiError{}
}

func (m *MultiError) Run(fn func() error) *RunError {
	if m.err == nil {
		m.err = fn()
	}
	return NewRunError(m.err)
}

type MultiError1[T any] struct {
	err error
}

func NewMultiError1[T any]() *MultiError1[T] {
	return &MultiError1[T]{}
}

func (m *MultiError1[T]) Run(fn func(T) error, arg1 T) *RunError {
	if m.err == nil {
		m.err = fn(arg1)
	}
	return NewRunError(m.err)
}

type MultiError2[T any, E any] struct {
	err error
}

func NewMultiError2[T any, E any]() *MultiError2[T, E] {
	return &MultiError2[T, E]{}
}

func (m *MultiError2[T, E]) Run(fn func(T, E) error, arg1 T, arg2 E) *RunError {
	if m.err == nil {
		m.err = fn(arg1, arg2)
	}
	return NewRunError(m.err)
}
