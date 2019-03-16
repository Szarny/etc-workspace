package main

import (
	"fmt"
)

func swap(S string, T string) (string, string) {
	return T, S
}

func main() {
	fmt.Println(swap("hoge", "fuga"))
}