package main

import (
	"fmt"
	"math"
)

func outputResult(imt float64) {
	result := fmt.Sprintf("Your index is: %.0f%%", imt)
	fmt.Println(result)
}

func calculateIMT(userHeight, userKg float64) (IMT float64) {
	const IMTPower = 2
	IMT = userKg / math.Pow(userHeight/100, IMTPower)
	return
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Enter your Hieght in santimeters ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your Kg ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func main() {
	fmt.Println("calc index")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}
