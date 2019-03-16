package main

import (
	"fmt"
)

var (
	maxint uint64 = 1<<64 - 1
)

func main() {
	fmt.Printf("Type:%T Value:%d\n", maxint, maxint)
}