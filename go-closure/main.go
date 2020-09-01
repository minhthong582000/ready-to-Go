package main

import "fmt"

func main() {
	//sum := func(a, b int) int { return a + b }(3, 4)
	//p := (*int)(nil)
	var f func(func(int, int) int, int) func(int, int) int
	fmt.Printf("%v\n", f)
}
