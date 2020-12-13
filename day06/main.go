// AOC2200 Day 6
package main

import (
	"strconv"
	"zhatt/aoc2020/aoc"
)

type customsData map[string]int

type groupData struct {
	groupSize int
	groupData customsData
}

func newGroupData() groupData {
	value := groupData{}
	value.groupData = make(customsData)

	return value
}

func parseInput(input []string) []groupData {
	data := []groupData{}

	currentGroupData := newGroupData()
	for _, line := range input {
		if line == "" {
			data = append(data, currentGroupData)
			currentGroupData = newGroupData()
			continue
		}

		currentGroupData.groupSize++
		for _, value := range line {
			currentGroupData.groupData[string(value)]++
		}
	}
	data = append(data, currentGroupData)

	return data
}

func part1(inputList []string) string {
	customsDataList := parseInput(inputList)

	numYes := 0

	for _, groupsAnswers := range customsDataList {
		numYes += len(groupsAnswers.groupData)
	}
	return strconv.Itoa(numYes)
}

func part2(inputList []string) string {
	customsDataList := parseInput(inputList)
	numAllYes := 0

	for _, groupsAnswers := range customsDataList {
		for _, question := range groupsAnswers.groupData {
			if question == groupsAnswers.groupSize {
				numAllYes++
			}
		}
	}
	return strconv.Itoa(numAllYes)
}

func main() {
	aoc.MainFunc(part1, part2)
}
