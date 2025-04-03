package main

import "fmt"

func findduplicate(num []int) []int {
	result := []int{}
	for i := 0; i < len(num); i++ {
		duplicate := false
		for j := 0; j < len(result); j++ {
			if num[i] == result[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, num[i])
		}
	}
	return result
}
func main() {
	numbers := []int{12, 13, 11, 10, 1, 0, 10, 10, 25, 25, 2, 2, 11}
	exactnumbers := findduplicate(numbers)
	fmt.Println("EXACT NUMBERS ARE : ", exactnumbers)
}
