package set

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
		delete(set.members, hash(member))
	}
}

func (set Set[T]) Contains(member T) bool {
	return set.members[hash(member)] != nil
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

func (set Set[T]) Map(action func(member T) T) Set[T] {
	s := Set[T]{}
	for _, member := range set.members {
		s.Add(action(*member))
	}
	return s
}

func hash[T any](member T) string {
	series, _ := json.Marshal(member)
	hashBytes := sha256.Sum256(series)
	return hex.EncodeToString(hashBytes[:])
}
