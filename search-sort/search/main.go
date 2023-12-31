package main

import (
	"fmt"
	"sort"
)

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
	fmt.Println()

	fmt.Println("Find duplicate sorted method")
	printDuplicateSort(arr3)
	fmt.Print("\nFind duplicate hash approach\n")
	printDuplicateHash(arr3)

	fmt.Println("Print duplicate count method")
	printDuplicateCount(arr3, 11)
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

func printDuplicateSort(arr []int) {
	sort.Ints(arr)
	fmt.Printf("%v\n", arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			fmt.Print(arr[i], " ")
		}
	}
}

func printDuplicateHash(arr []int) {
	m := make(map[int]int)
	for _, x := range arr {
		n, ok := m[x]
		if ok {
			fmt.Print(n, " ")

		} else {
			m[x] = x
		}
	}
	fmt.Println()
}

func printDuplicateCount(arr []int, max int) {
	count := make([]int, max)
	for i := 0; i < len(arr); i++ {
		count[i] = 0
	}

	for i, a := range arr {
		if count[arr[i]] == 1 {
			fmt.Print(a, ", ")
		} else {
			count[arr[i]]++
		}
	}
	fmt.Println()
}
