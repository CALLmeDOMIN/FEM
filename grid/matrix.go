package grid

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
	i "mes/integration"
)

func calculateHMatrix_global(grid c.Grid) *mat.Dense {
	H := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)

	for _, element := range grid.Elements {
		H_local := element.HMatrix

		for i, globalIDi := range element.IDs {
			for j, globalIDj := range element.IDs {
				H.Set(globalIDi-1, globalIDj-1, H.At(globalIDi-1, globalIDj-1)+H_local.At(i, j))
			}
		}
	}

	return H
}

func calculateHMatrix_local(element c.Element, nodeMap map[int]c.Node, conductivity float64, points int) *mat.Dense {
	weights := c.Points[points].Weights
	H := mat.NewDense(len(element.IDs), len(element.IDs), nil)

	jacobians := i.CalculateJacobian(element, nodeMap, points)
	dets := i.CalculateDetJacobian(jacobians)
	inverses := i.CalculateReverseJacobian(jacobians)

	dNdx := make([][]float64, points*points)
	dNdy := make([][]float64, points*points)

	for i := 0; i < points*points; i++ {
		dNdx[i] = make([]float64, 4)
		dNdy[i] = make([]float64, 4)

		for j := 0; j < 4; j++ {
			dNdx[i][j] = inverses[i].At(0, 0)*element.DNdKsi[i][j] + inverses[i].At(0, 1)*element.DNdEta[i][j]
			dNdy[i][j] = inverses[i].At(1, 0)*element.DNdKsi[i][j] + inverses[i].At(1, 1)*element.DNdEta[i][j]
		}
	}

	for i := 0; i < points*points; i++ {
		detJ := dets[i]
		weightX := weights[i%points]
		weightY := weights[i/points]

		for m := 0; m < 4; m++ {
			for n := 0; n < 4; n++ {
				H.Set(m, n, H.At(m, n)+conductivity*(dNdx[i][m]*dNdx[i][n]+dNdy[i][m]*dNdy[i][n])*detJ*weightX*weightY)
			}
		}
	}

	fmt.Println("H:")
	c.PrintMatrix(H)

	return H
}
