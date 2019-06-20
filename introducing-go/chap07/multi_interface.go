package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	total := 0.0

	for _, shape := range m.shapes {
		total += shape.area()
	}

	return total
}

func (m *MultiShape) perimeter() float64 {
	total := 0.0

	for _, shape := range m.shapes {
		total += shape.perimeter()
	}

	return total
}

type Circle struct {
	radius float64
}

type Square struct {
	width  float64
	height float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (s *Square) area() float64 {
	return s.width * s.height
}

func (s *Square) perimeter() float64 {
	return s.width*2 + s.height*2
}

func main() {
	multiShape := &MultiShape{
		[]Shape{
			&Circle{1},
			&Square{1, 2},
		},
	}

	fmt.Println(multiShape.area())
	fmt.Println(multiShape.perimeter())
}
