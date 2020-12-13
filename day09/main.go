// AOC2200 Day 9
package main

import (
	"strconv"
	"zhatt/aoc2020/aoc"
)

const standardPreambleLength = 25

func parseInput(input []string) []int {
	numbers := make([]int, len(input))

	for index, line := range input {
		number, err := strconv.Atoi(line)
		aoc.PanicOnError(err)

		numbers[index] = number
	}

	return numbers
}

func findBreaker1(preambleLength int, numbers []int) int {
	var breaker int
	for index, number := range numbers {
		if index < preambleLength {
			continue
		}

		found := false

	outer:
		for i, n1 := range numbers[index-preambleLength : index] {
			for _, n2 := range numbers[index-preambleLength+i+1 : index] {
				if n1+n2 == number {
					found = true
					break outer
				}
			}
		}

		if !found {
			breaker = number
			break
		}
	}

	return breaker
}

func findBreaker2(breakerValue int, numbers []int) int {
	for i, number1 := range numbers {
		sum := number1
		smallest := number1
		largest := number1

		for _, number2 := range numbers[i+1:] {
			if number2 < smallest {
				smallest = number2
			}
			if number2 > largest {
				largest = number2
			}
			sum += number2
			if sum == breakerValue {
				return smallest + largest
			}
		}
	}
	panic("failed")
}

func part1(inputList []string) string {
	numbers := parseInput(inputList)
	result := findBreaker1(standardPreambleLength, numbers)
	return strconv.Itoa(result)
}

func part2(inputList []string) string {
	numbers := parseInput(inputList)
	breaker1 := findBreaker1(standardPreambleLength, numbers)
	result := findBreaker2(breaker1, numbers)
	return strconv.Itoa(result)
}

func main() {
	aoc.MainFunc(part1, part2)
}
