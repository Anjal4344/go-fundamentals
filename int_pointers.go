package main

import "fmt"

func modify(x *int) {
	*x = *x + 20
}
func main() {
	num := 42
	fmt.Println("Before:", num)
	modify(&num)
	fmt.Println("After:", num)
}
