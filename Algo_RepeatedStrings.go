package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	var count int64 = 0
	var total int64 = 0
	var capacity int64 = n / int64(len(s))
	var remainder int64 = n % int64(len(s))

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" {
			count++
		}
	}

	total = count * capacity
	if remainder != 0 {
		for j := 0; j < int(remainder); j++ {
			if string(s[j]) == "a" {
				total++
			}
		}
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
