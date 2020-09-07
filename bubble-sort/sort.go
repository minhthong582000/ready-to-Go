package main

import (
	"fmt"
)

// InputSlice Initialize slice of integer
func InputSlice() []int {
	var sl []int

	for i := 0; i < 10; i++ {
		var input int
		fmt.Scan(&input)
		sl = append(sl, input)
	}

	return sl
}

// Swap Swap two integers pass by reference
func Swap(first *int, second *int) {
	*first, *second = *second, *first
}

// BubbleSort Sort array of integers using bubble sort algorithm
func BubbleSort(arr []int) {
	var arrLen int = len(arr)

	for i := range arr {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func main() {
	arr := InputSlice()
	BubbleSort(arr)

	fmt.Println(arr)
}
