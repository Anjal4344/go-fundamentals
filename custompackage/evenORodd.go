package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter a number")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input. PLEASE ENTER AN INTEGER")
		return
	}

	if a%2 == 0 {
		fmt.Println("The given number is EVEN")
	} else {
		fmt.Println("The given number is odd")
	}
}
