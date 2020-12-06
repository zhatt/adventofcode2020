// Package main for AOC 2020 day 1
//
package main

import (
	"errors"
	"strconv"
	"zhatt/aoc2020/aoc"
)

func findExpenseReportError(numList []int, targetSum int) (int, error) {
	for {
		a := numList[0]
		for _, b := range numList[1:] {
			if a+b == targetSum {
				return a * b, nil
			}
		}
		numList = numList[1:]
		if len(numList) == 0 {
			return 0, errors.New("no error found")
		}
	}
}

func findExpenseReportError1(numList []int) int {
	retVal, err := findExpenseReportError(numList, 2020)
	aoc.PanicOnError(err)
	return retVal
}

func findExpenseReportError2(numList []int) int {
	for {
		a := numList[0]
		result, err := findExpenseReportError(numList[1:], 2020-a)
		if err == nil {
			return result * a
		}
		numList = numList[1:]
	}
}

func inputToIntList(inputList []string) []int {
	retVal := make([]int, 0, len(inputList))

	for _, strVal := range inputList {
		intVal, err := strconv.Atoi(strVal)
		aoc.PanicOnError(err)

		retVal = append(retVal, intVal)
	}
	return retVal
}

func part1(inputList []string) string {
	retVal := findExpenseReportError1(inputToIntList(inputList))
	return strconv.Itoa(retVal)
}

func part2(inputList []string) string {
	retVal := findExpenseReportError2(inputToIntList(inputList))
	return strconv.Itoa(retVal)
}

func main() {
	aoc.MainFunc(part1, part2)
}
