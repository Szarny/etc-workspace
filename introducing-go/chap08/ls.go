package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")

	if err != nil {
		panic("os.Open Error")
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		panic("dir.Readdir Error")
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Size())
	}
}
