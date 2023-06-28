package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct{ radius float64 }

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea(areas ...Shape) float64 {
	var total float64
	for _, s := range areas {
		total += s.Area()
	}
	return total
}

func TotalPerimeter(peri ...Shape) float64 {
	var total float64
	for _, s := range peri {
		total += s.Perimeter()
	}
	return total
}

func main() {
	r := Rectangle{width: 9, height: 5}
	c := Circle{radius: 12}
	fmt.Println("Total area:", TotalArea(r, c))
	fmt.Println("Total perimeter:", TotalPerimeter(r, c))
}
