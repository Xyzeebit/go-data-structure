package main

import "fmt"

func main() {
	arr := []int{2, 8, 1, 9, 10, 1, 5, 7}
	v := 5
	find := linearSearch(arr, v)
	fmt.Printf("search in %v for %d : %d\n", arr, v, find)

	arr2 := []int{1, 2, 4, 6, 7}
	v = 2
	find = linearSearchSorted(arr2, v)
	fmt.Printf("search in %v for %d : %d\n", arr2, v, find)
}

func linearSearch(list []int, v int) int {
	for i, a := range list {
		if v == a {
			return i
		}
	}
	return -1
}

func linearSearchSorted(list []int, v int) int {
	for i, a := range list {
		if v < a {
			return -1
		}
		if v == a {
			return i
		}
	}
	return -1
}
