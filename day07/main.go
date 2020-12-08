// AOC2200 Day 4
package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2020/aoc"
)

type colorAndCount struct {
	color string
	count int
}

type bagData map[string][]colorAndCount

func parseInput(input []string) bagData {
	data := bagData{}

	for _, line := range input {
		line = strings.ReplaceAll(line, "bags.", "bag")
		line = strings.ReplaceAll(line, "bags", "bag")
		tokens := strings.Split(line, " bag contain ")
		thisBag := tokens[0]
		bagList := tokens[1]

		bagsInside := []colorAndCount{}
		if bagList != "no other bag" {
			for _, bagInfo := range strings.Split(bagList, ", ") {
				fields := strings.Fields(bagInfo)
				bagColor := fields[1] + " " + fields[2]
				bagCount, err := strconv.Atoi(fields[0])
				aoc.PanicOnError(err)
				cAndC := colorAndCount{color: bagColor, count: bagCount}
				bagsInside = append(bagsInside, cAndC)
			}
		}

		data[thisBag] = bagsInside
	}

	return data
}

func containsAColor(colorToCheckFor string, currentColor string, bagData bagData) bool {
	for _, bag := range bagData[currentColor] {
		if bag.color == colorToCheckFor {
			return true
		}

		if containsAColor(colorToCheckFor, bag.color, bagData) {
			return true
		}
	}

	return false
}

func numberOfBagsInside(bagColor string, bagData bagData) int {
	numBags := 0

	for _, bag := range bagData[bagColor] {
		// This bag
		numBags += bag.count

		// Bags inside this bag
		inside := numberOfBagsInside(bag.color, bagData) * bag.count
		numBags += inside
	}

	return numBags
}

func part1(inputList []string) string {
	bagData := parseInput(inputList)
	numHolders := 0

	for bagColor := range bagData {
		if bagColor == "shiny gold" {
			continue
		}

		contains := containsAColor("shiny gold", bagColor, bagData)
		if contains {
			numHolders++
		}
	}

	return strconv.Itoa(numHolders)
}

func part2(inputList []string) string {
	bagData := parseInput(inputList)
	numBags := numberOfBagsInside("shiny gold", bagData)
	return strconv.Itoa(numBags)
}

func main() {
	aoc.MainFunc(part1, part2)
}
