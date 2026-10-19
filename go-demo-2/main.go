package main

import "fmt"

func main() {
	transaction := [5]int{1, 2, 3, 4, 5}
	partial := transaction[1:4]
	fmt.Println(partial)
}
