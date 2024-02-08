package sort

import "fmt"

func BubbleSort(arr []int) []int {
	swapped := false

	for {
		for i, j := 0, 1; i < len(arr)-1; i, j = i+1, j+1 {
			if arr[i] > arr[j] {
				swap(arr, i, j)
				swapped = true
			}
		}
		if swapped {
			swapped = false
		} else {
			return arr
		}
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	fmt.Println("Entering Quicksort")
	if left >= right {
		return
	}

	pivotIndex := right
	pivot := arr[pivotIndex]
	fmt.Println("pivot " + fmt.Sprintf("%d", pivot))

	for left < right {
		if arr[left] > pivot {
			fmt.Printf("%d > %d\n", arr[left], pivot)
			for left < right {
				if arr[right] < pivot {
					fmt.Printf("%d - %d\n", arr[left], arr[right])
					swap(arr, left, right)
					left++
					right--
					fmt.Printf("indices: %d - %d\n", left, right)
					break
				} else {
					right--
				}
			}
		} else {
			fmt.Printf("%d < %d\n", arr[left], pivot)
			left++
		}
	}
	quickSort(arr, left, pivotIndex/2)
	quickSort(arr, (pivotIndex/2)+1, right)
}

func swap(arr []int, i int, j int) {
	fmt.Print("swapping: ")
	fmt.Println(arr)
	arr[i], arr[j] = arr[j], arr[i]
	fmt.Print("swapped: ")
	fmt.Println(arr)
}
