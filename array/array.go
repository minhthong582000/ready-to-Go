package main

import (
	"errors"
	"fmt"
)

// Array - dynamic Array structure
type Array struct {
	arr              []interface{}
	capacity, length int
}

// Constructor
func newArrayDefault() (*Array, error) {
	arr, err := newArray(16)

	return arr, err
}

// Constructor
func newArray(capacity int) (*Array, error) {
	if capacity < 0 {
		return nil, errors.New("illegal capacity")
	}

	arr := Array{make([]interface{}, capacity), capacity, 0}

	return &arr, nil
}

// Getter
func (arr Array) get(index int) interface{} {
	if index < 0 || index > arr.length-1 {
		// todo throw an error
		return nil
	}
	return arr.arr[index]
}

// Setter
func (arr Array) set(index int, newValue interface{}) {
	if index < 0 || index > arr.length-1 {
		// todo throw an error
		return
	}
	arr.arr[index] = newValue
}

func (arr Array) size() int {
	return arr.length
}

func (arr Array) isEmpty() bool {
	return arr.size() == 0
}

func (arr Array) clear() {
	for i := range arr.arr {
		arr.arr[i] = nil
	}

	arr.length = 0
}

func (arr *Array) add(elem interface{}) {
	if arr.length+1 > arr.capacity {
		if arr.capacity == 0 {
			arr.capacity = 1
		} else {
			arr.capacity *= 2
		}

		newArr := make([]interface{}, arr.capacity)
		for i := 0; i < arr.length; i++ {
			newArr[i] = arr.arr[i]
		}

		arr.arr = newArr
	}

	arr.arr[arr.length] = elem
	arr.length++
}

func (arr *Array) removeAt(index int) {
	if index < 0 || index > arr.length-1 {
		// todo throw an error
		return
	}

	arr.length--
	newArr := make([]interface{}, arr.capacity)

	for i := 0; i < arr.length; i++ {
		if i >= index {
			nextIndex := i + 1
			newArr[i] = arr.arr[nextIndex]
		} else {
			newArr[i] = arr.arr[i]
		}
	}

	arr.arr = newArr
}

func (arr *Array) remove(obj interface{}) {
	target := arr.indexOf(obj)

	if target < 0 {
		// todo throw an error
		return
	}

	arr.removeAt(target)
}

func (arr Array) indexOf(obj interface{}) int {
	for i := range arr.arr {
		if arr.arr[i] == obj {
			return i
		}
	}

	return -1
}

func (arr Array) contain(obj interface{}) bool {
	return arr.indexOf(obj) >= 0
}

func (arr Array) toString() string {
	var str string

	if arr.length == 0 {
		return "[]"
	}

	str += "["
	for i := 0; i < arr.length; i++ {
		s := fmt.Sprintf(" %v ", arr.arr[i])
		str += s
	}
	str += "]"

	return str
}

func main() {
	arr, _ := newArray(0)

	arr.add("nigga")

	fmt.Println(arr.indexOf("xx"))

	arr.add(2)
	fmt.Println(arr.contain("nigga"))

	arr.add(3)

	arr.remove("nigga")
	arr.removeAt(1)

	fmt.Println(arr.toString())

	arr.add(3)
	arr.set(-1, "dude")

	fmt.Println(arr.toString())
	fmt.Println(arr.get(-1))

	arr.removeAt(-1)
	fmt.Println(arr.toString())
}
