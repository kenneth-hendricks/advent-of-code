package main

import "fmt"
import "adventofcode2019/day03/input"
import "strings"
import "os"
import "strconv"

func main() {
	pathOne := strings.Split(input.PathOne, ",")
	pathTwo := strings.Split(input.PathTwo, ",")

	pathOneCoordinateSteps := plotCoordinateSteps(pathOne)
	pathTwoCoordinateSteps := plotCoordinateSteps(pathTwo)

	intersectionSteps := getIntersectionCombinedSteps(pathOneCoordinateSteps, pathTwoCoordinateSteps)

	var shortestSteps = intersectionSteps[0]

	for i := 1; i < len(intersectionSteps); i++ {
		if intersectionSteps[i] < shortestSteps {
			shortestSteps = intersectionSteps[i]
		}
	}

	fmt.Println(shortestSteps)
}

func plotCoordinateSteps(path []string) map[coordinate]int {
	var coordinates = map[coordinate]int{}
	currentCoordinate := coordinate{x: 0, y: 0}
	squaredEntered := 0

	for i := 0; i < len(path); i++ {
		direction := string(path[i][0:1])
		distance, _ := strconv.Atoi(string(path[i][1:]))

		for d := 1; d <= distance; d++ {
			switch direction {
			case "U":
				currentCoordinate.y++
			case "R":
				currentCoordinate.x++
			case "D":
				currentCoordinate.y--
			case "L":
				currentCoordinate.x--
			default:
				fmt.Printf("Direction '%s' not recognized \n", direction)
				os.Exit(1)
			}

			_, coordExists := coordinates[currentCoordinate]

			squaredEntered++
			if !coordExists {
				coordinates[currentCoordinate] = squaredEntered
			}
		}
	}

	return coordinates
}

func getIntersectionCombinedSteps(coordinatesA map[coordinate]int, coordinatesB map[coordinate]int) []int {
	combinedSteps := []int{}

	for coord := range coordinatesA {
		_, coordInBoth := coordinatesB[coord]
		if coordInBoth {
			combinedSteps = append(combinedSteps, coordinatesA[coord]+coordinatesB[coord])
		}
	}

	return combinedSteps
}

type coordinate struct {
	x int
	y int
}
