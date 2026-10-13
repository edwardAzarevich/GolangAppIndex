package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.92
	usdToRub = 90.0
	eurToRub = usdToRub / usdToEur
)

func userInput() float64 {
	fmt.Print("Enter amount: ")
	for {
		var amount float64
		_, err := fmt.Scan(&amount)
		if err == nil {
			return amount
		}
		fmt.Println("Invalid number. Please enter a number, for example: 100 or 12.5")
	}
}

func userCurrencyInput(exclude string) string {
	validCurrencies := map[string]bool{
		"usd":  true,
		"euro": true,
		"rub":  true,
	}
	if exclude == "" {
		fmt.Println("Please choose your currency: usd, euro, rub")
	} else {
		fmt.Println("Please choose target currency: usd, euro, rub")
	}
	for {
		var currency string
		fmt.Scan(&currency)

		currency = strings.ToLower(strings.TrimSpace(currency))

		if !validCurrencies[currency] {
			fmt.Println("Invalid currency. Please enter: usd, euro, or rub")
			continue
		}

		if currency == exclude {
			fmt.Println("Target currency must be different from source currency.")
			continue
		}

		return currency
	}
}

func processConversion() {
	fmt.Println("Welcome to priccess Converter!")
	fromCurrency := userCurrencyInput("")
	userValue := userInput()
	toCurrency := userCurrencyInput(fromCurrency)
	result := (calculateConversion(userValue, fromCurrency, toCurrency))
	fmt.Printf("%.2f %s = %.2f %s\n", userValue, fromCurrency, result, toCurrency)
}

func calculateConversion(amount float64, fromCurrency string, toCurrency string) float64 {
	fromCurrency = strings.ToLower(strings.TrimSpace(fromCurrency))
	toCurrency = strings.ToLower(strings.TrimSpace(toCurrency))
	switch fromCurrency {
	case "usd":
		switch toCurrency {
		case "euro":
			return amount * usdToEur
		case "rub":
			return amount * usdToRub
		}
	case "euro":
		switch toCurrency {
		case "usd":
			return amount / usdToEur
		case "rub":
			return amount * eurToRub
		}
	case "rub":
		switch toCurrency {
		case "usd":
			return amount / usdToRub
		case "euro":
			return amount / eurToRub
		}
	}

	return 0.0
}

func main() {
	processConversion()
}
