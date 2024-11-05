package common

type Node struct {
	Id int     `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

type Element struct {
	Id     int   `json:"id"`
	Ids    []int `json:"nodes"`
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
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	NumberHeight   int     `json:"numberHeight"`
	NumberWidth    int     `json:"numberWidth"`
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
