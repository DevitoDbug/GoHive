package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) Area() (area float64) {
	return r.length * r.width
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() (area float64) {
	return t.Height * (t.Base / 2) * 0.5
}

type Circle struct {
	radius float64
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.radius * c.radius
}

func TotalArea(shapes []Shape) (totalArea float64) {
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Triangle{
			Height: 12,
			Base:   56,
		},
		Circle{radius: 70},
		Rectangle{
			width:  20,
			length: 60,
		},
		Triangle{
			Height: 80,
			Base:   50,
		},
	}
	fmt.Printf("The area of the shapes you gave is %.4f", TotalArea(shapes))
}
