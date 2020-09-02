package main

import "fmt"

func binarySearch(arr []int, l, h, x int) bool {
	pivot := (l + h) / 2

	if l <= h {
		if x == arr[pivot] {
			return true
		}
		if x < arr[pivot] {
			return binarySearch(arr, l, pivot-1, x)
		}
		if x > arr[pivot] {
			return binarySearch(arr, pivot+1, h, x)
		}
	}

	return false
}

func searchForFirstAndLast(arr []int, l, h, x int) int {
	var pos int = h + 1

	for l <= h {
		pivot := (l + h) / 2

		if arr[pivot] >= x {
			pos = pivot
			h = pivot - 1
		} else {
			l = pivot + 1
		}
	}

	return pos
}

func main() {
	var x int
	fmt.Scan(&x)

	arr := []int{1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 6, 7, 9, 10}
	found := binarySearch(arr, 0, len(arr)-1, x)

	first := searchForFirstAndLast(arr, 0, len(arr)-1, x)
	last := searchForFirstAndLast(arr, 0, len(arr)-1, x+1) - 1
	if first <= last {
		fmt.Printf("%d is between %d and %d\n", x, first, last)
	}
	fmt.Println(found)
}
