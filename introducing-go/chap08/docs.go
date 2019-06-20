package main

import "fmt"

// Calculate the sum of a slice of int
func Sum(L []int) int {
	t := 0
	for _, v := range L {
		t += v
	}
	return t
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
}
