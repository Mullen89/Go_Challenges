package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // ignore this comment

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank"

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var x uint64
	var y float64
	var z string
	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	x, _ = strconv.ParseUint(scanner.Text(), 10, 10)
	scanner.Scan()
	y, _ = strconv.ParseFloat(scanner.Text(), 10)
	scanner.Scan()
	z = scanner.Text()
	// Print the sum of both integer variables on a new line.
	fmt.Printf("%v\n", i+x)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+y)
	// Concatenate and print the String variables on a new line
	fmt.Printf("%s%s\n", s, z)
	// The 's' variable above should be printed first.
}
