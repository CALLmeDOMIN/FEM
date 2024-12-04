package common

import "gonum.org/v1/gonum/mat"

type Node struct {
	ID int     `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	BC bool
}

type Geometry interface {
	GetIDs() []int
	GetDNdKsi(point int) []float64
	GetDNdEta(point int) []float64
}

type Element struct {
	ID      int   `json:"id"`
	NodeIDs []int `json:"nodes"`
	Ksi     []float64
	Eta     []float64
	DNdKsi  [][]float64
	DNdEta  [][]float64
	HMatrix *mat.Dense
}

func (e Element) GetIDs() []int {
	return e.NodeIDs
}

func (e Element) GetDNdKsi(point int) []float64 {
	return e.DNdKsi[point]
}

func (e Element) GetDNdEta(point int) []float64 {
	return e.DNdEta[point]
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
	BCNodes        []int   `json:"bcNodes"`
	HMatrix        *mat.Dense
}

type GlobalData struct {
	SimulationTime     int     `json:"simulationTime"`
	SimulationStepTime int     `json:"simulationStepTime"`
	Conductivity       float64 `json:"conductivity"`
	Alpha              float64 `json:"alpha"`
	AmbientTemperature float64 `json:"ambientTemperature"`
	InitialTemperature float64 `json:"initialTemperature"`
	Density            float64 `json:"density"`
	SpecificHeat       float64 `json:"specificHeat"`
	NodesNumber        int     `json:"nodesNumber"`
	ElementsNumber     int     `json:"elementsNumber"`
}
