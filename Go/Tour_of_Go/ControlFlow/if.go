package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) string {
	if n < 0 {
		return sqrt(-n) + "i"
	}
	
	return fmt.Sprint(math.Sqrt(n))
}

func main() {
	fmt.Println(sqrt(5), sqrt(-4))
}