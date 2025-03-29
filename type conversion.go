package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	a := 12896
	b := float64(a)
	fmt.Println("integer is:", a)
	fmt.Println("converted to float64 is:", b)
	//
	c := 52.25
	d := int(c)
	fmt.Println("float value is:", c)
	fmt.Println("integer value is:", d)
	//
	e := 125968
	f := strconv.Itoa(e)
	fmt.Println("integer is:", e)
	fmt.Println("string value is:", f)
	//
	g := "152625"
	h, err := strconv.Atoi(g)
	if err != nil {
		fmt.Println("Error In Converting To Integer", err)
	} else {
		fmt.Println("String is:", g)
		fmt.Println("integer is:", h)
	}
	//
	i := "2.54"
	j, err := strconv.ParseFloat(i, 64)
	if err != nil {
		fmt.Println("Error In Converting To Float", err)
	} else {
		fmt.Println("String is:", i)
		fmt.Println("Float is:", j)
	}
	//
	k := []byte{72, 101, 108, 108, 111, 32, 65, 78, 74, 65, 76}
	l := string(k)
	fmt.Println("Byte is:", k)
	fmt.Println("String is:", l)
	//
	m := "Hello ANJAL"
	n := []byte(m)
	fmt.Println("String is:", m)
	fmt.Println("Byte is:", n)
	//
	time := 10
	if time < 10 {
		fmt.Println("good day")
	} else if time > 10 {
		fmt.Println("bad day")
	} else {
		fmt.Println("average day")
	}
	count := 0
	for {
		fmt.Println("Infinite Loop:", count)
		count++
		if count >= 5 {
			break
		}
	}
	//
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	//
	counter := 0
	for {
		fmt.Println("counter :", counter)
		counter++
		if counter >= 11 {
			break
		}
	}
	nums := []int{10, 20, 30, 40}

	for _, value := range nums {
		fmt.Println(value)
	}
	//
	capitals := map[int]string{0: "Delhi", 1: "Washington", 2: "Paris"}

	for k, v := range capitals {
		fmt.Println(k, v)
	}
	//
	str := "Golang"

	for i, c := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}
	//
	angel := "Golang"
	for i, c := range angel {
		fmt.Println(i, c)
	}
	//
	err = errors.New("Golang Error")
	fmt.Println(err)
}
