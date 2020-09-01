package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int16
	address string
}

func newPerson(p *Person) *Person {
	p.name = "Oke"
	p.age = 33
	p.address = "33 CMT8"

	return p
}

func main() {
	var p Person
	var p1 Person = Person{name: "x", age: 69}
	p2 := new(Person)
	var x int16 = p2.age
	fmt.Println(*newPerson(&p))
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(x)
}
