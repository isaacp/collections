package stack

type (
	stack[T any] struct {
		store []T
	}
)

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		store: make([]T, 0),
	}
}

func (s stack[T]) Height() int { return len(s.store) }

func (s *stack[T]) Pop() T {
	top := s.store[0]
	s.store = s.store[1:]
	return top
}

func (s *stack[T]) Push(t T) {
	s.store = append([]T{t}, s.store...)
}

func (s *stack[T]) Peek() T { return s.store[0] }
