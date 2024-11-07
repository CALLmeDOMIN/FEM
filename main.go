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

	integrationPoints := 3

	gridFile, globalDataFile, err := c.ReadFromFile(file)
	if err != nil {
		fmt.Println("Error reading from file: ", err)
		return
	}

	grid, globalData := c.GenerateGrid(globalDataFile, gridFile, integrationPoints)

	fmt.Printf("GlobalData: %v\n", globalData)
	fmt.Printf("Grid: %v\n", grid)

	nodeMap := make(map[int]c.Node)
	for i, node := range grid.Nodes {
		nodeMap[i+1] = node
	}

	for _, element := range grid.Elements {
		H := i.CalculateHMatrix(element, nodeMap, globalData.Conductivity, integrationPoints)
		fmt.Printf("H matrix for element %v:\n", element.IDs)
		c.PrintMatrix(H)
	}
}
