package main

import "fmt"

func userInput() float64 {
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func calculateConversion(amount float64, fromCurrency string, toCurrency string) float64 {
	fmt.Println(amount, fromCurrency, toCurrency)
	return 0.0
}

func main() {
	const (
		usdToEur = 0.92
		usdToRub = 90.0
		eurToRub = usdToRub / usdToEur
	)
	fmt.Print(eurToRub)
	var userInput float64 = userInput()
	fmt.Print(calculateConversion(userInput, "USD", "EUR"))
}
