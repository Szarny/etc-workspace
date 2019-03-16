package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
