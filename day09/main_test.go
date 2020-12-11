package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"35",
	"20",
	"15",
	"25",
	"47",
	"40",
	"62",
	"55",
	"65",
	"95",
	"102",
	"117",
	"150",
	"182",
	"127",
	"219",
	"299",
	"277",
	"309",
	"576",
}

func TestParseInput(t *testing.T) {
	numbers := parseInput(exampleInput)
	assert.Equal(t, 25, numbers[3])
	assert.Equal(t, 55, numbers[7])
}

func TestPart1Example1(t *testing.T) {
	numbers := parseInput(exampleInput)
	result := findBreaker1(5, numbers)
	assert.Equal(t, 127, result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "217430975", result)
}

func TestPart2Example1(t *testing.T) {
	numbers := parseInput(exampleInput)
	breaker1 := findBreaker1(5, numbers)
	result := findBreaker2(breaker1, numbers)
	assert.Equal(t, 62, result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "28509180", result)
}
