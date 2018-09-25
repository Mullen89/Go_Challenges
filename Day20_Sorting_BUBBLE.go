package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	//Reads first number from input
	numElem, _ := strconv.ParseInt(scanner.Text(), 0, 0)
	//Creates an int[] with the length of the given number of elements
	arr := make([]int, numElem)
	scanner.Scan()
	elem := scanner.Text()
	elemArr := strings.Fields(elem)

	//Inserts all integers from input into an array to be sorted
	for i := 0; i < len(arr); i++ {

		n, _ := strconv.ParseInt(elemArr[i], 0, 0)
		arr[i] = int(n)
	}
	bubbleSort(arr)
}

func bubbleSort(arr []int) {
	totalNumSwaps := 0
	for i := 0; i < len(arr); i++ {
		numSwaps := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numSwaps++
				totalNumSwaps++
			}
		}
		if numSwaps == 0 {
			break
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", totalNumSwaps)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}
