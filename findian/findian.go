package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Using bufio to accept white spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var x string = scanner.Text()

	// Remove case-sensitive
	var xLow = strings.ToLower(x)

	/**
	 * print “Found!”
	 * if the entered string:
	 * 		Starts with the character ‘i’,
	 * 		Ends with the character ‘n’,
	 * 		Contains the character ‘a’.
	 * Otherwise, print "Not found"
	 */
	if strings.HasPrefix(xLow, "i") && strings.HasSuffix(xLow, "n") && strings.Contains(xLow, "a") {
		fmt.Println("Found.")
	} else {
		fmt.Printf("Not found.")
	}
}
