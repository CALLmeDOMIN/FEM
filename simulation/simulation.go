package simulation

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

func SimulateTemperature(grid c.Grid, globalData c.GlobalData) []mat.VecDense {
	timeSteps := globalData.SimulationTime / globalData.SimulationStepTime
	temperatureHistory := make([]mat.VecDense, timeSteps+1)

	t0 := mat.NewVecDense(grid.NodesNumber, nil)
	for i := range grid.Nodes {
		t0.SetVec(i, globalData.InitialTemperature)
	}
	temperatureHistory[0] = *t0

	for step := 0; step < timeSteps; step++ {
		t := CalculateTemperatureNonStationary(grid, globalData)
		temperatureHistory[step+1] = t

		for i := range grid.Nodes {
			grid.Nodes[i].Temperature = t.AtVec(i)
		}
	}

	return temperatureHistory
}

func CalculateTemperatureNonStationary(grid c.Grid, globalData c.GlobalData) mat.VecDense {
	dTau := float64(globalData.SimulationStepTime)

	A := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)
	CScaled := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)
	CScaled.Scale(1/dTau, grid.CMatrix)
	A.Add(grid.HMatrix, CScaled)

	t0 := mat.NewVecDense(grid.NodesNumber, nil)
	for i := range grid.Nodes {
		t0.SetVec(i, grid.Nodes[i].Temperature)
	}

	B := mat.NewVecDense(grid.NodesNumber, nil)
	B.MulVec(CScaled, t0)
	B.AddVec(B, grid.PVector)

	var t mat.VecDense
	err := t.SolveVec(A, B)
	if err != nil {
		panic(fmt.Sprintf("Cannot solve the equation: %v", err))
	}

	return t
}
