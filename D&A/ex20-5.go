package main

import (
	"fmt"
)

func main() {
	arr := []float32{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}
	value := sortTemperatures(arr) //2008
	fmt.Println(value)
}

func sortTemperatures(arr []float32) []float32 {
	hash := make(map[float32]int)
	sorted := []float32{}

	for _, num := range arr {

		if hash[num] == 0 {
			hash[num] = 1
		} else {
			hash[num]++
		}
	}

	for i := 970; i <= 990; i += 1 {

		fl := float32(i) / 10.0
		if hash[fl] != 0 {
			for y := 0; y < hash[fl]; y++ {
				sorted = append(sorted, fl)
			}
		}
	}
	return sorted
}
