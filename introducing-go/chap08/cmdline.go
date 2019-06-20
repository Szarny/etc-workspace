package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value")
	flag.Parse()
	args := flag.Args()
	fmt.Println(rand.Intn(*maxp))
	fmt.Println(args)
}
