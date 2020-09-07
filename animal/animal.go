package main

import (
	"bufio"
	"fmt"
	"os"
)

// Animal interface
type Animal interface {
	Eat()
	Speak()
	Move()
}

// Cow class inherit Animal
type Cow struct {
	name, food, locomotion, noise string
}

// Eat method print the animal’s food
func (cow Cow) Eat() {
	fmt.Printf("Food: %s\n", cow.food)
}

// Move method print the animal’s locomotion
func (cow Cow) Move() {
	fmt.Printf("Locomotion: %s\n", cow.locomotion)
}

// Speak method print the animal’s spoken sound
func (cow Cow) Speak() {
	fmt.Printf("Spoken sound: %s\n", cow.noise)
}

// Bird class inherit Animal
type Bird struct {
	name, food, locomotion, noise string
}

// Eat method print the animal’s food
func (bird Bird) Eat() {
	fmt.Printf("Food: %s\n", bird.food)
}

// Move method print the animal’s locomotion
func (bird Bird) Move() {
	fmt.Printf("Locomotion: %s\n", bird.locomotion)
}

// Speak method print the animal’s spoken sound
func (bird Bird) Speak() {
	fmt.Printf("Spoken sound: %s\n", bird.noise)
}

// Snake class inherit Animal
type Snake struct {
	name, food, locomotion, noise string
}

// Eat method print the animal’s food
func (snake Snake) Eat() {
	fmt.Printf("Food: %s\n", snake.food)
}

// Move method print the animal’s locomotion
func (snake Snake) Move() {
	fmt.Printf("Locomotion: %s\n", snake.locomotion)
}

// Speak method print the animal’s spoken sound
func (snake Snake) Speak() {
	fmt.Printf("Spoken sound: %s\n", snake.noise)
}

// AnimalsData struct to store animals data using map
type AnimalsData struct {
	// animals map with key is the animal's name and value is the animal
	animals map[string]Animal
}

// Handle user request
func (aData AnimalsData) handleRequest(action, name, req string) {
	if action == "create" {
		aData.createNewAnimal(name, req)
		return
	}

	if action == "query" {
		aData.queryAnimals(name, req)
		return
	}

	fmt.Println("--- Action must be one of: query, create")
}

// Create a new animal and append it to animals map
func (aData AnimalsData) createNewAnimal(name, aType string) {
	var a Animal

	switch aType {
	case "cow":
		cow := Cow{name, "grass", "walk", "moo"}
		a = cow
	case "bird":
		bird := Bird{name, "worms", "fly", "peep"}
		a = bird
	case "snake":
		snake := Snake{name, "mice", "slither", "hsss"}
		a = snake
	default:
		fmt.Println("--- Animal must be one of: cow, bird, snake")
		return
	}

	aData.animals[name] = a
	fmt.Println("Created it!")
}

// Query animals data from map
func (aData AnimalsData) queryAnimals(name, req string) {
	if _, found := aData.animals[name]; found == false {
		fmt.Printf("--- Animal with name \"%s\" not found.\n", name)
		return
	}

	switch req {
	case "eat":
		aData.animals[name].Eat()
	case "move":
		aData.animals[name].Move()
	case "speak":
		aData.animals[name].Speak()
	default:
		fmt.Println("--- Your request must be one of: eat, move, speak")
	}
}

func input() (string, string, string) {
	var action, name, req string
	var r = bufio.NewReader(os.Stdin)
	fmt.Print("Enter action, name and request in one single line: ")
	fmt.Fscanf(r, "%s %s %s ", &action, &name, &req)

	return action, name, req
}

func main() {
	animals := AnimalsData{animals: make(map[string]Animal)}

	for true {
		action, name, req := input()

		animals.handleRequest(action, name, req)
	}
}
