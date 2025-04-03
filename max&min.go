package main

import "fmt"

func maxmin(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("ARRAY CAN'T BE EMPTY")
	}
	min0, max0 := arr[0], arr[0]
	for _, num := range arr {
		if num < min0 {
			min0 = num
		}
		if num > max0 {
			max0 = num
		}
	}
	return min0, max0
}
func main() {
	var ARRAY []int
	fmt.Print("Enter number of elements of array =")
	_, err := fmt.Scan(&ARRAY)
	if err != nil {
		fmt.Println("Error reading arrray =", ARRAY)
		return
	}
	min1, max1 := maxmin(ARRAY)
	fmt.Printf("max = %.2d\n", max1)
	fmt.Printf("min = %.2d\n", min1)
}
