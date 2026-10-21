package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputReader = bufio.NewReader(os.Stdin)

func enterTransaction() []float64 {
	fmt.Println("Please Enter your transactions")
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

func main() {
	inputNumbers := enterTransaction()
	fmt.Printf("Your transaction: %v", inputNumbers)
}
