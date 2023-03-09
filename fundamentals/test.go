package main

import "fmt"

// STRUCTS
type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

func main() {

	rect1 := Rectangle{height: 15, width: 10}

	fmt.Println("Rectangle is", rect1.width, "wide")

	fmt.Println("Area of the rectangle =", rect1.area())

}
