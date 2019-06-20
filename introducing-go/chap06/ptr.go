package main

import "fmt"

func override(xPtr *int) {
	*xPtr = 123
}

func swap(p, q *int) {
	tmp := *p
	*p = *q
	*q = tmp
}

func main() {
	x := 1
	fmt.Println("x =", x)
	override(&x)
	fmt.Println("x =", x)

	p := new(int)
	*p = 12345
	fmt.Println(p, *p)

	a := 123
	b := 456
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
