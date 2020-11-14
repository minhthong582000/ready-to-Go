package main

import "fmt"

func computeD(phi, e int) int {
	var x, y int
	for i := 0; i < 1000; i++ {
		x = (i*phi + 1) / e
		y = (x * e) % phi
		if y == 1 {
			return x
		}
	}

	return -1
}

func main() {
	fmt.Println(computeD(880, 13))
}
