package main

import (
	"fmt"
)

func main() {
	fmt.Println(GCD(12, 21))
}

func GCD(x, y int) int {
	if y%x == 0 {
		return x
	} else {
		return GCD(y%x, x)
	}
}
