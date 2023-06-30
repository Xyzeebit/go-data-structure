package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercises")
}

// Write a function that returns the sum of all the integer in an array
func ArraySum(arr []int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}

// Write a method which will search a list for some give value
func SearchList(v int, list []int) int {
	for i, a := range list {
		if a == v {
			return i
		}
	}
	return -1
}

// Write a binary search function to search an array of integers, return the index when found otherwise return -1
func BinarySearch(v int, list []int) int {
	var first int = 0
	var last int = len(list) - 1
	var mid int
	for first <= last {
		mid = (first + last) / 2
		if list[mid] == v {
			return mid
		} else {
			if list[mid] < v {
				first = mid + 1
			} else {
				last = mid - 1
			}
		}

	}
	return -1
}

// Write a recursive binary search function
func BinaryRecursiveSearch(v int, list []int) int {
	first := 0
	last := len(list) - 1
	var mid int = (first + last) - 1
	if list[mid] == v {
		return mid
	} else if list[mid] < v {
		return BSH(v, list, ((mid + 1 + last) / 2), mid+1, last)
	} else {
		return BSH(v, list, ((mid - 1 + first) / 2), first, mid-1)
	}
}

func BSH(v int, list []int, mid, start, end int) int {
	if list[mid] == v {
		return mid
	} else if list[mid] < v {
		return BSH(v, list, ((mid + 1 + end) / 2), mid+1, end)
	} else {
		return BSH(v, list, ((mid - 1 + start) / 2), start, mid-1)
	}
}

// Given a list, you need to rotate its elements K number of times
// e.g [1,2,3,4,5,6] k=2 [3,4,5,6,1,2]
func RotateList(k int, list []int) []int {
	size := len(list)
	tmp := list[0:k]
	arr := make([]int, size)
	j := 0
	for i := k; i < size; i++ {
		arr[j] = list[i]
		j++
	}

	for i := 0; i < k; i++ {
		arr[j] = tmp[i]
		j++
	}

	return arr
}

// Write a function that sums the values of a 2 dimentional array
func Sum2DimensionalArray(arr [2][2]int) int {
	var total int
	for _, i := range arr {
		for _, j := range i {
			total += j
		}
	}
	return total
}

// Write a function that finds the second largest element in an array
func SecondLargestInList(list []int) int {
	var secondLargest int
	var largest int = list[0]
	for _, n := range list {
		if n > largest {
			secondLargest = largest
			largest = n
		}
	}
	return secondLargest
}

// Write a function that computes the sum on N e.g Sum(N) -> 1+2+3..+N
func Sum(n int) int {
	if n <= 1 {
		return n
	}
	return n + Sum(n-1)
}

// Write a function to reverse a list without creating a new list
func ReverseList(list []int) []int {
	i := 0
	j := len(list) - 1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}

// given a list of 0s and 1s, sort the list so that all the 0s comes before 1s
func Sort01(list []int) []int {
	i := 0
	j := len(list) - 1
	for i <= j {
		if list[i] == 1 {
			if list[j] == 0 {
				list[i], list[j] = list[j], list[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			if list[j] == 0 {
				i++
			} else {
				i++
				j--
			}
		}

	}
	return list
}

// Given an array of n elements, and a value v, find two elements in the array that sums to v
func FindSumInArray(v int, arr []int) []int {
	r := make([]int, 2)
	i := 0
	j := i + 1
	// 1,5,7,2,4 = 7 - [5,2]
	for r[0]+r[1] != v {
		if arr[j] < v {
			if arr[i]+arr[j] == v {
				r[0] = arr[i]
				r[1] = arr[j]
				break
			}
			if j == len(arr) {
				i++
				j = 0
			}
			if i == len(arr) {
				return []int{-1, -1}
			}
		}
		j++
	}
	return r
}
