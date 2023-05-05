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

	head := sets[0]
	neck := sets[1]
	tail := sets[1:]

	for i := 0; i < len(head.members); i++ {
		common := head.ToSlice()
		if !neck.Contains(common[i]) {
			head.Remove(common[i])
		}
	}

	return Intersection(append(tail, head)...)
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
