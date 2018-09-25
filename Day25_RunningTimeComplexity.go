package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//read input and place test cases into an array
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numTests, _ := strconv.ParseInt(scanner.Text(), 0, 0)

	arr := make([]int, numTests)
	for i := 0; i < len(arr); i++ {
		scanner.Scan()
		n, _ := strconv.ParseInt(scanner.Text(), 0, 0)
		arr[i] = int(n)
	}
	isPrime(arr)
}

//Checks each element in int slice to determine if it is prime
func isPrime(arr []int) {
	for i := 0; i < len(arr); i++ {
		prime := true
		for j := 2; j <= int(math.Floor(math.Sqrt(float64(arr[i])))); j++ {
			if arr[i]%j == 0 {
				prime = false
				break
			}
		}
		if arr[i] == 1 {
			fmt.Println("Not prime")
		} else if arr[i] == 2 {
			fmt.Println("Prime")
		} else if prime == true {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
