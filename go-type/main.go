package main

import (
	"fmt"
	"math"
)

func main() {
	// Bool
	var myBool bool = true
	fmt.Println("Bool: ", myBool)

	// String
	var myString string = "Hello"
	fmt.Println("String: ", myString)

	// Int, Int8, Int16, Int32, Int64
	// 1. Range
	// 2. Bits
	var myInt int = 10
	fmt.Println("Int: ", myInt)
	//Int8, 16, 32, 64
	fmt.Println("Int8 range: ", math.MinInt8, "-->", math.MaxInt8)
	fmt.Println("Int16 range: ", math.MinInt16, "-->", math.MaxInt16)
	fmt.Println("Int32 range: ", math.MinInt32, "-->", math.MaxInt32)
	fmt.Println("Int64 range: ", math.MinInt64, "-->", math.MaxInt64)

	// Unsigned Int (Khong Dau)
	// var myUInt uint = -19
	var myUInt uint = 19
	fmt.Println("Unsigned Int: ", myUInt)

	// Byte (Alias for uint8)
	// var myByte byte = 259
	var myByte byte = 255
	fmt.Println("Byte: ", myByte)
	fmt.Printf("Byte alias: %T\n", myByte) // uint8

	var a byte = 'A'
	fmt.Println("Dec of \"A\": ", a)    // Dec 65, A in ASCII
	fmt.Printf("Hex of \"A\": %x\n", a) // Hex 41

	// FLoat
	var (
		myFloat       float32 = 69.96
		mySecondFloat float64 = -33.332302
	)
	fmt.Println("Float: ", myFloat, ",", mySecondFloat)

	// Complex
	// x = a + bi
	var myComplex complex64 = 6 + 9i
	var mySecondComplex complex64 = -445 - 943i
	fmt.Println("Complex: ", myComplex+mySecondComplex)

	// Rune
	// Does not matter how many bytes the code point occupies, it can be represented by a rune
	// Alias of Int32
	// Rune Represent a Unicode code point in Go

	// Ex: he take 2 bit
	// 	   but hê take 3 bit

	var mySecondString = "Kiểu dữ liệu & Hằng số "
	fmt.Println(mySecondString)

	myRune := []rune(mySecondString)

	for i := 0; i < len(myRune); i++ {
		fmt.Printf("%c", myRune[i])
	}

	var mySecondRune rune = 'A'
	fmt.Printf("\nRune alias: %T", mySecondRune)
	fmt.Printf("\nRune unicode: %U\n", mySecondRune)

	// Type Conversions
	var first int = 1
	var second float64 = 6.9
	fmt.Println(float64(first) + second)

	// Constants
	// Can NOT be declared using the ":=" syntax
	// const PI := 3.14159
	const PI float32 = 3.14159
	fmt.Println(PI)

	// Unsigned Int Pointer
	// Store var address
}
