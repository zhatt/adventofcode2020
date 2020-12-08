package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

var exampleInput2 = []string{
	"shiny gold bags contain 2 dark red bags.",
	"dark red bags contain 2 dark orange bags.",
	"dark orange bags contain 2 dark yellow bags.",
	"dark yellow bags contain 2 dark green bags.",
	"dark green bags contain 2 dark blue bags.",
	"dark blue bags contain 2 dark violet bags.",
	"dark violet bags contain no other bags.",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput)

	expected := []colorAndCount{
		{color: "shiny gold", count: 2},
		{color: "faded blue", count: 9},
	}

	assert.Equal(t, 9, len(data))
	assert.Equal(t, expected, data["muted yellow"])
}

func TestNumBagsInside1(t *testing.T) {
	data := parseInput(exampleInput)

	assert.Equal(t, 0, numberOfBagsInside("faded blue", data))
	assert.Equal(t, 11, numberOfBagsInside("vibrant plum", data))
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "4", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "155", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "32", result)
}

func TestPart2Example2(t *testing.T) {
	result := part2(exampleInput2)
	assert.Equal(t, "126", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "54803", result)
}
