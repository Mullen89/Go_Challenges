package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Scan() {
		str := scanner.Text()

		stringSplitter(str)
	}
}
func stringSplitter(str string) {
	evens := make([]string, 0)
	odds := make([]string, 0)

	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			evens = append(evens, string(str[i]))
		} else {
			odds = append(odds, string(str[i]))
		}
	}
	evenStr := strings.Join(evens, "")
	oddStr := strings.Join(odds, "")

	fmt.Printf("%s %s\n", evenStr, oddStr)
}
