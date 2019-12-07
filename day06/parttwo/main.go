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

	var youPath = map[string]int{}

	currentObjectKey := "YOU"
	distance := 0

	for currentObjectKey != "COM" {
		youPath[currentObjectKey] = distance
		distance++
		currentObjectKey = objects[currentObjectKey]
	}

	foundIntersection := false

	currentObjectKey = "SAN"
	sanDistance := 0
	youDistance := 0

	for !foundIntersection {
		if dist, ok := youPath[currentObjectKey]; ok {
			foundIntersection = true
			youDistance = dist
			break
		}

		sanDistance++
		currentObjectKey = objects[currentObjectKey]
	}

	fmt.Println(youDistance + sanDistance - 2)
}

func calcPath(objects map[string]string, objectKey string) map[string]int {
	var path = map[string]int{}

	currentObjectKey := objectKey
	distance := 0

	for currentObjectKey != "COM" {
		path[currentObjectKey] = distance
		distance++
		currentObjectKey = objects[currentObjectKey]
	}

	return path
}
