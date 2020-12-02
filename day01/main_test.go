package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2020/aoc"
)

var exampleInput = []string{
	"1721",
	"979",
	"366",
	"299",
	"675",
	"1456",
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "514579", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "100419", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "241861950", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "265253940", result)
}
