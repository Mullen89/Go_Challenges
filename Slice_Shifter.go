package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(len(s), s)
	for i := 0; i < 3; i++ {
		x := s[0]        // get the 0 index element from slice
		s = s[1:]        // remove the 0 index element from slice
		s = append(s, x) // print 0 index element
	}
	fmt.Println(s)
}
