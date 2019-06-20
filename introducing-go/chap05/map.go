package main

import "fmt"

func main() {
	D := make(map[string]string)

	D["foo"] = "foo!"
	D["bar"] = "bar!"

	if elem, ok := D["baz"]; !ok {
		fmt.Println("Error")
	} else {
		fmt.Println(elem)
	}
}
