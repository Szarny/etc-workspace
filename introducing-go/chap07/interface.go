package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Square struct {
	width  float64
	height float64
}

func (s *Square) area() float64 {
	return s.width * s.height
}

func calcTotalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}

func getCircle(radius float64) Shape {
	return &Circle{radius}
}

func getSquare(width, height float64) Shape {
	return &Square{width, height}
}

func main() {
	c := getCircle(10.0)
	s := getSquare(3, 4)
	fmt.Println(calcTotalArea(c, s, Shape(&Square{1, 1})))
}
