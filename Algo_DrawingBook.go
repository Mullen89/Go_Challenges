package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func pageCount(n int32, p int32) int32 {
	if n == p || p == 1 {
		return 0
	} else if n%2 != 0 && n-p == 1 {
		return 0
	} else if (n%2 == 0 && n-p <= 2) || ((p-1)/2 <= 1) {
		return 1
	} else {
		var tempAns float64 = 0
		if p-1 < n-p {
			tempAns = (float64(p) - 1) / 2
			return int32(math.Round(tempAns))
		} else {
			tempAns = (float64(n) - float64(p)) / 2
			if n%2 != 0 && p%2 == 0 {
				return int32(tempAns)
			} else {
				return int32(math.Round(tempAns))
			}
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

	fmt.Fprintf(writer, "%d\n", result)

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
