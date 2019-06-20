package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Hello!"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "Hello!"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("[Chan 1]", msg1)
			case msg2 := <-c2:
				fmt.Println("[Chan 2]", msg2)
			case aft := <-time.After(time.Second * 3):
				fmt.Println(aft)
				// default:
				// 	fmt.Println("Nothing")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
