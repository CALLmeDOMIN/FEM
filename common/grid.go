package common

func GenerateElements(numberWidth, numberHeight, elementsNumber, points int) []Element {
	elements := make([]Element, elementsNumber)

	ksi := make([]float64, points)
	eta := make([]float64, points)

	for i := 0; i < points; i++ {
		for j := 0; j < points; j++ {
			ksi[i] = Points[points].Coords[i]
			eta[i] = Points[points].Coords[i]
		}
	}

	dNdKsi := make([][]float64, points)
	dNdEta := make([][]float64, points)

	for i := 0; i < len(ksi); i++ {
		ksiValue := ksi[i]
		etaValue := eta[i]

		dNdKsi[i] = []float64{
			-0.25 * (1 - etaValue),
			0.25 * (1 - etaValue),
			0.25 * (1 + etaValue),
			-0.25 * (1 + etaValue),
		}

		dNdEta[i] = []float64{
			-0.25 * (1 - ksiValue),
			-0.25 * (1 + ksiValue),
			0.25 * (1 + ksiValue),
			0.25 * (1 - ksiValue),
		}
	}

	for i := 0; i < numberHeight; i++ {
		for j := 0; j < numberWidth; j++ {
			ids := []int{
				i*(numberWidth+1) + j + 1,
				i*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 1,
			}

			elements[i*numberWidth+j] = Element{
				Ids:    ids,
				Ksi:    ksi,
				Eta:    eta,
				DNdKsi: dNdKsi,
				DNdEta: dNdEta,
			}
		}
	}

	return elements
}

func GenerateNodes(width, height, numW, numH, nodesNumber int) []Node {
	elementHeight := height / numH
	elementWidth := width / numW

	nodes := make([]Node, nodesNumber)

	for i := 0; i <= numW; i++ {
		for j := 0; j <= numH; j++ {
			node := Node{
				X: float64(i * elementWidth),
				Y: float64(j * elementHeight),
			}

			nodes[i*(numH+1)+j] = node
		}
	}

	return nodes
}
