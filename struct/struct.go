package main

import (
	"bufio"
	"fmt"
	"os"
)

type shape interface{}

func Draw(s shape) {
	switch sh := s.(type) {
	case Animal:
		fmt.Println(sh)
	}
}

// Animal class
type Animal struct {
	food, locomotion, noise string
}

// Eat method print the animal’s food
func (a Animal) Eat() {
	fmt.Printf("Food: %s\n", a.food)
}

// Move method print the animal’s locomotion
func (a Animal) Move() {
	fmt.Printf("Locomotion: %s\n", a.locomotion)
}

// Speak method print the animal’s spoken sound
func (a Animal) Speak() {
	fmt.Printf("Spoken sound: %s\n", a.noise)
}

func handleRequest(name, req string) {
	var animal Animal

	switch name {
	case "cow":
		animal = Animal{"grass", "walk", "moo"}
	case "bird":
		animal = Animal{"worms", "fly", "peep"}
	case "snake":
		animal = Animal{"mice", "slither", "hsss"}
	default:
		fmt.Println("--- Animal name must be one of: cow, bird, snake")
	}

	switch req {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("--- Your request must be one of: eat, move, speak")
	}
}

func input() (string, string) {
	var name, req string
	var r = bufio.NewReader(os.Stdin)
	fmt.Print("Enter name and request in one single line: ")
	fmt.Fscanf(r, "%s %s ", &name, &req)

	return name, req
}

func main() {
	for true {
		var name, req string
		name, req = input()

		handleRequest(name, req)
	}
}
