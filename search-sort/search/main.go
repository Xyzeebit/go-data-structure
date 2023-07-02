package main

import "fmt"

func main() {
	arr := []int{2, 8, 1, 9, 10, 1, 5, 7}
	v := 5
	find := linearSearch(arr, v)
	fmt.Printf("search in %v, %d : ", arr, v)
	fmt.Println(find)
}

func linearSearch(list []int, v int) int {
	for i, a := range list {
		if v == a {
			return i
		}
	}
	return -1
}
