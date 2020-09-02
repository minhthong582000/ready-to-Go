package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Person Class
type Person struct {
	fname string
	lname string
}

// Read lines from a file.
// Then append each line to a slice of string
// Return the slice
func readLines(filepath string) ([]string, error) {
	// Open file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Split lines by space and create a Person object
// Then append that person to a slice
func personSlice(lines []string) []Person {
	var pers []Person

	for _, n := range lines {
		// Split first name and last name by space
		n := strings.Fields(n)

		// Append to slice of Person
		pers = append(pers, Person{fname: n[0], lname: n[1]})
	}

	return pers
}

func main() {
	var filepath string

	// Input file path
	fmt.Print("Enter filename or path to file: ")
	fmt.Scan(&filepath)

	// Read file and append each line to lines slice
	lines, err := readLines(filepath)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Print(personSlice(lines))
}
