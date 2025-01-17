package main

import (
	"fmt"
	"os"

	"mes/common"
	"mes/grid"
	"mes/simulation"
)

func main() {
	file, err := os.Open("data2.json")
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

	// common.PrintGlobalData(globalData)
	// common.PrintGrid(simulationGrid)

	temperatureHistory := simulation.SimulateTemperature(simulationGrid, globalData)
	for step, temps := range temperatureHistory {
		fmt.Printf("Time step %d:\n", step)
		common.PrintVector(&temps)
	}
}
