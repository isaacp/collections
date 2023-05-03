package collections

import (
	"errors"
)

type (
	List[T comparable] struct {
		members[T]
	}

	members[T comparable] []T
)

func (c List[T]) Contains(item T) bool {
	for _, member := range c.members {
		if member == item {
			return true
		}
	}
	return false
}

func (c *List[T]) Add(item T) {
	c.members = append(c.members, item)
}

func (c *List[T]) Remove(item T) {
	index, err := c.Index(item)
	if err != nil {
		return
	}
	oneOver := index + 1
	c.members = append(c.members[:index], c.members[oneOver:]...)
}

func (c List[T]) Length() int {
	return len(c.members)
}

func (c List[T]) Index(item T) (int, error) {
	for index, itm := range c.members {
		if itm == item {
			return index, nil
		}
	}

	return -1, errors.New("item not found")
}
