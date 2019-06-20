package main

import "fmt"

type Square struct {
	width  int
	height int
}

func calcArea(square Square) int {
	return square.width * square.height
}

func (square Square) getArea() int {
	return square.width * square.height
}

func main() {
	square := Square{4, 5}
	fmt.Println(calcArea(square))
	fmt.Println(square.getArea())
}
