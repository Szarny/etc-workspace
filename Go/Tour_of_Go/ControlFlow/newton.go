package main

import (
	"fmt"
	"math"
)

func mySqrt(n float64) float64 {
	lower := 1.0
	upper := n
	var mid float64

	for i := 0; i < 100; i++ {
		mid = (lower + upper) / 2

		if math.Pow(mid, 2) > n {
			upper = mid
		} else {
			lower = mid
		}
	}

	return mid
}

func main() {
	fmt.Println(fmt.Sprint(mySqrt(10)))
	fmt.Println(fmt.Sprint(math.Sqrt(10)))
}