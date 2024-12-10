package main

import (
	"fmt"
)

func groupTemperatures(temps []float64, step int) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/float64(step)) * step
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10

	groupedTemps := groupTemperatures(temps, step)

	for key, values := range groupedTemps {
		fmt.Printf("%d: %v\n", key, values)
	}
}
