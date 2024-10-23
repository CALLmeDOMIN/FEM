package common

type Node struct {
	X float64
	Y float64
}

type Element struct {
	Ids    []int
	Ksi    []float64
	Eta    []float64
	DNdKsi [][]float64
	DNdEta [][]float64
}

type Grid struct {
	NodesNumber    int
	ElementsNumber int
	Nodes          []Node
	Elements       []Element
	Width          int `json:"width"`
	Height         int `json:"height"`
	NumberHeight   int `json:"numberHeight"`
	NumberWidth    int `json:"numberWidth"`
}

type GlobalData struct {
	SimulationTime     int `json:"simulationTime"`
	SimulationStepTime int `json:"simulationStepTime"`
	Conductivity       int `json:"conductivity"`
	Alfa               int `json:"alfa"`
	AmbientTemperature int `json:"ambientTemperature"`
	InitialTemperature int `json:"initialTemperature"`
	Density            int `json:"density"`
	SpecificHeat       int `json:"specificHeat"`
	NodesNumber        int `json:"nodesNumber"`
	ElementsNumber     int `json:"elementsNumber"`
}
