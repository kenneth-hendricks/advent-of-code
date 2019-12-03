package main

import "fmt"
import "adventofcode2019/day01/input"

func main() {
	totalFuel := 0
	for _, mass := range input.Input {
		totalFuel += (mass / 3) - 2
	}

	fmt.Println(totalFuel)
}
