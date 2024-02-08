package main

import (
	"fmt"

	"github.com/isaacp/collections/queue"
	//"github.com/isaacp/collections/sort"
)

func main() {
	//a := make([]rune, 0)
	//a = append(a, 'e')
	// fmt.Println(rune('0'))
	// test := []int{2, 1, 5, 3, 7, 54, 9, 4, 2, 5, 67, 8, 3, 2, 34, 6, 7, 90, 3, 4, 5, 33, 2}
	// test2 := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70}
	// fmt.Println(sort.BubbleSort(test))
	// fmt.Println(sort.BubbleSort(test2))

	// arry := []int{1, 200, 34, 56, 23, 45, 87, 6, 43, 223, 78}
	// arry2 := []int{1}
	// sort.QuickSort(arry2)
	// fmt.Println(arry2)
	// sort.QuickSort(arry)
	// fmt.Println(arry)

	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(43)
  q.Print()
	fmt.Println(q.Dequeue().Value)
	fmt.Println(q.Dequeue().Value)
	fmt.Println(q.Dequeue().Value)
	fmt.Println(q.Dequeue().Value)
	//fmt.Println(q.Dequeue().Value)

	q.Print()
}
