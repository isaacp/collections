package set

import (
	"fmt"
	"hash/fnv"
)

type (
	Set[T any] struct {
		members map[int]*T
	}
)

func (set *Set[T]) Add(members ...T) {
	if set.members == nil {
		set.members = make(map[int]*T)
	}
	for _, member := range members {
		id := hash(fmt.Sprintf("%v", member))
		set.members[id] = &member
	}
}

func (set *Set[T]) Remove(members ...T) {
	for _, member := range members {
		delete(set.members, hash(fmt.Sprintf("%v", member)))
	}
}

func (set Set[T]) Contains(Id int) bool {
	return set.members[Id] != nil
}

func (set Set[T]) ToSlice() []T {
	slice := make([]T, 0)
	for _, member := range set.members {
		slice = append(slice, *member)
	}
	return slice
}

func (set Set[T]) Count() int {
	return len(set.members)
}

func hash(input string) int {
	id, _ := fnv.New128().Write([]byte(input))
	return id
}
