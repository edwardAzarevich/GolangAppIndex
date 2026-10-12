package main

import "fmt"

func userInput() float64 {
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func userCurrensyInput() string {
	fmt.Println("Please choose you currency: usd, euro, rub")
	var currency string
	fmt.Scan(&currency)
	return currency
}

func processConversion() {
	fmt.Println("Welcome to priccess Converter!")
	for {
		userCurrency := userCurrensyInput()
	}
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
