package stack

import (
	"errors"
)

type (
	stack[T any] struct {
		store []T
	}
)

var stackEmpty = "stack is empty"

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		store: make([]T, 0),
	}
}

func (s stack[T]) Height() int { return len(s.store) }

func (s stack[T]) IsEmpty() bool { return len(s.store) == 0 }

func (s stack[T]) Peek() (*T, error) {
	if !s.IsEmpty() {
		return &s.store[0], nil
	}
	return nil, errors.New(stackEmpty)
}

func (s *stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New(stackEmpty)
	}
	top := s.store[0]
	s.store = s.store[1:]
	return &top, nil
}

func (s *stack[T]) Push(t T) {
	s.store = append([]T{t}, s.store...)
}
