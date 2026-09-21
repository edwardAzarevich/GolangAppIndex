package main

import "fmt"

func main() {
	const (
		usdToEur = 0.92
		usdToRub = 90.0
		eurToRub = usdToRub / usdToEur
	)
	fmt.Print(eurToRub)
}
