package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hello", "e"))
	fmt.Println(strings.Contains("hello", "v"))

	fmt.Println(strings.Count("hello", "l"))

	fmt.Println(
		strings.Join([]string{"hoge", "fuga"}, "/"))
}
