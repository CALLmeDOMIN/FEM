package grid

import (
	"gonum.org/v1/gonum/mat"

	c "mes/common"
	i "mes/integration"
	s "mes/surface"
)

func calculateHMatrixGlobal(grid c.Grid) *mat.Dense {
	H := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)

	for _, element := range grid.Elements {
		H_local := element.HMatrix

		for i, globalIDi := range element.NodeIDs {
			for j, globalIDj := range element.NodeIDs {
				H.Set(globalIDi-1, globalIDj-1, H.At(globalIDi-1, globalIDj-1)+H_local.At(i, j))
			}
		}
	}

	return H
}

func calculateHMatrixLocal(element c.Element, nodeMap map[int]c.Node, conductivity float64, alpha float64, points int) *mat.Dense {
	H := mat.NewDense(len(element.NodeIDs), len(element.NodeIDs), nil)
	weights := c.Points[points].Weights

	jacobians := i.CalculateJacobians(element, nodeMap, points)
	dets := i.CalculateDetJacobians(jacobians)
	inverses := i.CalculateReverseJacobians(jacobians)

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

		scale := conductivity * detJ * weightX * weightY

		for m := 0; m < 4; m++ {
			for n := 0; n < 4; n++ {
				value := (dNdx[i][m]*dNdx[i][n] + dNdy[i][m]*dNdy[i][n]) * scale
				H.Set(m, n, H.At(m, n)+value)
			}
		}
	}

	Hbc := calculateHbcMatrix(element, nodeMap, points, alpha)
	H.Add(H, Hbc)

	return H
}

func calculateHbcMatrix(element c.Element, nodeMap map[int]c.Node, points int, alpha float64) *mat.Dense {
	Hbc := mat.NewDense(len(element.NodeIDs), len(element.NodeIDs), nil)
	surfaces := []s.Surface{}

	for i := 0; i < 4; i++ {
		nodes := []c.Node{nodeMap[element.NodeIDs[i]], nodeMap[element.NodeIDs[(i+1)%4]]}
		ksiVals := make([]float64, points)
		etaVals := make([]float64, points)

		if nodes[0].BC && nodes[1].BC {
			if i%2 == 0 {
				ksiVals = c.Points[points].Coords
				if i == 0 {
					for j := 0; j < points; j++ {
						etaVals[j] = -1
					}
				} else {
					for j := 0; j < points; j++ {
						etaVals[j] = 1
					}
				}
			} else {
				if i == 1 {
					for j := 0; j < points; j++ {
						ksiVals[j] = 1
					}
				} else {
					for j := 0; j < points; j++ {
						ksiVals[j] = -1
					}
				}
				etaVals = c.Points[points].Coords
			}
			surface := s.NewSurface(i+1, nodes, points, ksiVals, etaVals)
			surfaces = append(surfaces, surface)
		}
	}

	for _, surface := range surfaces {
		detJ := surface.CalculateDetJ()
		HbcSurface := surface.CalculateHbcMatrix(alpha)

		HbcSurface.Scale(detJ, HbcSurface)

		Hbc.Add(Hbc, HbcSurface)
	}

	return Hbc
}

func calculateCMatrixGlobal(grid c.Grid) *mat.Dense {
	C := mat.NewDense(grid.NodesNumber, grid.NodesNumber, nil)

	for _, element := range grid.Elements {
		C_local := element.CMatrix

		for i, globalIDi := range element.NodeIDs {
			for j, globalIDj := range element.NodeIDs {
				C.Set(globalIDi-1, globalIDj-1, C.At(globalIDi-1, globalIDj-1)+C_local.At(i, j))
			}
		}
	}

	return C
}

func calculateCMatrixLocal(element c.Element, nodeMap map[int]c.Node, density float64, specificHeat float64, points int) *mat.Dense {
	C := mat.NewDense(len(element.NodeIDs), len(element.NodeIDs), nil)
	weights := c.Points[points].Weights

	jacobians := i.CalculateJacobians(element, nodeMap, points)
	dets := i.CalculateDetJacobians(jacobians)

	for i := 0; i < points*points; i++ {
		detJ := dets[i]
		weightX := weights[i%points]
		weightY := weights[i/points]

		scale := density * specificHeat * detJ * weightX * weightY

		for m := 0; m < 4; m++ {
			for n := 0; n < 4; n++ {
				value := element.N[i][m] * element.N[i][n] * scale
				C.Set(m, n, C.At(m, n)+value)
			}
		}
	}

	return C
}
