package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var inputReader = bufio.NewReader(os.Stdin)

func enterTransaction() []float64 {
	fmt.Println("Please enter your array:")
	for {
		userInput, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("Exit the program. Good luck!")
			break
		}
		stringNumbers := strings.Fields(userInput)
		var numbers []float64
		isValid := true
		for _, str := range stringNumbers {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Printf("Error: %q — not a number. Please try again.\n", str)
				isValid = false
				break
			}
			numbers = append(numbers, num)
		}
		if !isValid {
			continue
		}
		return numbers
	}
	return []float64{}
}

func userOperationInput() string {
	validCurrencies := map[string]bool{
		"avg": true,
		"sum": true,
		"med": true,
	}
	fmt.Println("Please choose target operation: AVG, SUM, MED")
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}

		currency := strings.ToLower(strings.TrimSpace(line))

		if !validCurrencies[currency] {
			fmt.Println("Invalid currency. Please enter: AVG, SUM, MED")
			continue
		}

		return currency
	}
}

func calcOperation(operation string, arr []float64) float64 {
	switch operation {
	case "sum":
		return sum(arr)
	case "avg":
		return avg(arr)
	case "med":
		return med(arr)
	default:
		return 0.0
	}
}

func sum(arr []float64) float64 {
	result := 0.0
	for _, v := range arr {
		result += v
	}
	return result
}

func avg(arr []float64) float64 {
	result := 0.0
	for _, v := range arr {
		result += v
	}
	result = result / float64(len(arr))
	return result
}

func med(arr []float64) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	cpy := slices.Clone(arr)

	slices.Sort(cpy)

	if n%2 == 1 {
		return cpy[n/2]
	} else {
		return (cpy[n/2-1] + cpy[n/2]) / 2
	}
}

func main() {
	fmt.Println("Welcome to calculation!")
	operation := userOperationInput()
	arr := enterTransaction()
	result := calcOperation(operation, arr)
	fmt.Printf("Your result operation %s: %0.2f", operation, result)
}
