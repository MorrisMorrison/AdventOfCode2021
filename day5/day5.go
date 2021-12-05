package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./day5/input.txt"

type Line struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)
	diagram := make([][]int, 1000)
	for i := 0; i < len(diagram); i++ {

		diagram[i] = make([]int, 1000)

	}
	for _, line := range input {
		if !isLineValid(line) {
			continue
		}
		if isHorizontal(line) {
			isIncreasing := line.From.X < line.To.X
			if isIncreasing {
				for i := line.From.X; i <= line.To.X; i++ {
					diagram[line.From.Y][i] += 1
				}
			} else {
				for i := line.From.X; i >= line.To.X; i-- {
					diagram[line.From.Y][i] += 1
				}
			}
		}
		if isVertical(line) {
			isIncreasing := line.From.Y < line.To.Y
			if isIncreasing {
				for i := line.From.Y; i <= line.To.Y; i++ {
					diagram[i][line.From.X] += 1
				}
			} else {
				for i := line.From.Y; i >= line.To.Y; i-- {
					diagram[i][line.From.X] += 1
				}
			}
		}
	}

	fmt.Println()
	for _, row := range diagram {
		fmt.Println(row)
	}
	fmt.Println()
	fmt.Println("Overlapping lines: ", countOverlappingLines(diagram))

	return ""
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)
	diagram := make([][]int, 1000)
	for i := 0; i < len(diagram); i++ {
		diagram[i] = make([]int, 1000)
	}

	for _, line := range input {
		x := line.From.X
		y := line.From.Y
		var distance int
		if line.From.X != line.To.X {
			distance = int(math.Abs(float64(line.From.X - line.To.X)))
		} else {
			distance = int(math.Abs(float64(line.From.Y - line.To.Y)))
		}

		for i := 0; i <= distance; i++ {

			diagram[y][x] += 1

			if line.From.X < line.To.X {
				x += 1
			}
			if line.From.X > line.To.X {
				x -= 1
			}
			if line.From.Y < line.To.Y {
				y += 1
			}
			if line.From.Y > line.To.Y {
				y -= 1
			}
		}
	}

	fmt.Println()
	for _, row := range diagram {
		fmt.Println(row)
	}

	fmt.Println()
	fmt.Println("Overlapping lines: ", countOverlappingLines(diagram))

	return ""
}

func parsePoint(text string) Point {
	coordinates := strings.Split(text, ",")
	x, err := strconv.Atoi(coordinates[0])
	check(err)
	y, err := strconv.Atoi(coordinates[1])
	check(err)
	return Point{X: x, Y: y}
}

func isHorizontal(line Line) bool {
	return line.From.Y == line.To.Y && line.From.X != line.To.X
}

func isVertical(line Line) bool {
	return line.From.X == line.To.X && line.From.Y != line.To.Y
}

func isLineValid(line Line) bool {
	return isHorizontal(line) || isVertical(line)
}

func countOverlappingLines(diagram [][]int) int {
	var count int
	for row := 0; row < len(diagram); row++ {
		for col := 0; col < len(diagram[0]); col++ {
			if diagram[row][col] > 1 {
				count++
			}
		}
	}

	return count
}

func parseInput(file *os.File) []Line {
	scanner := bufio.NewScanner(file)
	input := make([]Line, 0)
	var row int
	for scanner.Scan() {
		inputLineContent := scanner.Text()
		fields := strings.Fields(inputLineContent)
		line := Line{From: parsePoint(fields[0]), To: parsePoint(fields[2])}
		input = append(input, line)
		row++
	}
	return input
}
