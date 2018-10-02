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

func gradingStudents(grades []int32) []int32 {
	var gradeArr []int32
	for _, v := range grades {
		if v < 38 {
			gradeArr = append(gradeArr, v)
		} else if (math.Ceil(float64(v)/5)*5)-float64(v) < 3 {
			gradeArr = append(gradeArr, int32((math.Ceil(float64(v)/5) * 5)))
		} else {
			gradeArr = append(gradeArr, v)
		}
	}
	return gradeArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grades []int32

	for gradesItr := 0; gradesItr < int(n); gradesItr++ {
		gradesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
