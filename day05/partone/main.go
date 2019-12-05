package main

import "fmt"
import "strings"
import "strconv"
import "adventofcode2019/day05/input"
import "os"

func main() {
	input := strings.Split(input.Program, ",")
	intCodes := make([]int, len(input))
	for i := range input {
		intCodes[i], _ = strconv.Atoi(input[i])
	}

	programInput := 1

	opCodePosition := 0

	pm1, pm2, opCode := parseInstruction(intCodes[0])
	for opCode != 99 {
		switch opCode {
		case 1:
			intCodes[intCodes[opCodePosition+3]] = getParam(opCodePosition+1, pm1, intCodes) + getParam(opCodePosition+2, pm2, intCodes)
			opCodePosition += 4

		case 2:
			intCodes[intCodes[opCodePosition+3]] = getParam(opCodePosition+1, pm1, intCodes) * getParam(opCodePosition+2, pm2, intCodes)
			opCodePosition += 4

		case 3:
			intCodes[intCodes[opCodePosition+1]] = programInput
			opCodePosition += 2

		case 4:
			fmt.Println(getParam(opCodePosition+1, pm1, intCodes))
			opCodePosition += 2

		default:
			fmt.Printf("Op code '%d' not recognized \n", opCode)
			os.Exit(1)
		}

		pm1, pm2, opCode = parseInstruction(intCodes[opCodePosition])
	}
}

func getParam(index int, mode int, intCodes []int) int {
	if mode == 0 {
		return intCodes[intCodes[index]]
	}

	return intCodes[index]
}

func parseInstruction(intInst int) (int, int, int) {
	instruction := strconv.Itoa(intInst)

	numLeadingZeros := 5 - len(instruction)

	for i := 0; i < numLeadingZeros; i++ {
		instruction = string('0') + instruction
	}

	// fmt.Println(instruction)

	// pm3, _ := strconv.Atoi(string(instruction[0]))
	pm2, _ := strconv.Atoi(string(instruction[1]))
	pm1, _ := strconv.Atoi(string(instruction[2]))
	opCode, _ := strconv.Atoi(string(instruction[3:]))

	return pm1, pm2, opCode
}
