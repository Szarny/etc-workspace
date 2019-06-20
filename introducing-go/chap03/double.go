package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Print("Enter a number >> ")
	fmt.Scanf("%d", &input)

	output := input * 2

	fmt.Println(output)
}
