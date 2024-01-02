package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day18/input")
	return part1(input), part2(input)
}

func part1(input []string) string {
	trenches := processPlan(input)
	// printTrenches(trenches)
	return fmt.Sprint(countCubicMeters(trenches))
}

func countCubicMeters(trenches [][]string) int {
	totalArea := len(trenches) * len(trenches[0])

	// Keep track of already visited points
	seen := make(map[Point]bool)
	onEdge := make(map[Point]bool)
	var toProcess []Point

	// Add all the edge nodes to the toProcess list.
	// toProcess will then be used for BFS search.
	for x := 0; x < len(trenches[0]); x++ {
		toProcess = append(toProcess, Point{x: x, y: 0})
		toProcess = append(toProcess, Point{x: x, y: len(trenches) - 1})
	}
	for y := 0; y < len(trenches); y++ {
		toProcess = append(toProcess, Point{x: 0, y: y})
		toProcess = append(toProcess, Point{x: len(trenches[0]) - 1, y: y})
	}

	for len(toProcess) > 0 {
		currPoint := toProcess[0]
		if seen[currPoint] {
			toProcess = toProcess[1:]
			continue
		}
		seen[currPoint] = true

		if len(trenches[currPoint.y][currPoint.x]) == 0 {
			onEdge[currPoint] = true
			for _, neighbour := range getNeighbours(trenches, currPoint) {
				toProcess = append(toProcess, neighbour)
			}
		}

		toProcess = toProcess[1:]
	}

	return totalArea - len(onEdge)
}

func getNeighbours(trenches [][]string, pos Point) []Point {
	var neighbours []Point
	for _, posOffset := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		newX := pos.x + posOffset.x
		newY := pos.y + posOffset.y

		if newY >= 0 && newY < len(trenches) && newX >= 0 && newX < len(trenches[newY]) {
			neighbours = append(neighbours, Point{x: newX, y: newY})
		}
	}
	return neighbours
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Point struct {
	x int
	y int
}

func processPlan(plan []string) [][]string {
	holes := make(map[Point]string)
	currentPoint := Point{x: 0, y: 0}
	holes[currentPoint] = "#FFFFFF"

	// Vars to hold the edges of the trenches so we can normalize the possitions to 0,0.
	// maxY is set to highest int as visually a lower Y is higher up the screen (0)
	minX, maxY := math.MaxInt, math.MaxInt
	maxX, minY := math.MinInt, math.MinInt
	// Load points into a map based on relative positions (x & y can go into negative values)
	for _, instruction := range plan {
		dir := strings.Split(instruction, " ")[0]
		n, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		col := strings.Split(instruction, " ")[2]
		col = col[1 : len(col)-1]

		for steps := 0; steps < n; steps++ {
			dy, dx := getDyDx(dir)
			newX := currentPoint.x + dx
			newY := currentPoint.y + dy
			if newX < minX {
				minX = newX
			} else if newX > maxX {
				maxX = newX
			}
			if newY < maxY {
				maxY = newY
			} else if newY > minY {
				minY = newY
			}
			currentPoint = Point{x: newX, y: newY}
			holes[currentPoint] = col
		}
	}

	xSize := maxX - minX + 1
	ySize := minY - maxY + 1

	trenches := make([][]string, ySize)
	for row := 0; row < ySize; row++ {
		newRow := make([]string, xSize)
		trenches[row] = newRow
	}

	for pos, col := range holes {
		trenches[pos.y-maxY][pos.x-minX] = col
	}

	return trenches
}

func printTrenches(trenches [][]string) {
	for _, row := range trenches {
		for _, val := range row {
			if len(val) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getDyDx(dir string) (dy, dx int) {
	if dir == "R" {
		dx = 1
	} else if dir == "L" {
		dx = -1
	} else if dir == "U" {
		dy = -1
	} else if dir == "D" {
		dy = 1
	}
	return dy, dx
}
