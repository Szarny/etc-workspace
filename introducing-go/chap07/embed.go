package main

import "fmt"

type Horn struct {
	sound string
}

func (horn *Horn) push() {
	fmt.Println(horn.sound)
}

type Car struct {
	Horn
}

func (car *Car) drive() {
	fmt.Println("brrrr...")
}

func main() {
	car := Car{Horn{sound: "Paaaa!!!"}}
	car.drive()
	car.Horn.push()
}
