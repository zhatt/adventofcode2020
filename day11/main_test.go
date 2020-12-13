package main

import (
	"strings"
	"testing"
	"zhatt/aoc2020/aoc"
	"zhatt/aoc2020/coord"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

func TestParseInput(t *testing.T) {
	floorPlan := parseInput(exampleInput)

	assert.Equal(t, 10, floorPlan.xSize)
	assert.Equal(t, 10, floorPlan.ySize)
	assert.Equal(t, '.', floorPlan.data[coord.Coord{Xval: 1, Yval: 2}])
}

func TestFloorPlanString(t *testing.T) {
	floorPlan := parseInput(exampleInput)

	asString := floorPlan.String()
	expected := strings.Join(exampleInput, "\n") + "\n"

	assert.Equal(t, expected, asString)
}

func TestFloorPlanClone(t *testing.T) {
	floorPlan := parseInput(exampleInput)

	newFloorPlan := floorPlan.clone()
	asString := newFloorPlan.String()
	expected := strings.Join(exampleInput, "\n") + "\n"

	assert.Equal(t, expected, asString)
}

func TestFloorCountNeighbors(t *testing.T) {
	var input = []string{
		"#.##LL",
		"####LL",
		"#.#.L.",
	}

	floorPlan := parseInput(input)
	var count int

	count = floorPlan.countNeighbors(coord.Coord{Xval: 0, Yval: 0})
	assert.Equal(t, 2, count)

	count = floorPlan.countNeighbors(coord.Coord{Xval: 3, Yval: 1})
	assert.Equal(t, 4, count)

	count = floorPlan.countNeighbors(coord.Coord{Xval: 5, Yval: 2})
	assert.Equal(t, 0, count)
}

func TestFloorCountCanSee(t *testing.T) {
	var input1 = []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}

	var input2 = []string{
		".............",
		".L.L.#.#.#.#.",
		".............",
	}

	var count int
	floorPlan1 := parseInput(input1)
	count = floorPlan1.countCanSee(coord.Coord{Xval: 3, Yval: 4})
	assert.Equal(t, 8, count)

	floorPlan2 := parseInput(input2)
	count = floorPlan2.countCanSee(coord.Coord{Xval: 1, Yval: 1})
	assert.Equal(t, 0, count)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "37", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2263", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "26", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2002", result)
}
