package main

import "fmt"
import "strings"
import "strconv"
import "advent-of-code/day02/input"
import "os"

func main() {
	input := strings.Split(input.Input, ",")
	intCodes := make([]int, len(input))
	for i := range input {
		intCodes[i], _ = strconv.Atoi(input[i])
	}

	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			// Need to make a copy of intcodes because its a slice
			tempCodes := make([]int, len(intCodes))
			copy(tempCodes, intCodes)

			if runProgram(tempCodes, n, v) == 19690720 {
				fmt.Println((100 * n) + v)
				os.Exit(0)
				break
			}
		}
	}

}

func runProgram(intCodes []int, noun int, verb int) int {
	intCodes[1] = noun
	intCodes[2] = verb

	insructionPtr := 0

	for intCodes[insructionPtr] != 99 {
		var newValue int
		firstValue := intCodes[intCodes[insructionPtr+1]]
		secondValue := intCodes[intCodes[insructionPtr+2]]

		switch opCode := intCodes[insructionPtr]; opCode {
		case 1:
			newValue = firstValue + secondValue
		case 2:
			newValue = firstValue * secondValue
		default:
			fmt.Printf("Op code '%d' not recognized \n", opCode)
			os.Exit(1)
		}

		intCodes[intCodes[insructionPtr+3]] = newValue

		insructionPtr += 4
	}

	return intCodes[0]
}
