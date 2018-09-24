package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Set up scanner and reads initial input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)

	//Creates the phonebook
	phoneBook := make(map[string]string)
	for i := 0; i < int(n); i++ {
		scanner.Scan()
		namePhone := scanner.Text()
		tempArr := strings.Fields(namePhone)
		phoneBook[tempArr[0]] = tempArr[1]
	}
	//Scans through rest of os.Stdin and adds remain names to an array/slice
	var nameArr []string
	for scanner.Scan() {
		nameArr = append(nameArr, scanner.Text())
	}
	//Prints out key:value pairs if they exist
	for i := 0; i < len(nameArr); i++ {
		tempKey := nameArr[i]
		if val, ok := phoneBook[tempKey]; ok {
			fmt.Printf("%s=%s\n", tempKey, val)
		} else {
			fmt.Println("Not found")
		}
	}
}
