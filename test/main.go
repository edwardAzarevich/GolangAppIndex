package main

import "fmt"

func main() {
	arr := [5]int{100, 200, 300, 400, 500}
	firstEl := 50
	slice := arr[1:4]
	slice[0] = firstEl
	fmt.Println(arr, len(slice), cap(slice))
}
