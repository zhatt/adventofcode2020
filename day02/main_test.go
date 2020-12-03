package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func TestParseInput(t *testing.T) {
	expected := []passwordData{
		{1, 3, 'a', "abcde"},
		{1, 3, 'b', "cdefg"},
		{2, 9, 'c', "ccccccccc"},
	}
	passwordData := parseInput(exampleInput)
	assert.Equal(t, expected, passwordData)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "2", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "418", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "1", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "616", result)
}
