package main

import "fmt"

func maxmin(arr []int) (int, int) {
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
	a := make([]int, 5)
	fmt.Print("Enter 5 numbers of elements of array: ")
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&a[i])
		if err != nil {
			fmt.Println("type integer", err)
			return
		}
	}
	min1, max1 := maxmin(a)
	fmt.Printf("max = %d\n", max1)
	fmt.Printf("min = %d\n", min1)
}
