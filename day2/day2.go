package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./day2/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type submarinePosition struct {
	verticalPosition   int
	horizontalPosition int
	aim                int
}

type instruction struct {
	value           int
	instructionType instructionType
}

type instructionType int

const (
	up      instructionType = 0
	down    instructionType = 1
	forward instructionType = 2
)

func Star1() submarinePosition {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	var sumUp int
	var sumDown int
	var sumForward int

	for _, instruction := range input {
		switch instruction.instructionType {
		case up:
			sumUp += instruction.value
		case down:
			sumDown += instruction.value
		case forward:
			sumForward += instruction.value
		}
	}

	return submarinePosition{verticalPosition: sumForward, horizontalPosition: sumDown - sumUp}
}

func Star2() submarinePosition {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)
	submarinePosition := submarinePosition{0, 0, 0}
	for _, instruction := range input {
		switch instruction.instructionType {
		case up:
			submarinePosition.aim -= instruction.value
		case down:
			submarinePosition.aim += instruction.value
		case forward:
			submarinePosition.horizontalPosition += instruction.value
			submarinePosition.verticalPosition += submarinePosition.aim * instruction.value
		}
	}

	return submarinePosition
}

func parseInput(file *os.File) []instruction {
	scanner := bufio.NewScanner(file)
	input := make([]instruction, 0)
	for scanner.Scan() {
		instructionText := scanner.Text()
		instructionTextSlice := strings.Fields(instructionText)
		instructionType := parseInstructionType(instructionTextSlice[0])
		value, err := strconv.Atoi(instructionTextSlice[1])
		check(err)
		instruction := instruction{value, instructionType}
		input = append(input, instruction)
	}

	return input
}

func parseInstructionType(instructionType string) instructionType {
	switch instructionType {
	case "up":
		return up
	case "down":
		return down
	case "forward":
		return forward
	}

	return -1
}
