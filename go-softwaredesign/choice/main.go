package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	argc := len(os.Args) - 1
	if argc < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s choice1 choice2 ...", os.Args[0])
		os.Exit(1)
	}

	fmt.Println(os.Args[rand.Intn(argc)+1])
}
