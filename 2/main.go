package main

import (
	"fmt"

	"mes/2/integration"
)

func main() {
	result1 := integration.GaussIntegration(integration.F1, 2)
	result2 := integration.GaussIntegration(integration.F2, 3)

	fmt.Printf("Result of 1st integral: %f\n", result1)
	fmt.Printf("Result of 2nd integral: %f\n", result2)
}
