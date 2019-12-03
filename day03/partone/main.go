package main

import "fmt"
import "adventofcode2019/day03/input"
import "strings"
import "os"
import "strconv"

func main() {
	pathOne := strings.Split(input.PathOne, ",")
	pathTwo := strings.Split(input.PathTwo, ",")

	pathOneCoordinates := plotCoordinates(pathOne)
	pathTwoCoordinates := plotCoordinates(pathTwo)

	intersectionDistances := getIntersectionDistances(pathOneCoordinates, pathTwoCoordinates)

	var shortestDistance = intersectionDistances[0]

	for i := 0; i < len(intersectionDistances); i++ {
		if intersectionDistances[i] < shortestDistance {
			shortestDistance = intersectionDistances[i]
		}
	}

	fmt.Println(shortestDistance)
}

func plotCoordinates(path []string) map[coordinate]bool {
	var coordinates = map[coordinate]bool{}
	currentCoordinate := coordinate{x: 0, y: 0}

	for i := 0; i < len(path); i++ {
		instructionRune := []rune(path[i])
		direction := string(instructionRune[0:1])
		distance, _ := strconv.Atoi(string(instructionRune[1:]))

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

			coordinates[currentCoordinate] = true
		}
	}

	return coordinates
}

func getIntersectionDistances(coordinatesA map[coordinate]bool, coordinatesB map[coordinate]bool) []int {
	intDistances := []int{}

	for coord := range coordinatesA {
		_, coordInBoth := coordinatesB[coord]
		if coordInBoth {
			intDistances = append(intDistances, absInt(coord.x)+absInt(coord.y))
		}

	}

	return intDistances
}

type coordinate struct {
	x int
	y int
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
