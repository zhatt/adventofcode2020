package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput)

	assert.Equal(t, 5, len(data))

	two := data[2]
	assert.Equal(t, 3, len(two.groupData))
	assert.Equal(t, groupData{
		groupSize: 2,
		groupData: customsData{"a": 2, "b": 1, "c": 1},
	},
		two)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "11", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "6662", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "6", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "3382", result)
}
