package integration

import c "mes/common"

func GaussIntegration(f func(float64, float64) float64, points int) (result float64) {
	coords := c.Points[points].Coords
	weights := c.Points[points].Weights

	for i := 0; i < points; i++ {
		for j := 0; j < points; j++ {
			result += f(coords[i], coords[j]) * weights[i] * weights[j]
		}
	}

	return
}
