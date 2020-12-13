// AOC2200 Day 10
package main

import (
	"sort"
	"strconv"
	"zhatt/aoc2020/aoc"
)

type joltageValue int

const maxJoltageDifference = 3

func parseInput(input []string) []joltageValue {
	numbers := make([]joltageValue, len(input))

	for index, line := range input {
		number, err := strconv.Atoi(line)
		aoc.PanicOnError(err)

		numbers[index] = joltageValue(number)
	}

	return numbers
}

func findJoltageUsingAll(adapters []joltageValue) int {
	sortedRatings := make([]joltageValue, len(adapters))
	copy(sortedRatings, adapters)

	sort.Slice(sortedRatings, func(i, j int) bool {
		return sortedRatings[i] < sortedRatings[j]
	})

	num1jolt := 0
	num3jolt := 1 // To my adapter

	joltage1 := joltageValue(0)
	for _, joltage2 := range sortedRatings {
		switch joltage2 - joltage1 {
		case 1:
			num1jolt++

		case 3:
			num3jolt++
		}
		joltage1 = joltage2
	}

	return num1jolt * num3jolt
}

type compatibleJoltageList []joltageValue
type adapterData map[joltageValue]compatibleJoltageList

type pathCache map[joltageValue]int

// countPaths performs a depth first search for all paths.
func countPaths(adapters adapterData, start, end joltageValue, cache pathCache) int {
	// Initialize the cache if needed.  We initialize the cache with
	// a value for the end indicating there is one path from it to the end.
	if cache == nil {
		cache = pathCache{end: 1}
	}

	// Use the cached value of the number of paths from hear to end.
	if cachedValue, exists := cache[start]; exists {
		return cachedValue
	}

	// Perform depth first search for all paths.  Calculate patCount
	// which is the number of paths from here to end.
	pathCount := 0
	for _, neighbor := range adapters[start] {
		pathCount += countPaths(adapters, neighbor, end, cache)
	}

	cache[start] = pathCount
	return pathCount
}

func buildAdapterGraph(adapterList []joltageValue) (adapterData, joltageValue) {
	graphData := adapterData{}

	// Initialize the list with a 0 for the outlet.
	sortedRatings := make([]joltageValue, 1, len(adapterList)+2)

	// Add the rest of the values and sort
	sortedRatings = append(sortedRatings, adapterList...)
	sort.Slice(sortedRatings, func(i, j int) bool {
		return sortedRatings[i] < sortedRatings[j]
	})

	myDeviceJoltage := sortedRatings[len(sortedRatings)-1] + 3
	sortedRatings = append(sortedRatings, myDeviceJoltage)

	for index, joltage1 := range sortedRatings {
		canConnectList := compatibleJoltageList{}

		if index != len(sortedRatings)-1 {
			for _, joltage2 := range sortedRatings[index+1:] {
				if joltage2-joltage1 <= maxJoltageDifference {
					canConnectList = append(canConnectList, joltage2)
				} else {
					break
				}
			}
		}
		graphData[joltage1] = canConnectList
	}

	return graphData, myDeviceJoltage
}

func part1(inputList []string) string {
	numbers := parseInput(inputList)
	return strconv.Itoa(findJoltageUsingAll(numbers))
}

func part2(inputList []string) string {
	numbers := parseInput(inputList)
	graph, myDeviceJoltage := buildAdapterGraph(numbers)
	count := countPaths(graph, 0, myDeviceJoltage, nil)
	return strconv.Itoa(count)
}

func main() {
	aoc.MainFunc(part1, part2)
}
