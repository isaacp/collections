package list

import (
	"errors"
)

type (
	list[T comparable] struct {
		members[T]
	}

	members[T comparable] []T
)

func List[T comparable](members []T) *list[T] {
	return &list[T]{
		members: members,
	}
}

func (c list[T]) Contains(item T) bool {
	for _, member := range c.members {
		if member == item {
			return true
		}
	}
	return false
}

func (c *list[T]) Add(item T) {
	c.members = append(c.members, item)
}

func (c *list[T]) Remove(item T) {
	index, err := c.Index(item)
	if err != nil {
		return
	}
	oneOver := index + 1
	c.members = append(c.members[:index], c.members[oneOver:]...)
}

func (c list[T]) Length() int {
	return len(c.members)
}

func (c list[T]) Index(item T) (int, error) {
	for index, itm := range c.members {
		if itm == item {
			return index, nil
		}
	}

	return -1, errors.New("item not found")
}

func (c list[T]) Map(action func(member T) T) list[T] {
	l := list[T]{}
	for _, m := range l.members {
		l.Add(action(m))
	}
	return l
}
