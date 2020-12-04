package main

import (
	"strconv"
	"zhatt/aoc2020/aoc"
	"zhatt/aoc2020/coord"
)

type mapData struct {
	mapHeight int
	mapWidth  int
	isTree    map[coord.Coord]bool
}

func newMapData() mapData {
	mapData := mapData{}
	mapData.isTree = make(map[coord.Coord]bool)
	return mapData
}

func parseInput(input []string) mapData {
	mapData := newMapData()
	for y, line := range input {
		mapData.mapWidth = len(line)
		mapData.mapHeight = y + 1
		for x, value := range line {
			if value == '#' {
				mapData.isTree[coord.Coord{Xval: x, Yval: y}] = true
			}
		}
	}
	return mapData
}

func simulatOneSledRun(mapData mapData, slope coord.Slope) int {
	position := coord.Coord{Xval: 0, Yval: 0}
	treeCount := 0

	for position.Yval < mapData.mapHeight {
		if mapData.isTree[position] {
			treeCount++
		}
		position = coord.Coord{
			Xval: (position.Xval + slope.Run) % mapData.mapWidth,
			Yval: position.Yval + slope.Rise,
		}
	}

	return treeCount
}

func simulateSledRuns(mapData mapData, slopes []coord.Slope) int {
	numTrees := 1
	for _, slope := range slopes {
		numTrees *= simulatOneSledRun(mapData, slope)
	}
	return numTrees
}

func part1(inputList []string) string {
	mapData := parseInput(inputList)
	slopes := []coord.Slope{
		{Run: 3, Rise: 1},
	}
	return strconv.Itoa(simulateSledRuns(mapData, slopes))
}

func part2(inputList []string) string {
	mapData := parseInput(inputList)
	slopes := []coord.Slope{
		{Run: 1, Rise: 1},
		{Run: 3, Rise: 1},
		{Run: 5, Rise: 1},
		{Run: 7, Rise: 1},
		{Run: 1, Rise: 2},
	}
	return strconv.Itoa(simulateSledRuns(mapData, slopes))
}

func main() {
	aoc.MainFunc(part1, part2)
}
