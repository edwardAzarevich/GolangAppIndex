package main

import "fmt"

func userInput() float64 {
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func calculateConversion(userInput float64, current, target float64) float64 {
	return userInput * (target / current)
}

func main() {
	const (
		usdToEur = 0.92
		usdToRub = 90.0
		eurToRub = usdToRub / usdToEur
	)
	fmt.Print(eurToRub)
}
