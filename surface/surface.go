package surface

import (
	"math"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

type Surface struct {
	ID        int
	Nodes     []c.Node
	NumPoints int
	Ksi       []float64
	Eta       []float64
	N         []*mat.VecDense
}

func NewSurface(ID int, nodes []c.Node, numPoints int, Ksi []float64, Eta []float64) Surface {
	return Surface{
		ID:        ID,
		Nodes:     nodes,
		NumPoints: numPoints,
		Ksi:       Ksi,
		Eta:       Eta,
		N:         calculateN(numPoints, Ksi, Eta),
	}
}

func (s Surface) CalculateDetJ() float64 {
	dx := s.Nodes[1].X - s.Nodes[0].X
	dy := s.Nodes[1].Y - s.Nodes[0].Y

	return math.Sqrt(dx*dx+dy*dy) / 2.0
}

func calculateN(numPoints int, Ksi []float64, Eta []float64) []*mat.VecDense {
	N := make([]*mat.VecDense, numPoints)

	for i := 0; i < numPoints; i++ {
		N[i] = mat.NewVecDense(4, nil)

		N[i].SetVec(0, 0.25*(1-Ksi[i])*(1-Eta[i]))
		N[i].SetVec(1, 0.25*(1+Ksi[i])*(1-Eta[i]))
		N[i].SetVec(2, 0.25*(1+Ksi[i])*(1+Eta[i]))
		N[i].SetVec(3, 0.25*(1-Ksi[i])*(1+Eta[i]))
	}

	return N
}

func (s Surface) CalculateHbcMatrix(alpha float64) *mat.Dense {
	Hbc := mat.NewDense(4, 4, nil)

	if !s.Nodes[0].BC || !s.Nodes[1].BC {
		return Hbc
	}

	for i := 0; i < s.NumPoints; i++ {
		weight := c.Points[s.NumPoints].Weights[i]
		N := s.N[i]

		HbcPc := mat.NewDense(4, 4, nil)
		HbcPc.Outer(1.0, N, N)
		HbcPc.Scale(weight*alpha, HbcPc)
		Hbc.Add(Hbc, HbcPc)
	}

	return Hbc
}
