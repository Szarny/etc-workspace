package main

import "fmt"

type S struct {
	x int
	y float64
	z string
	w bool
}

func main() {
	s1 := new(S)
	fmt.Println(s1)

	s2 := S{x: 1, y: 2.2, z: "foo", w: true}
	fmt.Println(s2)
}
