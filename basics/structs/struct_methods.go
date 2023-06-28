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

type MyInt int

func (i MyInt) increment1() {
	i = i + 1
}

func (i *MyInt) increment2() {
	*i = *i + 1
}

func main() {
	r := Rectangle{width: 15, height: 10}
	fmt.Println("Area:", r.Area())
	fmt.Printf("Perimeter:%.f\n", r.Perimeter())

	ptr := &Rectangle{width: 8.87, height: 12.25}
	fmt.Printf("Area: %.2f\nPerimeter :%.2f\n", ptr.Area(), ptr.Perimeter())

	fmt.Println("Changing type internal data")
	var n MyInt = 1
	fmt.Println("Before increment()", n)
	n.increment1()
	fmt.Println("After increment1()", n)
	n.increment2()
	fmt.Println("After increment2()", n)
}
