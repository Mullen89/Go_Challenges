package main

import (
	"fmt"
)

var arr []string

func BalanceParens(left int, right int, str string) {
	if left == 0 && right == 0 {
		arr = append(arr, str)
	}
	if left > 0 {
		BalanceParens(left-1, right+1, str+"(")
	}
	if right > 0 {
		BalanceParens(left, right-1, str+")")
	}
}

func main() {
	BalanceParens(3, 0, "")
	fmt.Printf("%v\n", arr)
}
