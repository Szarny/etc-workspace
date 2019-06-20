package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Open("test.txt")

	if err != nil {
		panic("File open error")
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()

	if err != nil {
		panic("Fileinfo error")
	}

	buffer := make([]byte, fileInfo.Size())

	_, err = fp.Read(buffer)

	if err != nil {
		panic("File read error")
	}

	str := string(buffer)
	fmt.Println(str)
}
