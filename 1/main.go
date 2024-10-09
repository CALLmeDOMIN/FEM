package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFromFile(file *os.File) (grid Grid, globalData GlobalData) {
	scanner := bufio.NewScanner(file)

	fields := []*int{
		&globalData.simulationTime,
		&globalData.simulationStepTime,
		&globalData.conductivity,
		&globalData.alfa,
		&globalData.ambientTemperature,
		&globalData.initialTemperature,
		&globalData.density,
		&globalData.specificHeat,
		&globalData.nodesNumber,
		&globalData.elementsNumber,
		&grid.height,
		&grid.width,
		&grid.numberHeight,
		&grid.numberWidth,
	}

	i := 0
	for scanner.Scan() {
		if i >= len(fields) {
			fmt.Printf("Error: more lines in file than expected")
			return Grid{}, GlobalData{}
		}

		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting line to int: %v", err)
			return Grid{}, GlobalData{}
		}

		*fields[i] = val
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	grid.elementsNumber = globalData.elementsNumber
	grid.nodesNumber = globalData.nodesNumber

	return grid, globalData
}

func generateElements(numberWidth, numberHeight, elementsNumber int) []Element {
	elements := make([]Element, elementsNumber)

	for i := 0; i < numberHeight; i++ {
		for j := 0; j < numberWidth; j++ {
			ids := []int{
				i*(numberWidth+1) + j + 1,
				i*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 2,
				(i+1)*(numberWidth+1) + j + 1,
			}
			fmt.Printf("ids: %v\n", ids)

			elements[i*numberWidth+j] = Element{ids}
		}
	}

	return elements
}

func generateNodes(width, height, numW, numH, nodesNumber int) []Node {
	elementHeight := height / numH
	elementWidth := width / numW

	nodes := make([]Node, nodesNumber)

	for i := 0; i <= numW; i++ {
		for j := 0; j <= numH; j++ {
			node := Node{
				float64(i * elementWidth),
				float64(j * elementHeight),
			}

			fmt.Printf("i: %v\n", i)
			fmt.Printf("j: %v\n", j)
			fmt.Printf("Node: %v\n", node)

			nodes[i*(numH+1)+j] = node
		}
	}

	return nodes
}

func main() {
	file, err := os.Open("task.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	grid, globalData := readFromFile(file)
	grid.elements = generateElements(grid.numberWidth, grid.numberHeight, grid.elementsNumber)
	grid.nodes = generateNodes(grid.width, grid.height, grid.numberWidth, grid.numberHeight, grid.nodesNumber)

	fmt.Printf("Grid: %v\n", grid)
	fmt.Printf("GlobalData: %v\n", globalData)
}
