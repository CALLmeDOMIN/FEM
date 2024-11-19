package common

import "gonum.org/v1/gonum/mat"

type Node struct {
	ID int     `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

type Element struct {
	ID      int   `json:"id"`
	IDs     []int `json:"nodes"`
	Ksi     []float64
	Eta     []float64
	DNdKsi  [][]float64
	DNdEta  [][]float64
	HMatrix *mat.Dense
}

type Grid struct {
	NodesNumber    int
	ElementsNumber int
	Nodes          []Node
	NodeMap        map[int]Node
	Elements       []Element
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	NumberHeight   int     `json:"numberHeight"`
	NumberWidth    int     `json:"numberWidth"`
	HMatrix        *mat.Dense
}

type GlobalData struct {
	SimulationTime     int     `json:"simulationTime"`
	SimulationStepTime int     `json:"simulationStepTime"`
	Conductivity       float64 `json:"conductivity"`
	Alpha              int     `json:"alfa"`
	AmbientTemperature float64 `json:"ambientTemperature"`
	InitialTemperature float64 `json:"initialTemperature"`
	Density            float64 `json:"density"`
	SpecificHeat       float64 `json:"specificHeat"`
	NodesNumber        int     `json:"nodesNumber"`
	ElementsNumber     int     `json:"elementsNumber"`
}
