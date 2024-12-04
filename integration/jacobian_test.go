package integration

import (
	"testing"

	"gonum.org/v1/gonum/mat"

	c "mes/common"
)

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
				NodeIDs: []int{1, 2, 3, 4},
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
				NodeIDs: []int{1, 2, 3, 4},
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
			res := CalculateJacobians(tc.element, tc.nodeMap, tc.points)

			for i, j := range res {
				if !mat.EqualApprox(j, tc.result[i], 1e-6) {
					t.Errorf("Expected %v, got %v", tc.result[i], j)
				}
			}
		})
	}
}
