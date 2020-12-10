// AOC2200 Day 4
package main

import (
	"errors"
	"strconv"
	"strings"
	"zhatt/aoc2020/aoc"
)

type instruction struct {
	opcode  string
	operand int
}

func parseInput(input []string) []instruction {
	instructions := make([]instruction, len(input), len(input))

	for index, line := range input {

		fields := strings.Fields(line)
		opcode := fields[0]
		operand, err := strconv.Atoi(fields[1])
		aoc.PanicOnError(err)

		instructions[index] = instruction{opcode: opcode, operand: operand}
	}

	return instructions
}

func simulate(instructions []instruction) (int, error) {
	pc := 0
	acc := 0
	visited := make(map[int]bool)
	var err error

	for {
		visited[pc] = true
		var nextAcc int
		var nextPc int
		switch instructions[pc].opcode {
		case "nop":
			nextPc = pc + 1
			nextAcc = acc

		case "acc":
			nextPc = pc + 1
			nextAcc = acc + instructions[pc].operand

		case "jmp":
			nextPc = pc + instructions[pc].operand
			nextAcc = acc
		default:
			aoc.PanicOnError(errors.New("unknown opcode"))
		}
		if visited[nextPc] {
			err = errors.New("looped")
			break
		}

		pc = nextPc
		acc = nextAcc
		if pc >= len(instructions) {
			break
		}
	}
	return acc, err
}

func part1(inputList []string) string {
	instructions := parseInput(inputList)
	finalAcc, _ := simulate(instructions)
	return strconv.Itoa(finalAcc)
}

func part2(inputList []string) string {
	instructions := parseInput(inputList)
	var finalAcc int
	var err error

	for index := range instructions {
		instructionsCopy := make([]instruction, len(instructions))
		copy(instructionsCopy, instructions)

		if instructionsCopy[index].opcode == "nop" {
			instructionsCopy[index].opcode = "jmp"

		} else if instructionsCopy[index].opcode == "jmp" {
			instructionsCopy[index].opcode = "nop"
		}
		finalAcc, err = simulate(instructionsCopy)
		if err == nil {
			break
		}
	}
	aoc.PanicOnError(err)
	return strconv.Itoa(finalAcc)
}

func main() {
	aoc.MainFunc(part1, part2)
}
