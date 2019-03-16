package main

import (
	"fmt"
)

func main() {
	i := 1
	f := float64(i)

	fmt.Printf("Type:%T Value:%v\n", i, i)
	fmt.Printf("Type:%T Value:%v\n", f, f)
}