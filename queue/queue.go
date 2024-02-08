package queue

import (
	"errors"
	"fmt"
)

type (
	Queue[T any] struct {
		members []T
	}

	DequeueOutput[T any] struct {
		Value T
		Error error
	}
)

func New[T any]() Queue[T] {
	return Queue[T]{
		members: make([]T, 0),
	}
}

func FromArray[T any](array []T) Queue[T] {
	q := New[T]()
	for _, val := range array {
		q.Enqueue(val)
	}
	return q
}

func (q *Queue[T]) Enqueue(item T) {
	q.members = append(q.members, item)
}

func (q *Queue[T]) Dequeue() DequeueOutput[T] {
	if len(q.members) == 0 {
		return DequeueOutput[T]{
			Error: errors.New("queue is empty"),
		}
	}
	t := q.members[0]
	if len(q.members) > 1 {
		q.members = q.members[1:]
	} else {
		q.members = []T{}
	}
	return DequeueOutput[T]{
		Value: t,
		Error: nil,
	}
}

func (q Queue[T]) Print() {
	fmt.Println(q.members)
}
