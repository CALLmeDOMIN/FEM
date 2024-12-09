package main

import (
	"fmt"
	"os"

	"mes/common"
	"mes/grid"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	integrationPoints := 3

	gridFile, globalDataFile, err := common.ReadFromFile(file)
	if err != nil {
		fmt.Println("Error reading from file: ", err)
		return
	}

	simulationGrid, globalData := grid.GenerateGrid(globalDataFile, gridFile, integrationPoints)

	common.PrintGlobalData(globalData)
	common.PrintGrid(simulationGrid)
}
