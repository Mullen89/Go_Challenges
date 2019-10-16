package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Write anything: ")

	// Takes in input
	input := bufio.NewReader(os.Stdin)

	// Parses to text, or error if invalid input
	text, err := input.ReadString('\n')

	// Error handling
	if err != nil {
		panic(err)
	}

	// Test string
	var x string = "hello"

	fmt.Println(x)
	fmt.Println(text)
	// In order to use "CountStringWords(text)", you must run command: "go run 0-Hello.go CountWords.go"
	// Exported functions must always be capitalized
	fmt.Println("Number of words you wrote = ", CountStringWords(text))

}
