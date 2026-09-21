package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	userHeight, userkg := 1.9, 89.0
	IMT := userkg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Halo, your IMT is %.2f\n", IMT)
}
