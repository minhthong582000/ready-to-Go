package main

import (
	"fmt"
)

func main() {
	// If condition
	var num int8 = 10

	if num == 10 {
		fmt.Println(true)
	}

	if num < 69 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// If statement; condition { // Do someshit}
	if a := 100; a > 100 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// Switch case
	fmt.Println("-----------")
	switch num {
	case 1, 3, 19:
		fmt.Println(1)
		// auto break
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Ah shiet")
	}

	fmt.Println("-----------")
	switch {
	case num > 9:
		fmt.Println(1)
	case num <= 9:
		fmt.Println(2)
	default:
		fmt.Println("Ah shiet")
	}

	fmt.Println("-----------")
	num = 2
	switch num { // FallThrough
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	case 10:
		goto handleNumEqual10
	handleNumEqual10:
		fmt.Println("Nice")
	case 3:
		fmt.Println(3)
		fallthrough
	case 4:
		fmt.Println(4)
		fallthrough
	case 5:
		fmt.Println(5)
		fallthrough
	case 6:
		fmt.Println(6)
		fallthrough
	default:
		fmt.Println("Default")
	}

	fmt.Println("-----------")

	// For loop
	// While loop
	var j int8 = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("-----------")

	// Nested loop, multiple condition
	for i, j := 0, 0; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i + j)
	}
}
