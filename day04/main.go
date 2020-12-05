package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2020/aoc"
)

type passportData map[string]string

func parseInput(input []string) []passportData {
	data := []passportData{}

	currentRecord := passportData{}
	for _, line := range input {
		if line == "" {
			data = append(data, currentRecord)
			currentRecord = passportData{}
			continue
		}

		for _, fieldAndValue := range strings.Split(line, " ") {
			tokens := strings.Split(fieldAndValue, ":")
			currentRecord[tokens[0]] = tokens[1]
		}
	}
	data = append(data, currentRecord)

	return data
}

func isValidYear(yearString string, min int, max int) bool {
	year, err := strconv.ParseInt(yearString, 10, 0)

	return err == nil && int(year) >= min && int(year) <= max
}

func isValidByr(year string) bool {
	return isValidYear(year, 1920, 2002)
}

func isValidIyr(year string) bool {
	return isValidYear(year, 2010, 2020)
}

func isValidEyr(year string) bool {
	return isValidYear(year, 2020, 2030)
}

func isValidHgt(heightString string) bool {

	if len(heightString) < 2 {
		return false
	}

	height, err := strconv.ParseInt(heightString[:len(heightString)-2], 10, 0)

	if heightString[len(heightString)-2:] == "cm" {
		return err == nil && int(height) >= 150 && int(height) <= 193
	}
	
	if heightString[len(heightString)-2:] == "in" {
		return err == nil && int(height) >= 59 && int(height) <= 76

	} 
	return false
}

func isValidHcl(color string) bool {
	if color[0] != '#' {
		return false
	}
	if len(color) != 7 {
		return false
	}

	_, err := strconv.ParseInt(color[1:], 16, 0)

	return err == nil
}

func isValidEcl(color string) bool {
	switch color {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func isValidPid(pidString string) bool {
	if len(pidString) != 9 {
		return false
	}

	_, err := strconv.ParseInt(pidString, 10, 0)
	return err == nil
}

var requiredPassportFields = map[string]func(string) bool{
	"byr": isValidByr,
	"iyr": isValidIyr,
	"eyr": isValidEyr,
	"hgt": isValidHgt,
	"hcl": isValidHcl,
	"ecl": isValidEcl,
	"pid": isValidPid,
}

func countValidPasswords(passwordData []passportData, validate bool) int {
	numValid := 0

	for _, data := range passwordData {
		isGood := true
		for requiredField, validator := range requiredPassportFields {
			if value, exists := data[requiredField]; exists {
				if validate {
					isGood = validator(value)
					if !isGood {
						break
					}
				}

			} else {
				isGood = false
				break
			}
		}
		if isGood {
			numValid++
		}
	}
	return numValid
}

func part1(inputList []string) string {
	passwordData := parseInput(inputList)
	numValidPasswords := countValidPasswords(passwordData, false)
	return strconv.Itoa(numValidPasswords)
}

func part2(inputList []string) string {
	passwordData := parseInput(inputList)
	numValidPasswords := countValidPasswords(passwordData, true)
	return strconv.Itoa(numValidPasswords)
}

func main() {
	aoc.MainFunc(part1, part2)
}
