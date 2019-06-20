package main

import (
	"container/list"
	"fmt"
)

func main() {
	var L list.List
	L.PushBack(1)
	L.PushBack(2)
	L.PushBack(3)

	for l := L.Front(); l != nil; l = l.Next() {
		fmt.Println(l.Value.(int) + 1)
	}
}
