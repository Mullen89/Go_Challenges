package main

import (
	"bufio"
	"strings"
)

func CountStringWords(str string) int {
	// Creates a scanner to read over a string
	scan := bufio.NewScanner(strings.NewReader(str))

	// Split up the string
	scan.Split(bufio.ScanWords)

	// Can just write "count := 0"
	var count int = 0

	// A "for" loop set as a "while" loop
	// Counts the words in the string
	for scan.Scan() {
		count++
	}

	return count
}

// func main() {
// 	input := bufio.NewReader(os.Stdin)
// 	text, _ := input.ReadString('\n')

// 	fmt.Println(CountStringWords(text))
// }
