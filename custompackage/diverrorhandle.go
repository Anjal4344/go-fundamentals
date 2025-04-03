package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	var x, y float64

	// Taking user input
	fmt.Print("Enter numerator: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Invalid input. PLEASE ENTER AN INTEGER")
		return
	}

	fmt.Print("Enter denominator: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		fmt.Println("Invalid input. PLEASE ENTER AN INTEGER")
		return
	}
	result, err := div(x, y)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
