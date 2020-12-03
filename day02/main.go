package main

import (
	"fmt"
	"strconv"
	"strings"

	"zhatt/aoc2020/aoc"
)

type passwordData struct {
	value1   int
	value2   int
	letter   byte
	password string
}

func parseInput(input []string) []passwordData {
	retVal := make([]passwordData, 0, len(input))
	for _, line := range input {
		var data passwordData
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &data.value1, &data.value2,
			&data.letter, &data.password)
		aoc.PanicOnError(err)
		retVal = append(retVal, data)
	}
	return retVal
}

func countValidPasswordsOldAlgorithm(passwordData []passwordData) int {
	numberOfValidPasswords := 0
	for _, password := range passwordData {
		count := strings.Count(password.password, string(password.letter))

		if password.value1 <= count && count <= password.value2 {
			numberOfValidPasswords++
		}
	}

	return numberOfValidPasswords
}

func countValidPasswordsNewAlgorithm(passwordData []passwordData) int {
	numberOfValidPasswords := 0
	for _, password := range passwordData {
		count := 0
		if password.password[password.value1-1] == password.letter {
			count++
		}
		if password.password[password.value2-1] == password.letter {
			count++
		}

		if count == 1 {
			numberOfValidPasswords++
		}
	}

	return numberOfValidPasswords
}

func part1(inputList []string) string {
	return strconv.Itoa(countValidPasswordsOldAlgorithm(parseInput(inputList)))
}

func part2(inputList []string) string {
	return strconv.Itoa(countValidPasswordsNewAlgorithm(parseInput(inputList)))
}

func main() {
	aoc.MainFunc(part1, part2)
}
