package integration

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

func CalculateJacobians(geometry c.Geometry, nodeMap map[int]c.Node, points int) []*mat.Dense {
	jacobians := make([]*mat.Dense, 0)

	for i := 0; i < points*points; i++ {
		jacobian := mat.NewDense(2, 2, nil)

		for j := 0; j < 4; j++ {
			nodeID := geometry.GetIDs()[j]
			x, y := nodeMap[nodeID].X, nodeMap[nodeID].Y

			jacobian.Set(0, 0, jacobian.At(0, 0)+geometry.GetDNdKsi(i)[j]*x)
			jacobian.Set(0, 1, jacobian.At(0, 1)+geometry.GetDNdKsi(i)[j]*y)
			jacobian.Set(1, 0, jacobian.At(1, 0)+geometry.GetDNdEta(i)[j]*x)
			jacobian.Set(1, 1, jacobian.At(1, 1)+geometry.GetDNdEta(i)[j]*y)
		}

		jacobians = append(jacobians, jacobian)
	}

	return jacobians
}

func CalculateDetJacobians(jacobians []*mat.Dense) []float64 {
	dets := make([]float64, 0)

	for _, jacobian := range jacobians {
		dets = append(dets, mat.Det(jacobian))
	}

	return dets
}

func CalculateReverseJacobians(jacobians []*mat.Dense) []*mat.Dense {
	inverses := make([]*mat.Dense, 0)

	for _, jacobian := range jacobians {
		var matInverse mat.Dense
		err := matInverse.Inverse(jacobian)
		if err != nil {
			fmt.Println("Error calculating inverse: ", err)
			continue
		}

		inverses = append(inverses, &matInverse)
	}

	return inverses
}
