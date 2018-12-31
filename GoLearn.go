package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, world!")
	// rand.Seed(time.Now().UTC().UnixNano())
	// for i := 0; i < 11; i++ {
	// 	fmt.Println("Random number: ", rand.Intn(100))
	// }

	var r rune = 'z'

	fmt.Println(r)

	fmt.Println(((75 % 10) * 10) + (75 / 10))
}
