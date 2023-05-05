package set

func Difference[T any](first Set[T], second Set[T]) Set[T] {
	result := EmptySet[T]()
	for _, member := range first.members {
		if !second.Contains(*member) {
			result.Add(*member)
		}
	}
	return result
}

func Disjoint[T any](first Set[T], second Set[T]) bool {
	result := true
	for _, member := range first.members {
		if second.Contains(*member) {
			return false
		}
	}
	return result
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

	last := sets[len(sets)-1]
	nextToLast := sets[len(sets)-2]
	front := sets[:len(sets)-2]
	lastMembers := last.ToSlice()
	for i := 0; i < len(last.members); i++ {
		if !nextToLast.Contains(lastMembers[i]) {
			last.Remove(lastMembers[i])
		}
	}

	return Intersection(append(front, last)...)
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
