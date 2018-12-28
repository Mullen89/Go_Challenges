package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	fmt.Println(hourglassSum(arr))
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

func hourglassSum(array [][]int32) int32 {
	var tempSum int32 = 0

	var i int32 = 0
	var j int32 = 1

	tempSum = array[0][0+i] + array[0][1+i] + array[0][2+i]
	tempSum += array[1][1+i]
	tempSum += array[2][0+i] + array[2][1+i] + array[2][2+i]

	var sum int32 = tempSum

	for i < 4 {
		for j < 4 {
			tempSum = array[j][0+i] + array[j][1+i] + array[j][2+i]
			tempSum += array[j+1][1+i]
			tempSum += array[j+2][0+i] + array[j+2][1+i] + array[j+2][2+i]

			if tempSum > sum {
				sum = tempSum
			}
			tempSum = 0
			j++
		}
		j = 0
		i++
	}
	return sum
}
