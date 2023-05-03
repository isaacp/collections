package main

import (
	"fmt"

	"../code/go/collections/set"
	"github.com/isaacp/collections/set"
)

type (
	Burger struct {
		Name  string
		price int
	}
)

func main() {

	b1 := Burger{Name: "jaeger"}
	b2 := Burger{Name: "bologna"}

	s1 := set.NewSet()
	s2 := set.NewSet()
	s1.Add(b1)
	s1.Add(b2)

	s2.Add(b1)

	diff := set.Difference(s1, s2)
	fmt.Println(diff)
}
