package main

import "fmt"

func arithmatics(x, y float64) (float64, float64, float64, float64) {
	if y == 0 {
		fmt.Println("Error : Division by zero")
		return x + y, x - y, x * y, 0
	}
	return x + y, x - y, x * y, x / y
}
func main() {
	var a, b float64
	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input. PLEASE ENTER AN INTEGER")
		return
	}
	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid input. PLEASE ENTER AN INTEGER")
		return
	}
	sum, diff, prod, quot := arithmatics(a, b)
	fmt.Println("Sum = ", sum, "\nDifference = ", diff, "\nProduct = ", prod)
	if b != 0 {
		fmt.Printf("Quotient = %.2f\n", quot)
	}
}
