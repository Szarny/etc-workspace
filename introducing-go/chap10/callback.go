package main

import (
	"fmt"
	"time"
)

func waitAndRun(n int, f func(string), s string) {
	select {
	case _ = <-time.After(time.Second * time.Duration(n)):
		f(s)
	}
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	go waitAndRun(3, greet, "tsubasa")

	var input string
	fmt.Scanln(&input)
}
