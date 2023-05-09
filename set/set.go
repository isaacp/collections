package set

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	for _, m := range members {
		set.members[hash(m)] = &m
	}
}

func (set *Set[T]) Remove(members ...T) {
	for _, m := range members {
		delete(set.members, hash(m))
	}
}

func (set Set[T]) Contains(member T) bool {
	return set.members[hash(member)] != nil
}

func (set Set[T]) ToSlice() []T {
	s := make([]T, 0)
	for _, m := range set.members {
		s = append(s, *m)
	}
	return s
}

func (set Set[T]) Count() int {
	return len(set.members)
}

func (set Set[T]) Map(action func(member T) T) Set[T] {
	s := Set[T]{}
	for _, m := range set.members {
		s.Add(action(*m))
	}
	return s
}

func hash[T any](member T) string {
	s, _ := json.Marshal(member)
	b := sha256.Sum256(s)
	return hex.EncodeToString(b[:])
}

func (s Set[T]) Copy() Set[T] {
	n := NewSet(s.ToSlice()...)
	return *n
}

// Set Operations
func Difference[T any](first Set[T], second Set[T]) Set[T] {
	result := EmptySet[T]()
	for _, member := range first.members {
		if !second.Contains(*member) {
			result.Add(*member)
		}
	}
	return result
}

func Disjoint[T any](sets ...Set[T]) bool {
	r := EmptySet[T]()
	c := 0
	for _, s := range sets {
		c += s.Count()
		r.Add(s.ToSlice()...)
	}
	return r.Count() == c
}

func EmptySet[T any]() Set[T] {
	return Set[T]{
		members: make(map[string]*T),
	}
}

func Intersection[T any](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return EmptySet[T]()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	head := sets[0].Copy()
	deletions := EmptySet[T]()
	for _, next := range sets[1:] {
		headMembers := head.ToSlice()
		for i := 0; i < head.Count(); i++ {
			if !next.Contains(headMembers[i]) {
				deletions.Add(headMembers[i])
			}
		}
		for _, member := range deletions.ToSlice() {
			head.Remove(member)
		}
	}
	return head
}

func NewSet[T any](members ...T) *Set[T] {
	set := EmptySet[T]()
	for _, member := range members {
		set.Add(member)
	}
	return &set
}

func Union[T any](sets ...Set[T]) Set[T] {
	result := EmptySet[T]()
	for _, set := range sets {
		result.Add(set.ToSlice()...)
	}
	return result
}
