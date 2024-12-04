package grid

import (
	"gonum.org/v1/gonum/mat"

	c "mes/common"
	i "mes/integration"
	s "mes/surface"
)

func calculateHMatrix_global(grid c.Grid) *mat.Dense {
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

func calculateHMatrix_local(element c.Element, nodeMap map[int]c.Node, conductivity float64, alpha float64, points int) *mat.Dense {
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
		ksi_vals := make([]float64, points)
		eta_vals := make([]float64, points)

		if nodes[0].BC && nodes[1].BC {
			if i%2 == 0 {
				ksi_vals = c.Points[points].Coords
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
				eta_vals = c.Points[points].Coords
			}
			surface := s.NewSurface(i+1, nodes, points, ksi_vals, eta_vals)
			surfaces = append(surfaces, surface)
		}
	}

	for _, surface := range surfaces {
		detJ := surface.CalculateDetJ()
		Hbc_surface := surface.CalculateHbcMatrix(alpha)

		Hbc_surface.Scale(detJ, Hbc_surface)

		Hbc.Add(Hbc, Hbc_surface)
	}

	return Hbc
}
