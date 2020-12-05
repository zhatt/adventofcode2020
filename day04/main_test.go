package main

import (
	"testing"
	"zhatt/aoc2020/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
	"byr:1937 iyr:2017 cid:147 hgt:183cm",
	"",
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
	"hcl:#cfa07d byr:1929",
	"",
	"hcl:#ae17e1 iyr:2013",
	"eyr:2024",
	"ecl:brn pid:760753108 byr:1931",
	"hgt:179cm",
	"",
	"hcl:#cfa07d eyr:2025 pid:166559648",
	"iyr:2011 ecl:brn hgt:59in",
}

func TestParseInput(t *testing.T) {
	passportData := parseInput(exampleInput)
	_ = passportData

	assert.Equal(t, 4, len(passportData))
	assert.Equal(t, "350", passportData[1]["cid"])
	assert.Equal(t, "brn", passportData[3]["ecl"])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "2", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "230", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "2", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "156", result)
}
