package integration

import (
	"math"
	"testing"
)

const tolerance = 1e-6

func TestGaussIntegration(t *testing.T) {
	type tcData struct {
		name   string
		f      func(float64, float64) float64
		points int
		result float64
	}

	testCases := []tcData{
		{
			name:   "should return 16.0 for \"f(x, y) = -2*x*x*y + 2*x*y + 4\" and 2 points",
			f:      F1,
			points: 2,
			result: 16.0,
		},
		{
			name:   "should return 40.0 for \"f(x, y) = -5*x*x*y + 2*x*y + 10\" and 3 points",
			f:      F2,
			points: 3,
			result: 40.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := GaussIntegration(tc.f, tc.points)

			if math.Abs(res-tc.result) > tolerance {
				t.Errorf("Expected %f, got %f", tc.result, res)
			}
		})
	}
}
