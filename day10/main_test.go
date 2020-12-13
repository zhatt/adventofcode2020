package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"16",
	"10",
	"15",
	"5",
	"1",
	"11",
	"7",
	"19",
	"6",
	"12",
	"4",
}

var exampleInput2 = []string{
	"28",
	"33",
	"18",
	"42",
	"31",
	"14",
	"46",
	"20",
	"48",
	"47",
	"24",
	"23",
	"49",
	"45",
	"19",
	"38",
	"39",
	"11",
	"1",
	"32",
	"25",
	"35",
	"8",
	"17",
	"7",
	"9",
	"4",
	"2",
	"34",
	"10",
	"3",
}

func TestParseInput(t *testing.T) {
	numbers := parseInput(exampleInput1)
	assert.Equal(t, joltageValue(15), numbers[2])
	assert.Equal(t, joltageValue(19), numbers[7])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "35", result)
}

func TestPart1Example2(t *testing.T) {
	result := part1(exampleInput2)
	assert.Equal(t, "220", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2346", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "8", result)
}

func TestPart2Example2(t *testing.T) {
	result := part2(exampleInput2)
	assert.Equal(t, "19208", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "6044831973376", result)
}
