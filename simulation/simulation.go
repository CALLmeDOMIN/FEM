package simulation

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

func SimulateTemperature(grid c.Grid, globalData c.GlobalData) *mat.VecDense {
	timeSteps := globalData.SimulationTime / globalData.SimulationStepTime

	t0 := mat.NewVecDense(grid.NodesNumber, nil)
	for i := range grid.Nodes {
		t0.SetVec(i, globalData.InitialTemperature)
	}

	for i := 0; i < timeSteps; i++ {
		t0 = calculateTemperatureDistribution(grid, globalData, t0)
	}

	return t0
}

func calculateTemperatureDistribution(grid c.Grid, globalData c.GlobalData, t0 *mat.VecDense) *mat.VecDense {
	dTau := float64(globalData.SimulationStepTime)
	CScaled := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)
	CScaled.Scale(1/dTau, grid.CMatrix)

	A := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)
	A.Add(grid.HMatrix, CScaled)

	B := mat.NewVecDense(grid.NodesNumber, nil)
	B.MulVec(CScaled, t0)
	B.AddVec(B, grid.PVector)

	var t1 mat.VecDense
	err := t1.SolveVec(A, B)
	if err != nil {
		panic(fmt.Sprintf("Cannot solve the equation: %v", err))
	}

	min := t1.RawVector().Data[0]
	max := t1.RawVector().Data[0]

	for _, data := range t1.RawVector().Data {
		if data < min {
			min = data
		}
		if data > max {
			max = data
		}
	}

	fmt.Printf("Min: %f, Max: %f\n", min, max)

	return &t1
}
