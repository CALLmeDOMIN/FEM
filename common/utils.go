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

func ReadFromFile(file *os.File, integrationPoints int) (Grid, GlobalData, error) {
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

	grid.NodesNumber = globalData.NodesNumber
	grid.ElementsNumber = globalData.ElementsNumber

	if len(grid.Nodes) == 0 {
		grid.Nodes = GenerateNodes(grid.Width, grid.Height, grid.NumberWidth, grid.NumberHeight, grid.NodesNumber)
	}

	nodeMap := make(map[int]Node)
	for _, node := range grid.Nodes {
		nodeMap[node.ID] = node
	}

	if len(grid.Elements) == 0 {
		grid.Elements = GenerateElements(grid.NumberWidth, grid.NumberHeight, grid.ElementsNumber)
	}

	grid.Elements = GenerateShapeFunctionData(grid.Elements, grid.NumberWidth, grid.NumberHeight, integrationPoints)

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
