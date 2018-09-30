package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//prints min of first 4 numbers and max of last 4 numbers
func miniMaxSum(arr []int32) {
	fmt.Printf("%v %v\n", int64(arr[0])+int64(arr[1])+int64(arr[2])+int64(arr[3]), int64(arr[len(arr)-1])+int64(arr[len(arr)-2])+int64(arr[len(arr)-3])+int64(arr[len(arr)-4]))
}

//Sorts the arrays from min to max
func bubbleSort(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		numSwaps := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numSwaps++
			}
		}
		if numSwaps == 0 {
			break
		}
	}
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	arr2 := bubbleSort(arr)
	miniMaxSum(arr2)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
