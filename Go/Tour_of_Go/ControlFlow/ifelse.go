package main

import (
	"fmt"
)

func double(n int) string {
	if n2 := n * 2; n2  > 10 {
		return "2倍したら10を超える"
	} else {
		return "2倍しても10を超えない"
	}
}

func main() {
	fmt.Println(double(3))
	fmt.Println(double(44))
}