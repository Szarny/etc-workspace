package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {
	if value := math.Pow(x, n); value < limit {
		return value
	}

	return limit
}

func main() {
	fmt.Println(fmt.Sprint(pow(3,3, 100)))
	fmt.Println(fmt.Sprint(pow(3,3, 10)))
}