package main

import (
	"testing"
	"zhatt/aoc2020/coord"

	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestParseInput(t *testing.T) {
	input := []string{
		".##",
		"#..",
	}
	mapData := parseInput(input)
	expected := newMapData()
	expected.mapWidth = 3
	expected.mapHeight = 2
	expected.isTree[coord.Coord{Xval: 1, Yval: 0}] = true
	expected.isTree[coord.Coord{Xval: 2, Yval: 0}] = true
	expected.isTree[coord.Coord{Xval: 0, Yval: 1}] = true
	assert.Equal(t, expected, mapData)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "7", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "220", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "336", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2138320800", result)
}
