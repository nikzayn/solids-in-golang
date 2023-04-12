package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	r := Rectangle{width: 4.5, height: 6.6}
	c := Circle{radius: 4.2}

	PrintArea(r) // Prints area of rectangle
	PrintArea(c) // Prints area of circle
}
