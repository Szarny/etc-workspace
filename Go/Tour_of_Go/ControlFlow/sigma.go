package main

import (
	"fmt"
)

func sigma(n int) (sum int) {
	sum = 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return
}

func main() {
	fmt.Println(sigma(5))
	fmt.Println(sigma(10))
}