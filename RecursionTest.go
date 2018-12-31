package main

import (
	"fmt"
)

func f(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + f(n-1)
}

func main() {
	fmt.Println(f(24))
}
