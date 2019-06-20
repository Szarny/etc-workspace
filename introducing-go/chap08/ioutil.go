package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buffer, err := ioutil.ReadFile("test.txt")

	if err != nil {
		panic("ReadFile Error")
	}

	fmt.Println(string(buffer))
}
