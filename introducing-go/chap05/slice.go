package main

import "fmt"

func main() {
	A := make([]int, 3, 9)
	B := A[:6]

	fmt.Println(A, len(A), cap(A))
	fmt.Println(B, len(B), cap(B))
}
