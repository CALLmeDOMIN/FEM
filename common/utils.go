package common

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFromFile(file *os.File) (Grid, GlobalData, error) {
	var grid Grid
	var globalData GlobalData

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&struct {
		*Grid
		*GlobalData
	}{
		&grid,
		&globalData,
	})
	if err != nil {
		return Grid{}, GlobalData{}, fmt.Errorf("error decoding file: %v", err)
	}

	grid.NodesNumber = globalData.NodesNumber
	grid.ElementsNumber = globalData.ElementsNumber

	return grid, globalData, nil
}
