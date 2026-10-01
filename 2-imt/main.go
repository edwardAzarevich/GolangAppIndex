package main

import (
	"fmt"
	"math"
)

func outputResult(imt float64) {
	result := fmt.Sprintf("Your index is: %.0f%%", imt)
	fmt.Println(result)
}

func repeat() bool {
	var choice string
	fmt.Println("Would you like to continue? yes/no")
	fmt.Scan(&choice)
	return choice == "yes"
}

func processIMTCalc() {
	fmt.Println("Welcome to the IMT Calculator")
	for {
		userHeight, userKg := getUserInput()
		IMT := calculateIMT(userHeight, userKg)
		outputResult(IMT)
		if !repeat() {
			break
		}
	}
}

func calculateIMT(userHeight, userKg float64) (IMT float64) {
	const IMTPower = 2
	IMT = userKg / math.Pow(userHeight/100, IMTPower)
	switch {
	case IMT < 16:
		fmt.Println("You are very underweight")
	case IMT < 18.5:
		fmt.Println("You are underweight")
	case IMT < 25:
		fmt.Println("You are normal")
	case IMT < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("You have a degree of obesity")
	}
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
	processIMTCalc()
	// userHeight, userKg := getUserInput()
	// IMT := calculateIMT(userHeight, userKg)
	// outputResult(IMT)
}
