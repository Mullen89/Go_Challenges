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
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	//Convert to binary, returned as a string
	strN := strconv.FormatInt(nTemp, 2)

	fmt.Printf("%d\n", consecutiveOnes(strN))
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

func consecutiveOnes(str string) int {
	count := 0
	tempCount := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "1" {
			tempCount++
		}
		if tempCount > count {
			count = tempCount
		}
		if string(str[i]) == "0" {
			tempCount = 0
		}
	}
	return count
}
