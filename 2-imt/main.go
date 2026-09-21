package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print("Enter your Hieght in meters ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your Kg ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
