package grid

import (
	"slices"

	c "mes/common"
)

func GenerateGrid(globalData c.GlobalData, grid c.Grid, integrationPoints int) (c.Grid, c.GlobalData) {
	grid.NodesNumber = globalData.NodesNumber
	grid.ElementsNumber = globalData.ElementsNumber

	if len(grid.Nodes) == 0 {
		grid.Nodes = generateNodes(grid.Width, grid.Height, grid.NumberWidth, grid.NumberHeight, grid.NodesNumber)
	}

	for i := range grid.Nodes {
		if slices.Contains(grid.BCNodes, grid.Nodes[i].ID) {
			grid.Nodes[i].BC = true
		}
	}

	nodeMap := make(map[int]c.Node)
	for _, node := range grid.Nodes {
		nodeMap[node.ID] = node
	}
	grid.NodeMap = nodeMap

	if len(grid.Elements) == 0 {
		grid.Elements = generateElements(grid.NumberWidth, grid.NumberHeight, grid.ElementsNumber)
	}

	grid.Elements = generateShapeFunctionData(grid.Elements, grid.NumberWidth, grid.NumberHeight, integrationPoints)

	for i, element := range grid.Elements {
		grid.Elements[i].HMatrix = calculateHMatrixLocal(element, nodeMap, globalData.Conductivity, globalData.Alpha, integrationPoints)
		grid.Elements[i].PVector = calculatePVectorLocal(element, nodeMap, globalData.Alpha, globalData.AmbientTemperature, integrationPoints)
	}

	grid.HMatrix = calculateHMatrixGlobal(grid)
	grid.PVector = calculatePVectorGlobal(grid)

	return grid, globalData
}

func generateNodes(width, height float64, numW, numH, nodesNumber int) []c.Node {
	elementHeight := height / float64(numH)
	elementWidth := width / float64(numW)

	nodes := make([]c.Node, nodesNumber)

	for i := 0; i <= numW; i++ {
		for j := 0; j <= numH; j++ {
			node := c.Node{
				X: float64(i) * elementWidth,
				Y: float64(j) * elementHeight,
			}

			nodes[i*(numH+1)+j] = node
		}
	}

	return nodes
}

func generateElements(numberWidth, numberHeight, elementsNumber int) []c.Element {
	elements := make([]c.Element, elementsNumber)

	for i := 0; i < numberHeight; i++ {
		for j := 0; j < numberWidth; j++ {
			ids := []int{
				i*(numberWidth+1) + j + 1,
				i*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 1,
			}

			elements[i*numberWidth+j] = c.Element{
				NodeIDs: ids,
			}
		}
	}

	return elements
}

func generateShapeFunctionData(elements []c.Element, numberWidth, numberHeight, points int) []c.Element {
	ksi := make([]float64, points*points)
	eta := make([]float64, points*points)

	for i := 0; i < points; i++ {
		for j := 0; j < points; j++ {
			index := i*points + j
			ksi[index] = c.Points[points].Coords[j]
			eta[index] = c.Points[points].Coords[i]
		}
	}

	dNdKsi := make([][]float64, points*points)
	dNdEta := make([][]float64, points*points)

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
			IDs_copy := elements[i*numberWidth+j].NodeIDs

			elements[i*numberWidth+j] = c.Element{
				ID:      i*numberWidth + j + 1,
				NodeIDs: IDs_copy,
				Ksi:     ksi,
				Eta:     eta,
				DNdKsi:  dNdKsi,
				DNdEta:  dNdEta,
			}
		}
	}

	return elements
}
