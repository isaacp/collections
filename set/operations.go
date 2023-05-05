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

	intersection := sets[0]
	remainingSets := sets[1:]
	for i := 0; i < len(remainingSets); i++ {
		for _, member := range remainingSets[i].members {
			if !intersection.Contains(*member) {
				intersection.Remove(*member)
			}
		}
	}

	return intersection
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
