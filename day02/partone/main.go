package main

import "fmt"
import "strings"
import "strconv"
import "adventofcode2019/day02/input"
import "os"

func main() {
	input := strings.Split(input.Input, ",")
	intCodes := make([]int, len(input))
	for i := range input {
		intCodes[i], _ = strconv.Atoi(input[i])
	}

	intCodes[1] = 12
	intCodes[2] = 2

	opCodePosition := 0

	for intCodes[opCodePosition] != 99 {
		var newValue int
		firstValue := intCodes[intCodes[opCodePosition+1]]
		secondValue := intCodes[intCodes[opCodePosition+2]]

		switch opCode := intCodes[opCodePosition]; opCode {
		case 1:
			newValue = firstValue + secondValue
		case 2:
			newValue = firstValue * secondValue
		default:
			fmt.Printf("Op code '%d' not recognized \n", opCode)
			os.Exit(1)
		}

		intCodes[intCodes[opCodePosition+3]] = newValue

		opCodePosition += 4
	}

	output := make([]string, len(intCodes))

	for i := range intCodes {
		output[i] = strconv.Itoa(intCodes[i])
	}

	fmt.Println(strings.Join(output, ","))
}
