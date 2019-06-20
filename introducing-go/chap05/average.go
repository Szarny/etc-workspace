package main

import "fmt"

func main() {
	var A [5]float64 = [5]float64{
		13,
		244,
		32,
		45,
		54,
	}

	var total float64 = 0
	for _, a := range A {
		total += a
	}

	fmt.Println(total / float64(len(A)))
}
