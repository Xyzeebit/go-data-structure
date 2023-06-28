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

func slices() {
	var sli []int
	fmt.Println(sli)

	for i := 0; i < 17; i++ {
		sli = append(sli, i)
		fmt.Printf("%v -> len:%d -> cap: %d\n", sli, len(sli), cap(sli))
	}
}

func main() {
	arrays()
	slices()
}
