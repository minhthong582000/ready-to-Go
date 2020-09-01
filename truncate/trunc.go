package main

import "fmt"

func main() {
	// var x float32

	// Hash table
	var idMap map[string]int
	idMap = make(map[string]int)

	// Map literal
	nameMap := map[string]int{
		"James": 6,
	}

	idMap["James"] = 69
	delete(idMap, "James")
	nameMap["John Cena"] = 96
	id, p := nameMap["John Cena"]
	fmt.Printf("%d %d %d %t\n", idMap["James"], nameMap["John Cena"], id, p)
	for id, v := range nameMap {
		fmt.Printf("%s %d\n", id, v)
	}

	// Array
	x := [...]int{1, 2, 3, 4}

	// Slide
	sl := x[1:2] // {2}
	slMake := make([]int, 5, 9)

	// Append to the end of the slide
	sl = append(sl, 66)

	fmt.Println(len(sl))
	fmt.Println(cap(slMake))

	for i, v := range x {
		fmt.Println(i + sl[1])
		fmt.Println(v)
	}

}
