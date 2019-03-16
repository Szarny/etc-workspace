package main

import (
	"fmt"
)

func sayGoodbye() string {
	return "Good Bye!"
}

func main() {
	defer fmt.Println(sayGoodbye())

	fmt.Println("Hello!")
}
