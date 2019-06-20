package main

import "fmt"

func average(A []float64) float64 {
	total := 0.0

	for _, v := range A {
		total += v
	}

	return total / float64(len(A))
}

func main() {
	A := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(A))
}
