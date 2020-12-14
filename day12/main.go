// AOC2200 Day 12
package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2020/aoc"
	"zhatt/aoc2020/coord"
)

type partType int

const (
	PART1 partType = iota
	PART2
)

type commandType int

const (
	cmdLeft    commandType = 'L'
	cmdRight   commandType = 'R'
	cmdNorth   commandType = 'N'
	cmdSouth   commandType = 'S'
	cmdEast    commandType = 'E'
	cmdWest    commandType = 'W'
	cmdForward commandType = 'F'
)

type actionType struct {
	command commandType
	amount  int
}

type headingType int

const (
	headingNorth headingType = 0
	headingEast  headingType = 90
	headingSouth headingType = 180
	headingWest  headingType = 270
)

func (heading *headingType) add(degrees int) {
	*heading = (*heading + headingType(degrees)) % 360

	for *heading < 0 {
		*heading += 360
	}
}

type shipType struct {
	startingLocation coord.Coord
	location         coord.Coord
	waypoint         coord.Coord
	heading          headingType
}

func newShip() *shipType {
	return &shipType{
		heading:  headingEast,
		waypoint: coord.Coord{Xval: 10, Yval: 1},
	}
}

func moveLocation(location *coord.Coord, action actionType) {
	switch action.command {
	case cmdNorth:
		increment := coord.Coord{Xval: 0, Yval: action.amount}
		*location = coord.Add(*location, increment)
	case cmdSouth:
		increment := coord.Coord{Xval: 0, Yval: -action.amount}
		*location = coord.Add(*location, increment)
	case cmdEast:
		increment := coord.Coord{Xval: action.amount, Yval: 0}
		*location = coord.Add(*location, increment)
	case cmdWest:
		increment := coord.Coord{Xval: -action.amount, Yval: 0}
		*location = coord.Add(*location, increment)
	case cmdForward, cmdLeft, cmdRight:
		panic("bug")
	}
}

// Normalize to [0:360)
func normalizeDegrees(degrees int) int {
	for degrees < 0 {
		degrees += 360
	}
	return degrees % 360
}

func rotateWaypoint(waypoint coord.Coord, degrees int) coord.Coord {
	degrees = normalizeDegrees(degrees)

	for degrees > 0 {
		waypoint = coord.Coord{Xval: waypoint.Yval, Yval: -waypoint.Xval}
		degrees -= 90
	}

	if degrees != 0 {
		panic("bug")
	}

	return waypoint
}

func (ship *shipType) move(action actionType, part partType) {
	switch action.command {
	case cmdLeft:
		if part == PART1 {
			ship.heading.add(-action.amount)
		} else {
			ship.waypoint = rotateWaypoint(ship.waypoint, -action.amount)
		}
	case cmdRight:
		if part == PART1 {
			ship.heading.add(action.amount)
		} else {
			ship.waypoint = rotateWaypoint(ship.waypoint, action.amount)
		}

	case cmdNorth, cmdSouth, cmdEast, cmdWest:
		if part == PART1 {
			moveLocation(&ship.location, action)
		} else {
			moveLocation(&ship.waypoint, action)
		}
	case cmdForward:
		if part == PART1 {
			switch ship.heading {
			case headingNorth:
				ship.move(actionType{command: cmdNorth, amount: action.amount}, PART1)
			case headingEast:
				ship.move(actionType{command: cmdEast, amount: action.amount}, PART1)
			case headingSouth:
				ship.move(actionType{command: cmdSouth, amount: action.amount}, PART1)
			case headingWest:
				ship.move(actionType{command: cmdWest, amount: action.amount}, PART1)

			default:
				panic("bad heading")
			}
		} else {
			// NB.  These may be negative.
			northAmount := ship.waypoint.Yval * action.amount
			eastAmount := ship.waypoint.Xval * action.amount
			ship.move(actionType{command: cmdNorth, amount: northAmount}, PART1)
			ship.move(actionType{command: cmdEast, amount: eastAmount}, PART1)
		}

	default:
		panic("bad command")
	}
}

func (ship *shipType) distanceTraveled() int {
	return coord.DistanceManhattan(ship.location, ship.startingLocation)
}

func parseInput(input []string) []actionType {
	commands := []actionType{}

	for _, line := range input {
		var action actionType
		_, err := fmt.Sscanf(line, "%c%d", &action.command, &action.amount)
		aoc.PanicOnError(err)
		commands = append(commands, action)
	}
	return commands
}

func part1(inputList []string) string {
	actions := parseInput(inputList)
	ship := newShip()

	for _, action := range actions {
		ship.move(action, PART1)
	}
	distance := ship.distanceTraveled()
	return strconv.Itoa(distance)
}

func part2(inputList []string) string {
	actions := parseInput(inputList)
	ship := newShip()

	for _, action := range actions {
		ship.move(action, PART2)
	}
	distance := ship.distanceTraveled()
	return strconv.Itoa(distance)
}

func main() {
	aoc.MainFunc(part1, part2)
}
