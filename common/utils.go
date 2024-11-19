package common

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

var Points = map[int]struct {
	Coords  []float64
	Weights []float64
}{
	2: {
		Coords:  []float64{-math.Sqrt(1.0 / 3.0), math.Sqrt(1.0 / 3.0)},
		Weights: []float64{1, 1},
	},
	3: {
		Coords:  []float64{-math.Sqrt(3.0 / 5.0), 0, math.Sqrt(3.0 / 5.0)},
		Weights: []float64{5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0},
	},
}

func PrintGlobalData(globalData GlobalData) {
	fmt.Println("========================================")
	fmt.Println("Global data:")
	fmt.Printf("  Simulation time: %v\n", globalData.SimulationTime)
	fmt.Printf("  Simulation step time: %v\n", globalData.SimulationStepTime)
	fmt.Printf("  Conductivity: %v\n", globalData.Conductivity)
	fmt.Printf("  Alpha: %v\n", globalData.Alpha)
	fmt.Printf("  Ambient temperature: %v\n", globalData.AmbientTemperature)
	fmt.Printf("  Initial temperature: %v\n", globalData.InitialTemperature)
	fmt.Printf("  Density: %v\n", globalData.Density)
	fmt.Printf("  Specific heat: %v\n", globalData.SpecificHeat)
	fmt.Printf("  Nodes number: %v\n", globalData.NodesNumber)
	fmt.Printf("  Elements number: %v\n", globalData.ElementsNumber)
	fmt.Println("========================================")
}

func PrintGrid(grid Grid) {
	fmt.Println("========================================")
	fmt.Println("Grid:")
	fmt.Printf("  Nodes number: %v\n", grid.NodesNumber)
	fmt.Printf("  Elements number: %v\n", grid.ElementsNumber)
	fmt.Println("  Nodes:")
	for i, node := range grid.Nodes {
		fmt.Printf("    Node %v: x: %v, y: %v\n", i, node.X, node.Y)
	}
	fmt.Println("  Elements:")
	for i, element := range grid.Elements {
		fmt.Printf("\n    Element %v: \n", i)
		fmt.Printf("      IDs: %v\n", element.IDs)
		fmt.Printf("      Ksi: %v\n", element.Ksi)
		fmt.Printf("      Eta: %v\n", element.Eta)
		fmt.Printf("      DNdKsi: %v\n", element.DNdKsi)
		fmt.Printf("      DNdEta: %v\n", element.DNdEta)
		fmt.Printf("      HMatrix: \n")
		PrintMatrix(element.HMatrix)
	}
	fmt.Printf("  Width: %v\n", grid.Width)
	fmt.Printf("  Height: %v\n", grid.Height)
	fmt.Printf("  Number height: %v\n", grid.NumberHeight)
	fmt.Printf("  Number width: %v\n", grid.NumberWidth)
	fmt.Println("  HMatrix:")
	PrintMatrix(grid.HMatrix)
	fmt.Println("========================================")
}

func ReadFromFile(file *os.File) (Grid, GlobalData, error) {
	var grid Grid
	var globalData GlobalData

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&struct {
		*Grid
		*GlobalData
	}{
		&grid,
		&globalData,
	})
	if err != nil {
		return Grid{}, GlobalData{}, fmt.Errorf("error decoding file: %v", err)
	}

	return grid, globalData, nil
}

func PrintMatrix(matrix *mat.Dense) {
	fmt.Println(mat.Formatted(matrix, mat.Prefix(""), mat.Squeeze()))
}

func PrintMatrixArray(matrices []*mat.Dense) {
	for _, matrix := range matrices {
		PrintMatrix(matrix)
	}
}
