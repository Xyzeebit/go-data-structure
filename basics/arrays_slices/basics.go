package main

import "fmt"

func arrays() {
	var arr [10]int
	fmt.Println(arr)

	for i := 0; i < 10; i++ {
		arr[i] = i * i
	}
	fmt.Println(arr)
}

func main() {
	arrays()
}
