package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxSubArray(arr []int) (string, string) {
	maxEnding := 0
	maxSoFar := arr[0]

	for _, v := range arr {
		maxEnding = int(math.Max(float64(v), float64(maxEnding+v)))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEnding)))
	}
	//------------------------------------------------------------------
	sort.Ints(arr)
	sum := 0

	if arr[len(arr)-1] <= 0 {
		sum = arr[len(arr)-1]
	} else {
		for _, v := range arr {
			if v > 0 {
				sum += v
			}
		}
	}
	return strconv.Itoa(maxSoFar), strconv.Itoa(sum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < t; tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int

		for i := 0; i < n; i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arr = append(arr, arrItem)
		}
		x, y := maxSubArray(arr)
		writer.WriteString(x + " ")
		writer.WriteString(y + "\n")
	}
	writer.Flush()
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
