package main

import (
	"fmt"
	"math"
)

func outputResult(imt float64) string {
	result := fmt.Sprintf("Your index is: %.0f", imt)
	fmt.Println(result)
}

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("calc index")
	fmt.Print("Enter your Hieght in santimeters ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your Kg ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	outputResult(IMT)
}
