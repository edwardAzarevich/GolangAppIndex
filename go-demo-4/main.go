package main

import "fmt"

func main() {
	a := 2
	double(&a)
	fmt.Println(a)
}

func double(num *int) {
	*num = *num * 2
}

func reverse(arr *[]int) {
	s := *arr
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
