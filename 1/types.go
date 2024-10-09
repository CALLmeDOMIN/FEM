package main

type Node struct {
	x float64
	y float64
}

type Element struct {
	ids []int
}

type Grid struct {
	nodesNumber    int
	elementsNumber int
	nodes          []Node
	elements       []Element
	width          int
	height         int
	numberHeight   int
	numberWidth    int
}

type GlobalData struct {
	simulationTime     int
	simulationStepTime int
	conductivity       int
	alfa               int
	ambientTemperature int
	initialTemperature int
	density            int
	specificHeat       int
	nodesNumber        int
	elementsNumber     int
}
