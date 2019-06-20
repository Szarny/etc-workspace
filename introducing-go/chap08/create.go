package main

import "os"

func main() {
	fp, err := os.Create("create.txt")

	if err != nil {
		panic("os.Create error")
	}
	defer fp.Close()

	fp.WriteString("Hello!")
}
