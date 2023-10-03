package main

import "fmt"

func main() {
	fmt.Println(convertTemperature(36.50))
}

func convertTemperature(celsius float64) []float64 {

	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
