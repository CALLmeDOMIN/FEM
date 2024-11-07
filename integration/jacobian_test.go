package integration

import (
	"testing"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

func TestCalculateHMatrix(t *testing.T) {
	const conductivity = 30.0

	type tcData struct {
		name    string
		element c.Element
		nodeMap map[int]c.Node
		points  int
		result  *mat.Dense
	}

	testCases := []tcData{
		{
			name: "should return correct H matrix",
			nodeMap: map[int]c.Node{
				1: {ID: 1, X: 0.0, Y: 0.0},
				2: {ID: 2, X: 0.025, Y: 0.0},
				3: {ID: 3, X: 0.025, Y: 0.025},
				4: {ID: 4, X: 0.0, Y: 0.025},
			},
			points: 3,
			element: c.Element{
				ID:  1,
				IDs: []int{1, 2, 3, 4},
				Ksi: []float64{-0.44365, -0.25, -0.05635, -0.44365, -0.25, -0.05635, -0.44365, -0.25, -0.05635},
				Eta: []float64{-0.444, -0.25, -0.056, -0.444, -0.25, -0.056, -0.444, -0.25, -0.056},
				DNdKsi: [][]float64{
					{-0.44365, 0.44365, 0.05635, -0.0564},
					{-0.25, 0.25, 0.25, -0.25},
					{-0.05635, 0.05635, 0.44365, -0.4436},
					{-0.44365, 0.44365, 0.05635, -0.0564},
					{-0.25, 0.25, 0.25, -0.25},
					{-0.05635, 0.05635, 0.44365, -0.4436},
					{-0.44365, 0.44365, 0.05635, -0.0564},
					{-0.25, 0.25, 0.25, -0.25},
					{-0.05635, 0.05635, 0.44365, -0.4436},
				},
				DNdEta: [][]float64{
					{-0.444, -0.056, 0.056, 0.443649},
					{-0.444, -0.056, 0.056, 0.443649},
					{-0.444, -0.056, 0.056, 0.443649},
					{-0.25, -0.25, 0.25, 0.25},
					{-0.25, -0.25, 0.25, 0.25},
					{-0.25, -0.25, 0.25, 0.25},
					{-0.056, -0.444, 0.444, 0.056351},
					{-0.056, -0.444, 0.444, 0.056351},
					{-0.056, -0.444, 0.444, 0.056351},
				},
			},
			result: mat.NewDense(4, 4, []float64{
				20, -5, -10, -5,
				-5, 20, -5, -10,
				-10, -5, 20, -5,
				-5, -10, -5, 20,
			}),
		},
		{
			name: "should return correct H matrix",
			nodeMap: map[int]c.Node{
				1: {ID: 1, X: 0.01, Y: -0.01},
				2: {ID: 2, X: 0.025, Y: 0},
				3: {ID: 3, X: 0.025, Y: 0.025},
				4: {ID: 4, X: 0, Y: 0.025},
			},
			points: 2,
			element: c.Element{
				IDs: []int{1, 2, 3, 4},
				Ksi: []float64{-0.57735, 0.57735, 0.57735, -0.57735},
				Eta: []float64{-0.57735, -0.57735, 0.57735, 0.57735},
				DNdKsi: [][]float64{
					{-0.394338, 0.394338, 0.105662, -0.10566},
					{-0.394338, 0.394338, 0.105662, -0.10566},
					{-0.105662, 0.105662, 0.394338, -0.39434},
					{-0.105662, 0.105662, 0.394338, -0.39434},
				},
				DNdEta: [][]float64{
					{-0.394338, -0.10566, 0.105662, 0.394338},
					{-0.105662, -0.39434, 0.394338, 0.105662},
					{-0.105662, -0.39434, 0.394338, 0.105662},
					{-0.394338, -0.10566, 0.105662, 0.394338},
				},
			},
			result: mat.NewDense(4, 4, []float64{
				20.563, -13.788, -9.436, 2.6619,
				-13.788, 28.304, -1.788, -12.726,
				-9.436, -1.788, 20.563, -9.338,
				2.6619, -12.726, -9.338, 19.402,
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := CalculateHMatrix(tc.element, tc.nodeMap, conductivity, tc.points)

			if !mat.EqualApprox(res, tc.result, 1e-2) {
				t.Errorf("Expected %v, got %v", tc.result, res)
			}
		})
	}
}

func TestCalculateJacobian(t *testing.T) {
	type tcData struct {
		name    string
		element c.Element
		nodeMap map[int]c.Node
		points  int
		result  []*mat.Dense
	}

	testCases := []tcData{
		{
			name: "should return [0.0125 0; 0 0.0125]",
			nodeMap: map[int]c.Node{
				1: {ID: 1, X: 0, Y: 0},
				2: {ID: 2, X: 0.025, Y: 0},
				3: {ID: 3, X: 0.025, Y: 0.025},
				4: {ID: 4, X: 0, Y: 0.025},
			},
			points: 1,
			element: c.Element{
				IDs: []int{1, 2, 3, 4},
				DNdKsi: [][]float64{
					{-0.39434, 0.394338, 0.105662, -0.10566},
					{-0.39434, 0.394338, 0.105662, -0.10566},
					{-0.10566, 0.105662, 0.394338, -0.39434},
					{-0.10566, 0.105662, 0.394338, -0.39434},
				},
				DNdEta: [][]float64{
					{-0.39434, -0.10566, 0.105662, 0.394338},
					{-0.10566, -0.39434, 0.394338, 0.105662},
					{-0.39434, -0.10566, 0.105662, 0.394338},
					{-0.10566, -0.39434, 0.394338, 0.105662},
				},
			},
			result: []*mat.Dense{
				mat.NewDense(2, 2, []float64{0.0125, 0, 0, 0.0125}),
			},
		},
		{
			name: "should return [0.008556624, 0.0039434; -0.00394338, 0.0164434]",
			nodeMap: map[int]c.Node{
				1: {ID: 1, X: 0.01, Y: -0.01},
				2: {ID: 2, X: 0.025, Y: 0},
				3: {ID: 3, X: 0.025, Y: 0.025},
				4: {ID: 4, X: 0, Y: 0.025},
			},
			points: 1,
			element: c.Element{
				IDs: []int{1, 2, 3, 4},
				DNdKsi: [][]float64{
					{-0.394338, 0.394338, 0.105662, -0.10566},
					{-0.394338, 0.394338, 0.105662, -0.10566},
					{-0.105662, 0.105662, 0.394338, -0.39434},
					{-0.105662, 0.105662, 0.394338, -0.39434},
				},
				DNdEta: [][]float64{
					{-0.394338, -0.10566, 0.105662, 0.394338},
					{-0.105662, -0.39434, 0.394338, 0.105662},
					{-0.394338, -0.10566, 0.105662, 0.394338},
					{-0.105662, -0.39434, 0.394338, 0.105662},
				},
			},
			result: []*mat.Dense{
				mat.NewDense(2, 2, []float64{0.008556624, 0.0039434, -0.00394338, 0.0164434}),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := CalculateJacobian(tc.element, tc.nodeMap, tc.points)

			for i, j := range res {
				if !mat.EqualApprox(j, tc.result[i], 1e-6) {
					t.Errorf("Expected %v, got %v", tc.result[i], j)
				}
			}
		})
	}
}
