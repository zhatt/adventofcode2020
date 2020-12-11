package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestParseInput(t *testing.T) {
	instructions := parseInput(exampleInput)
	assert.Equal(t, 9, len(instructions))
	assert.Equal(t, opcodeType("acc"), instructions[3].opcode)
	assert.Equal(t, 3, instructions[3].operand)
	assert.Equal(t, opcodeType("jmp"), instructions[7].opcode)
	assert.Equal(t, -4, instructions[7].operand)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "5", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "1384", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "8", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "761", result)
}
