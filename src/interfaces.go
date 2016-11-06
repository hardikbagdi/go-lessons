package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	length, width float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func main() {

	shape := Rectangle{2, 3}
	fmt.Println(shape.area())
}
