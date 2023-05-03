package set

func Union[T any](sets ...Set[T]) Set[T] {
	result := EmptySet[T]()
	for _, set := range sets {
		result.Add(set.ToSlice()...)
	}
	return result
}

func Intersection[T any](first Set[T], sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return first
	}

	result := EmptySet[T]()
	for _, set := range sets {
		for key, member := range first.members {
			if set.Contains(key) {
				result.Add(*member)
			}
		}
	}
	return result
}

func Difference[T any](first Set[T], second Set[T]) Set[T] {
	result := EmptySet[T]()
	for key, member := range first.members {
		if !second.Contains(key) {
			result.Add(*member)
		}
	}
	return result
}

func EmptySet[T any]() Set[T] {
	return Set[T]{
		members: make(map[int]*T),
	}
}

func NewSet[T any](members ...T) *Set[T] {
	set := EmptySet[T]()
	for _, member := range members {
		set.Add(member)
	}
	return &set
}
