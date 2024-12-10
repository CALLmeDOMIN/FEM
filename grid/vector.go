package grid

import (
	"gonum.org/v1/gonum/mat"

	"mes/common"
	"mes/surface"
)

func calculatePVectorGlobal(grid common.Grid) *mat.VecDense {
	P := mat.NewVecDense(grid.NodesNumber, nil)

	for _, element := range grid.Elements {
		P_local := element.PVector

		for i, globalID := range element.NodeIDs {
			P.SetVec(globalID-1, P.AtVec(globalID-1)+P_local.AtVec(i))
		}
	}

	return P
}

func calculatePVectorLocal(element common.Element, nodeMap map[int]common.Node, alpha float64, ambientTemperature float64, points int) *mat.VecDense {
	P := mat.NewVecDense(len(element.NodeIDs), nil)
	surfaces := []surface.Surface{}

	for i := 0; i < 4; i++ {
		nodes := []common.Node{nodeMap[element.NodeIDs[i]], nodeMap[element.NodeIDs[(i+1)%4]]}
		ksi_vals := make([]float64, points)
		eta_vals := make([]float64, points)

		if nodes[0].BC && nodes[1].BC {
			if i%2 == 0 {
				ksi_vals = common.Points[points].Coords
				if i == 0 {
					for j := 0; j < points; j++ {
						eta_vals[j] = -1
					}
				} else {
					for j := 0; j < points; j++ {
						eta_vals[j] = 1
					}
				}
			} else {
				if i == 1 {
					for j := 0; j < points; j++ {
						ksi_vals[j] = 1
					}
				} else {
					for j := 0; j < points; j++ {
						ksi_vals[j] = -1
					}
				}
				eta_vals = common.Points[points].Coords
			}
			surface := surface.NewSurface(i+1, nodes, points, ksi_vals, eta_vals)
			surfaces = append(surfaces, surface)
		}
	}

	for _, surface := range surfaces {
		detJ := surface.CalculateDetJ()
		Pbc := surface.CalculatePbcVector(alpha, ambientTemperature)

		Pbc.ScaleVec(detJ, Pbc)

		P.AddVec(P, Pbc)
	}

	return P
}
