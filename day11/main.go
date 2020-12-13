// AOC2200 Day 11
package main

import (
	"reflect"
	"strconv"
	"strings"
	"zhatt/aoc2020/aoc"
	"zhatt/aoc2020/coord"
)

type floorPlanType struct {
	xSize int
	ySize int
	data  map[coord.Coord]rune
}

func (floorplan *floorPlanType) clone() *floorPlanType {
	clone := &floorPlanType{}
	clone.xSize = floorplan.xSize
	clone.ySize = floorplan.ySize
	clone.data = map[coord.Coord]rune{}
	for k, v := range floorplan.data {
		clone.data[k] = v
	}
	return clone
}

func (floorplan *floorPlanType) load(input []string) {
	floorplan.xSize = len(input[0])
	floorplan.ySize = len(input)
	floorplan.data = map[coord.Coord]rune{}

	for yval, line := range input {
		for xval, symbol := range line {
			coordXY := coord.Coord{Xval: xval, Yval: yval}
			floorplan.data[coordXY] = symbol
		}
	}
}

func (floorplan *floorPlanType) String() string {
	var buffer strings.Builder

	for y := 0; y < floorplan.ySize; y++ {
		for x := 0; x < floorplan.xSize; x++ {
			coordXY := coord.Coord{Xval: x, Yval: y}
			buffer.WriteRune(floorplan.data[coordXY])
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

var neighborIncrements = []coord.Coord{
	{Xval: 0, Yval: 1},
	{Xval: 1, Yval: 1},
	{Xval: 1, Yval: 0},
	{Xval: 1, Yval: -1},
	{Xval: 0, Yval: -1},
	{Xval: -1, Yval: -1},
	{Xval: -1, Yval: 0},
	{Xval: -1, Yval: 1},
}

func (floorplan *floorPlanType) countNeighbors(seatCoord coord.Coord) int {
	count := 0
	for _, increment := range neighborIncrements {
		adjacentCoord := coord.Add(seatCoord, increment)

		if floorplan.data[adjacentCoord] == '#' {
			count++
		}
	}
	return count
}

func (floorplan *floorPlanType) countCanSee(seatCoord coord.Coord) int {
	count := 0

	for _, increment := range neighborIncrements {
		adjacentCoord := seatCoord
		for {
			adjacentCoord = coord.Add(adjacentCoord, increment)
			value, exists := floorplan.data[adjacentCoord]

			if !exists {
				break
			}

			if value == 'L' {
				break
			}

			if value == '#' {
				count++
				break
			}
		}
	}

	return count
}

func simulateSeating(floorplan *floorPlanType,
	neighborCounter func(seatCoord coord.Coord) int,
	neighborThreshold int) int {
	for {
		newFloorplan := floorplan.clone()

		for y := 0; y < floorplan.ySize; y++ {
			for x := 0; x < floorplan.xSize; x++ {
				coordXY := coord.Coord{Xval: x, Yval: y}
				current := floorplan.data[coordXY]

				if current == '.' {
					continue
				}

				numberOfNeighbors := neighborCounter(coordXY)
				if current == 'L' {
					if numberOfNeighbors == 0 {
						newFloorplan.data[coordXY] = '#'
					}
				} else if current == '#' {
					if numberOfNeighbors >= neighborThreshold {
						newFloorplan.data[coordXY] = 'L'
					}
				}
			}
		}

		if reflect.DeepEqual(floorplan, newFloorplan) {
			break
		}

		*floorplan = *newFloorplan
	}

	count := 0
	for y := 0; y < floorplan.ySize; y++ {
		for x := 0; x < floorplan.xSize; x++ {
			coordXY := coord.Coord{Xval: x, Yval: y}
			if floorplan.data[coordXY] == '#' {
				count++
			}
		}
	}
	return count
}

func simulateSeating1(floorplan *floorPlanType) int {
	return simulateSeating(
		floorplan,
		floorplan.countNeighbors,
		4,
	)
}

func simulateSeating2(floorplan *floorPlanType) int {
	return simulateSeating(
		floorplan,
		floorplan.countCanSee,
		5,
	)
}

func parseInput(input []string) *floorPlanType {
	floorPlan := &floorPlanType{}
	floorPlan.load(input)

	return floorPlan
}

func part1(inputList []string) string {
	floorPlan := parseInput(inputList)
	count := simulateSeating1(floorPlan)
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	floorPlan := parseInput(inputList)
	count := simulateSeating2(floorPlan)
	return strconv.Itoa(count)
}

func main() {
	aoc.MainFunc(part1, part2)
}
