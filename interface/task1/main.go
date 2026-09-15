package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	circle := NewCircle(5)
	rectangle := NewRectangle(5, 10)

	fmt.Println("Circle S: ", circle.Area())
	fmt.Println("Circle P: ", circle.Perimeter())

	fmt.Println("Rectangle S: ", rectangle.Area())
	fmt.Println("Rectangle P: ", rectangle.Perimeter())
}
