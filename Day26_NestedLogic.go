package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//creates a scanner and places each input into a string array
	//that separates each value into a separate index
	scanner := bufio.NewScanner(os.Stdin)

	//array that contains the returned date
	scanner.Scan()
	returned := scanner.Text()
	returnedArr := strings.Fields(returned)

	//array that contains the due date
	scanner.Scan()
	due := scanner.Text()
	dueArr := strings.Fields(due)

	//initializes these two variables as an []int so we can work with them
	retIntArr, dueIntArr := arrParser(returnedArr, dueArr)

	//determines the fine to be paid
	fmt.Printf("%d", fineFinder(retIntArr, dueIntArr))
}

//determines the total fine to pay
func fineFinder(ret, due []int) int {
	if ret[2] < due[2] {
		return 0
	} else if ret[2] > due[2] {
		return 10000
	} else {
		if ret[1] < due[1] {
			return 0
		} else if ret[1] > due[1] {
			return 500 * (ret[1] - due[1])
		} else {
			if ret[0] < due[0] {
				return 0
			} else if ret[0] > due[0] {
				return 15 * (ret[0] - due[0])
			} else {
				return 0
			}
		}
	}
}

//parses two string arrays and returns two []int
func arrParser(ret, due []string) ([]int, []int) {
	x, y := make([]int, 0), make([]int, 0)
	for _, v := range ret {
		n, _ := strconv.ParseInt(v, 0, 0)
		x = append(x, int(n))
	}
	for _, v := range due {
		m, _ := strconv.ParseInt(v, 0, 0)
		y = append(y, int(m))
	}
	return x, y
}
