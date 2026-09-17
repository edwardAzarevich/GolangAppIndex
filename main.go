package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userkg := 1.9, 89.0
	IMT := userkg / math.Pow(userHeight, 2)
	fmt.Printf("Halo, your IMT is %.2f\n", IMT)
}
