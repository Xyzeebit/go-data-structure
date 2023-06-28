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
