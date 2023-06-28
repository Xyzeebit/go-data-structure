package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rectangle{width: 15, height: 10}
	fmt.Println("Area:", r.Area())
	fmt.Printf("Perimeter:%.f\n", r.Perimeter())

	ptr := &Rectangle{width: 8.87, height: 12.25}
	fmt.Printf("Area: %.2f\nPerimeter :%.2f\n", ptr.Area(), ptr.Perimeter())
}
