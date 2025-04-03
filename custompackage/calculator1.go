package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter OPERATOR(+,-,*,/): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	switch operator {
	case "+":
		fmt.Println("The result is: ", num1+num2)
	case "-":
		fmt.Println("The result is: ", num1-num2)
	case "*":
		fmt.Println("The result is: ", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result is: ", num1/num2)
		} else {
			fmt.Println("ERROR,Division by zero is not allowed")
		}
	default:
		fmt.Println("Invalid operator")
	}

}
