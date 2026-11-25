package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var exchangeRate = map[string]float64{
	"usd":  1,
	"euro": 0.92,
	"rub":  90.0,
}

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

func userCurrencyInput(rates *map[string]float64, exclude string) string {
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

		if _, ok := (*rates)[currency]; !ok {
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

func processConversion(rates *map[string]float64) {
	fmt.Println("Welcome to process Converter!")
	fromCurrency := userCurrencyInput(rates, "")
	userValue := amountInput()
	toCurrency := userCurrencyInput(rates, fromCurrency)
	result, err := calculateConversion(userValue, fromCurrency, toCurrency, rates)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", userValue, fromCurrency, result, toCurrency)
}

func calculateConversion(amount float64, fromCurrency string, toCurrency string, rates *map[string]float64) (float64, error) {
	fromCurrency = strings.ToLower(strings.TrimSpace(fromCurrency))
	toCurrency = strings.ToLower(strings.TrimSpace(toCurrency))

	fromRate, ok := (*rates)[fromCurrency]
	if !ok {
		return 0, fmt.Errorf("unsupported conversion: %s -> %s", fromCurrency, toCurrency)
	}

	toRate, ok := (*rates)[toCurrency]
	if !ok {
		return 0, fmt.Errorf("unsupported conversion: %s -> %s", fromCurrency, toCurrency)
	}

	return amount / fromRate * toRate, nil
}

func main() {
	processConversion(&exchangeRate)
}