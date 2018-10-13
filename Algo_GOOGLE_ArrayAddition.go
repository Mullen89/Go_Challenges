package main

import (
	"fmt"
	"sort"
)

func ArrayAddition(arr []int) bool {

	//finds max value in arr
	max := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			index = i
		}
	}
	//removes max value from arr
	arr = append(arr[:index], arr[index+1:]...)

	//Sorts the array from largest to smallest int
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	//fmt.Println(arr)

	//Finds if any combination of ints from arr equal the max value
	sum := 0
	for j := 0; j < len(arr); j++ {
		sum = arr[j]
		for k := 0; k < len(arr); k++ {
			if j == k && k == len(arr)-1 {
				break
			} else if j == k && k < len(arr)-1 {
				k++
			}
			//fmt.Println(arr[j], arr[k])
			sum += arr[k]
			if sum > max {
				sum -= arr[k]
			}
			//fmt.Println(j, k, sum)
			//fmt.Println()
		}
		if sum == max {
			return true
		}
	}
	return false
}

func main() {
	array1 := []int{5, 7, 16, 1, 2}
	fmt.Printf("Expected: false, actual: %t\n", ArrayAddition(array1))

	array2 := []int{3, 5, -1, 8, 12}
	fmt.Printf("Expected: true, actual: %t\n", ArrayAddition(array2))

	array3 := []int{33, 50, 17, 25, 10, 48, 36, 28, 15, 11, 5, 2, 8, 6, 39}
	fmt.Printf("Expected: true, actual: %t\n", ArrayAddition(array3))
}
