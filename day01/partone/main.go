package main

import "fmt"
import "advent-of-code/day01/input"

func main() {
	totalFuel := 0
	for _, mass := range input.Input {
		totalFuel += (mass / 3) - 2
	}

	fmt.Println(totalFuel)
}
