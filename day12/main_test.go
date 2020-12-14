package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"F10",
	"N3",
	"F7",
	"R90",
	"F11",
}

func TestParseInput(t *testing.T) {
	commands := parseInput(exampleInput)
	assert.Equal(t, actionType{command: cmdForward, amount: 7}, commands[2])
	assert.Equal(t, actionType{command: cmdRight, amount: 90}, commands[3])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "25", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "1177", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "286", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "46530", result)
}
