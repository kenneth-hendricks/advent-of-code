package main

import (
	"adventofcode2019/day06/input"
	"fmt"
	"strings"
)

func main() {
	orbits := strings.Split(input.Orbits, "\n")

	var objects = map[string]string{}

	objects["COM"] = ""

	for i := range orbits {
		parentObject := orbits[i][:3]
		childObject := orbits[i][4:]

		objects[childObject] = parentObject
	}

	orbitCount := 0

	// For each object we need to calculate distance to COM
	for child := range objects {
		currentObject := child

		for currentObject != "COM" {
			orbitCount++
			currentObject = objects[currentObject]
		}
	}

	fmt.Println(orbitCount)
}
