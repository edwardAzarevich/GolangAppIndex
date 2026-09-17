package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 = 1.9
	var userkg float64 = 89
	var IMT = userkg / math.Pow(userHeight, 2)
	fmt.Printf("Halo, your IMT is %.2f\n", IMT)
}
