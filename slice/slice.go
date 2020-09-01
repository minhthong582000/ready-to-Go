package main

import (
	"fmt"
	"strconv"
)

func partition(a []int, low, high int) int {
	p := a[high]
	for j := low; j < high; j++ {
		if a[j] < p {
			a[j], a[low] = a[low], a[j]
			low++
		}
	}
	a[low], a[high] = a[high], a[low]

	return low
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	index := partition(arr, low, high)

	quickSort(arr, low, index-1)
	quickSort(arr, index+1, high)
}

func slice() {
	var sl []int
	var x string

	for true {
		fmt.Print("Input: ")
		fmt.Scan(&x)

		if x == "x" {
			fmt.Println("Exited!")
			break
		}
		if i, err := strconv.Atoi(x); err == nil {
			sl = append(sl, i)
		} else {
			fmt.Println("You must enter an integer!")
			continue
		}

		quickSort(sl, 0, len(sl)-1)
		fmt.Println(sl)
	}
}

func main() {
	slice()
}
