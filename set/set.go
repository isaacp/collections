package set

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type (
	Set[T any] struct {
		members map[string]*T
	}
)

func (set *Set[T]) Add(members ...T) {
	if set.members == nil {
		set.members = make(map[string]*T)
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

func (set Set[T]) Contains(Id string) bool {
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

func hash(input string) string {
	hashBytes := sha256.Sum256([]byte(input))
	id := hex.EncodeToString(hashBytes[:])
	return id
}
