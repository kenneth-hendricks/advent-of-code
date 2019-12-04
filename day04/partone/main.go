package main

import "fmt"
import "adventofcode2019/day04/input"
import "strconv"

func main() {
	start := input.Start
	end := input.End

	validCount := 0

	for i := start; i <= end; i++ {
		if isValidPassword(strconv.Itoa(i)) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func isValidPassword(pass string) bool {
	lastInt := -1
	currentDigitCount := 1
	foundDouble := false

	for i := 0; i < len(pass); i++ {
		currentInt, _ := strconv.Atoi(string(pass[i]))

		if currentInt < lastInt {
			return false
		}

		if currentInt == lastInt {
			currentDigitCount++
		} else {
			if currentDigitCount == 2 {
				foundDouble = true
			}
			currentDigitCount = 1
		}

		lastInt = currentInt
	}

	if currentDigitCount == 2 {
		foundDouble = true
	}

	return foundDouble
}
