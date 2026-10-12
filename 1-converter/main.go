package main

import (
	"fmt"
	"strings"
)

func userInput() float64 {
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func userCurrencyInput() string {
	validCurrencies := map[string]bool{
		"usd":  true,
		"euro": true,
		"rub":  true,
	}
	for {
		fmt.Println("Please choose you currency: usd, euro, rub")
		var currency string
		fmt.Scan(&currency)

		currency = strings.ToLower(strings.TrimSpace(currency))
		if validCurrencies[currency] {
			return currency
		}
		fmt.Println("Invalid currency. Please enter: usd, euro, or rub")
	}
}

func processConversion() {
	fmt.Println("Welcome to priccess Converter!")
	for {
		userCurrency := userCurrencyInput()
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
