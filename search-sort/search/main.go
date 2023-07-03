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

	v = 7
	n := binarySearch(arr2, v)
	fmt.Printf("search in %v for %d : %d\n", arr2, v, n)

	arr3 := []int{1, 2, 4, 6, 2, 7, 4, 6, 5, 9, 1}
	fmt.Printf("Array: %v\n", arr3)
	printDuplicateBrute(arr3)
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

func binarySearch(list []int, v int) int {
	start := 0
	end := len(list) - 1
	var mid int

	for start <= end {
		mid = (start + end) / 2
		if list[mid] == v {
			return mid
		} else if list[mid] < v {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// Given an array of numbers, print the duplicate elements in the array
func printDuplicateBrute(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Print(arr[j], " ")
				break
			}
		}
	}
}
