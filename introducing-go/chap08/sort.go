package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type Persons []Person

func (persons Persons) Len() int {
	return len(persons)
}

func (persons Persons) Less(i, j int) bool {
	return persons[i].age < persons[j].age
}

func (persons Persons) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func main() {
	persons := Persons{
		{"piyo", 30},
		{"hoge", 10},
		{"fuga", 20},
	}

	sort.Sort(sort.Reverse(persons))
	fmt.Println(persons)
}
