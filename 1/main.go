package main

import (
	"fmt"
	"os"

	c "mes/common"
)

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

			elements[i*numberWidth+j] = c.Element{Ids: ids}
		}
	}

	return elements
}

func generateNodes(width, height, numW, numH, nodesNumber int) []c.Node {
	elementHeight := height / numH
	elementWidth := width / numW

	nodes := make([]c.Node, nodesNumber)

	for i := 0; i <= numW; i++ {
		for j := 0; j <= numH; j++ {
			node := c.Node{
				X: float64(i * elementWidth),
				Y: float64(j * elementHeight),
			}

			nodes[i*(numH+1)+j] = node
		}
	}

	return nodes
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	grid, globalData, err := c.ReadFromFile(file)
	if err != nil {
		fmt.Println("Error reading from file: ", err)
		return
	}

	grid.Elements = generateElements(grid.NumberWidth, grid.NumberHeight, grid.ElementsNumber)
	grid.Nodes = generateNodes(grid.Width, grid.Height, grid.NumberWidth, grid.NumberHeight, grid.NodesNumber)

	fmt.Printf("GlobalData: %v\n", globalData)
	fmt.Printf("Grid: %v\n", grid)
}
