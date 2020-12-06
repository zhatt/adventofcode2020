package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

func TestSeatID(t *testing.T) {
	assert.Equal(t, 567, seatID("BFFFBBFRRR"))
	assert.Equal(t, 119, seatID("FFFBBBFRRR"))
	assert.Equal(t, 820, seatID("BBFFBBFRLL"))
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "855", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "552", result)
}
