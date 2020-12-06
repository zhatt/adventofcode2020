package main

import (
	"sort"
	"strconv"
	"strings"
	"zhatt/aoc2020/aoc"
)

func seatID(locator string) int {
	locator = strings.ReplaceAll(locator, "F", "0")
	locator = strings.ReplaceAll(locator, "B", "1")
	locator = strings.ReplaceAll(locator, "L", "0")
	locator = strings.ReplaceAll(locator, "R", "1")

	row, err := strconv.ParseInt(locator[0:7], 2, 0)
	aoc.PanicOnError(err)
	col, err := strconv.ParseInt(locator[7:10], 2, 0)
	aoc.PanicOnError(err)

	seatID := row*8 + col

	return int(seatID)
}

func part1(inputList []string) string {
	highestSeatID := 0

	for _, locator := range inputList {
		seatID := seatID(locator)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	return strconv.Itoa(highestSeatID)
}

func part2(inputList []string) string {
	seatIDList := []int{}

	for _, locator := range inputList {
		seatID := seatID(locator)
		seatIDList = append(seatIDList, seatID)
	}

	sort.Ints(seatIDList)

	firstSeatID := seatIDList[0]
	mySeatID := 0

	for index, s := range seatIDList {
		if index+firstSeatID != s {
			mySeatID = index + firstSeatID
			break
		}
	}

	return strconv.Itoa(mySeatID)
}

func main() {
	aoc.MainFunc(part1, part2)
}
