package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", ":9999")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Received :", msg)
	}

	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}

	msg := "Hello, World!"
	fmt.Println("Sending :", msg)

	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		panic(err)
	}

	conn.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
