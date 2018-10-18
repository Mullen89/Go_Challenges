package main

import (
	"fmt"
	"sort"
)

func MissingNum(arr []int) int {
	sort.Ints(arr)
	numCounter := arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-numCounter == 1 {
			numCounter++
		} else {
			return numCounter + 1
		}
	}
	return 666
}

func main() {
	numArr := []int{-6, -3, -5, -4, -8, -9, 0, -1, -2}
	numArr2 := []int{1, 2, 5, 3}

	fmt.Printf("%v\n", MissingNum(numArr))
	fmt.Printf("%v\n", MissingNum(numArr2))
}
