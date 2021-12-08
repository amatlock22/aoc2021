package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const diagramSize = 1000

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diagram := initDiagram()

	for scanner.Scan() {
		x1, y1, x2, y2 := processCoordinates(scanner.Text())

		if x1 == x2 {
			verticalLine(diagram, x1, y1, y2)
		} else if y1 == y2 {
			horizontalLine(diagram, y1, x1, x2)
		}
	}

	return countOverlaps(diagram)
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diagram := initDiagram()

	for scanner.Scan() {
		x1, y1, x2, y2 := processCoordinates(scanner.Text())

		if x1 == x2 {
			verticalLine(diagram, x1, y1, y2)
		} else if y1 == y2 {
			horizontalLine(diagram, y1, x1, x2)
		} else {
			diagonalLine(diagram, x1, x2, y1, y2)
		}
	}

	return countOverlaps(diagram)
}

func processCoordinates(line string) (x1, y1, x2, y2 int) {
	coordinates := strings.Split(line, " -> ")
	x := strings.Split(coordinates[0], ",")
	y := strings.Split(coordinates[1], ",")

	x1, _ = strconv.Atoi(x[0])
	y1, _ = strconv.Atoi(x[1])
	x2, _ = strconv.Atoi(y[0])
	y2, _ = strconv.Atoi(y[1])

	return
}

func verticalLine(diagram [][]int, x, y1, y2 int) {
	diffY := diff(y1, y2)
	minY := min(y1, y2)

	for y := minY; y <= minY+diffY; y++ {
		diagram[y][x]++
	}
}

func horizontalLine(diagram [][]int, y, x1, x2 int) {
	diffX := diff(x1, x2)
	minX := min(x1, x2)

	for x := minX; x <= minX+diffX; x++ {
		diagram[y][x]++
	}
}

func diagonalLine(diagram [][]int, x1, x2, y1, y2 int) {
	diffX := diff(x1, x2)
	minX := min(x1, x2)

	if (x1 > x2 && y1 > y2) || (x1 < x2 && y1 < y2) { // both are increasing or decreasing
		for i, j := minX, min(y1, y2); i <= minX+diffX; i, j = i+1, j+1 {
			diagram[j][i]++
		}
	} else { // only one increases, the other decreases
		for i, j := minX, max(y1, y2); i <= minX+diffX; i, j = i+1, j-1 {
			diagram[j][i]++
		}
	}
}

func countOverlaps(diagram [][]int) int {
	overlapCount := 0

	for _, row := range diagram {
		for _, column := range row {
			if column >= 2 {
				overlapCount++
			}
		}
	}

	return overlapCount
}

func initDiagram() [][]int {
	diagram := make([][]int, diagramSize)

	for i := 0; i < diagramSize; i++ {
		diagram[i] = make([]int, diagramSize)
	}

	return diagram
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
