package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var temp_x uint64
	var x float64

	scanner.Scan()
	temp_x, _ = strconv.ParseUint(scanner.Text(), 0, 0)
	x = float64(temp_x)

	fmt.Printf("%t\n", PowersOfTwo(x))

}

func PowersOfTwo(num float64) bool {
	fmt.Println(num)
	if num/2 > 2 {
		return PowersOfTwo(num / 2)
	} else if math.Mod(num, 2) == 0 {
		return true
	} else {
		return false
	}
}
