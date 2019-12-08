package main

import (
	"adventofcode2019/day08/input"
	"fmt"
)

func main() {
	layerLength := 25 * 6
	layerDigits := input.Input

	layerStart := 0

	var lowestZeroDigits int
	var bestScore int

	for layerStart < len(layerDigits) {
		numZeroDigits := 0
		numOneDigits := 0
		numTwoDigits := 0

		for i := layerStart; i < layerStart+layerLength; i++ {
			if string(layerDigits[i]) == "0" {
				numZeroDigits++
			}

			if string(layerDigits[i]) == "1" {
				numOneDigits++
			}

			if string(layerDigits[i]) == "2" {
				numTwoDigits++
			}

		}

		if layerStart == 0 {
			lowestZeroDigits = numZeroDigits
		}

		if numZeroDigits <= lowestZeroDigits {
			lowestZeroDigits = numZeroDigits
			bestScore = numOneDigits * numTwoDigits
		}

		layerStart += layerLength
	}

	fmt.Println(bestScore)
}
