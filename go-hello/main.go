package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	// Type Inference
	var email = "abc@xyz"
	var z = 10
	fmt.Println(email)
	fmt.Println(z)

	// Declare many variables
	// 1. Same type
	// 2. Difference type
	// ------------------------
	// 1
	var a, b int
	a = 1
	b = 2
	fmt.Println(a + b)

	var x, y int = 10, 11
	fmt.Println(x + y)

	//2
	var (
		name    string = "Thongdz"
		address string = "Cantho"
		age     int    = 69
	)

	fmt.Println(name, address, age)

	// Short variable declarations
	damn := "Bro"
	fmt.Println(damn)
}
