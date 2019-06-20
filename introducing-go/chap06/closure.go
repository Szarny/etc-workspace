package main

import "fmt"

func genFib() func() int {
	x := 1
	y := 1

	return func() int {
		z := x + y
		x = y
		y = z
		return z
	}
}

func main() {
	getFib := genFib()
	for i := 0; i < 10; i++ {
		fmt.Println(getFib())
	}
}
