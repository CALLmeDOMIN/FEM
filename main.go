package main

import (
	"fmt"
	"os"

	c "mes/common"
	i "mes/integration"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	grid, globalData, err := c.ReadFromFile(file)
	if err != nil {
		fmt.Println("Error reading from file: ", err)
		return
	}

	nodeMap := make(map[int]c.Node)
	for i, node := range grid.Nodes {
		nodeMap[i+1] = node
	}

	for _, element := range grid.Elements {
		jacobians := i.CalculateJacobian(element, nodeMap)
		fmt.Printf("Jacobians for element %v:\n", element.IDs)
		c.PrintMatrixArray(jacobians)

		dets := i.CalculateDetJacobian(jacobians)
		fmt.Printf("Determinants for element %v: %v\n", element.IDs, dets)

		inverses := i.CalculateReverseJacobian(jacobians)
		fmt.Printf("Inverse Jacobians for element %v:\n", element.IDs)
		c.PrintMatrixArray(inverses)

		H := i.CalculateHMatrix(element, nodeMap, globalData.Conductivity)
		fmt.Printf("H matrix for element %v:\n", element.IDs)
		c.PrintMatrix(H)
	}

	fmt.Printf("GlobalData: %v\n", globalData)
	fmt.Printf("Grid: %v\n", grid)

	result1 := i.GaussIntegration(i.F1, 2)
	result2 := i.GaussIntegration(i.F2, 3)

	fmt.Printf("Result of 1st integral: %f\n", result1)
	fmt.Printf("Result of 2nd integral: %f\n", result2)
}
