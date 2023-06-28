package main

import "fmt"

func main() {
	s := "Hello world"
	r := []rune(s)
	r[0] = 'Y'
	x := string(r)
	fmt.Println(x)
}
