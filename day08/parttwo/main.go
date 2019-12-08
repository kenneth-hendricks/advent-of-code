package main

import (
	"adventofcode2019/day08/input"
	"fmt"
)

func main() {
	layerLength := 25 * 6
	layerDigits := input.Input

	layerStart := 0
	message := make([]string, layerLength)
	for i := range message {
		message[i] = "2"
	}

	for layerStart < len(layerDigits) {
		for i := layerStart; i < layerStart+layerLength; i++ {
			if message[i-layerStart] == "2" && string(layerDigits[i]) == "0" {
				message[i-layerStart] = "0"
			}

			if message[i-layerStart] == "2" && string(layerDigits[i]) == "1" {
				message[i-layerStart] = "1"
			}
		}

		layerStart += layerLength
	}

	for i := 0; i < layerLength; i++ {
		if i%25 == 0 {
			fmt.Println()
		}

		if message[i] == "1" {
			fmt.Print("X")
		} else {
			fmt.Print(" ")

		}
	}

	fmt.Println()

}
