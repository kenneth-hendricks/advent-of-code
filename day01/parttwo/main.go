package main

import "fmt"
import "advent-of-code/day01/input"

func main() {
	totalFuel := 0
	for _, mass := range input.Input {
		totalFuel += getModuleFuel(mass)
	}

	fmt.Println(totalFuel)
}

func getModuleFuel(mass int) int {
	fuel := (mass / 3) - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + getModuleFuel(fuel)
}
