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
