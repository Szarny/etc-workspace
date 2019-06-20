package main

import (
	"fmt"
)

func sum(A ...int) int {
	total := 0

	for _, v := range A {
		total += v
	}

	return total
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(sum(A...))
}
