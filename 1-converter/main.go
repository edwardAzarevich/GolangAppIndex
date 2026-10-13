package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usdToEur = 0.92
	usdToRub = 90.0
	eurToRub = usdToRub / usdToEur
)

var inputReader = bufio.NewReader(os.Stdin)

func amountInput() float64 {
	fmt.Print("Enter amount: ")

	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}

		line = strings.TrimSpace(line)
		amount, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Invalid number. Please enter a number, for example: 100 or 12.5")
			continue
		}

		if amount <= 0 {
			fmt.Println("Invalid number. Please enter more than 0, for example: 10 or 20.5")
			continue
		}

		return amount
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
		line, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}

		currency := strings.ToLower(strings.TrimSpace(line))

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
	fmt.Println("Welcome to process Converter!")
	fromCurrency := userCurrencyInput("")
	userValue := amountInput()
	toCurrency := userCurrencyInput(fromCurrency)
	result, err := calculateConversion(userValue, fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", userValue, fromCurrency, result, toCurrency)
}

func calculateConversion(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	fromCurrency = strings.ToLower(strings.TrimSpace(fromCurrency))
	toCurrency = strings.ToLower(strings.TrimSpace(toCurrency))
	switch fromCurrency {
	case "usd":
		switch toCurrency {
		case "euro":
			return amount * usdToEur, nil
		case "rub":
			return amount * usdToRub, nil
		}
	case "euro":
		switch toCurrency {
		case "usd":
			return amount / usdToEur, nil
		case "rub":
			return amount * eurToRub, nil
		}
	case "rub":
		switch toCurrency {
		case "usd":
			return amount / usdToRub, nil
		case "euro":
			return amount / eurToRub, nil
		}
	}

	return 0, fmt.Errorf("unsupported conversion: %s -> %s", fromCurrency, toCurrency)
}

func main() {
	processConversion()
}
