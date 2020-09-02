package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Using bufio to accept white spaces
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Input your address: ")
	scanner.Scan()
	addr := scanner.Text()

	personMap := map[string]string{
		"name": name,
		"addr": addr,
	}

	personJSON, _ := json.Marshal(personMap)
	fmt.Println(string(personJSON))
}
