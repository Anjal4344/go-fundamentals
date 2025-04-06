package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) details() {
	fmt.Printf("%s's age is %d\n", p.name, p.age)
}
func main() {
	personal := person{name: "AJU", age: 24}
	personal.details()
}
