package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	crc := crc32.NewIEEE()
	crc.Write([]byte("aaa"))
	value1 := crc.Sum32()
	fmt.Println(value1)

	hash := sha1.New()
	hash.Write([]byte("aaa"))
	value2 := hash.Sum([]byte(""))
	fmt.Println(value2)
}
