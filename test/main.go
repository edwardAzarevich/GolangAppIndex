package main

import "fmt"

func main() {
	numbers := [5]int{7, 14, 21, 28, 35}
	s := make([]int, 0, 3)

	_ = numbers
	_ = s

	for _, k := range numbers {
		s = append(s, k)
		fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	}
}
