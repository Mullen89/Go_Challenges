package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var temp_num uint64
	var num rune
	var str string

	scanner.Scan()
	str = scanner.Text()

	scanner.Scan()
	temp_num, _ = strconv.ParseUint(scanner.Text(), 0, 0)
	num = rune(temp_num)

	fmt.Printf("%s\n", CaesarCipher(str, num))

}

func CaesarCipher(s string, n rune) string {
	r := []rune(s)
	var cipher []rune
	for i := 0; i < len(s); i++ {
		temp_r := r[i] + n
		cipher = append(cipher, temp_r)
	}
	return string(cipher)
}
