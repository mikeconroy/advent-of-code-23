package day18

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day18/input")
	return part1(input), part2(input)
}

func part1(input []string) string {
	// trenches := processPlan(input)
	// return fmt.Sprint(countCubicMeters(trenches))
	return fmt.Sprint(getArea(getPerimeter(input)))
}

// Shoelace algorithm + Pick's theorem
// Flood Fill is not efficient enough for the sizes involved in this part.
func part2(input []string) string {
	updatedInstructions := updateInstructions(input)
	perimeter := getPerimeter(updatedInstructions)
	area := getArea(perimeter)
	return fmt.Sprint(area)
}

// Uses Pick's Theorem to calculate the total area including the perimeter.
// A = i + b/2 - 1
// i - interior points		b - boundary points
func getArea(perim []Point) int {
	i := getInnerArea(perim)
	b := len(perim)
	a := i + (b / 2) + 1
	return a
}

// Uses Shoelace Formula to calculate the inside area.
func getInnerArea(perim []Point) int {
	var area int
	for i := 0; i < len(perim); i++ {
		p1 := perim[i]
		var p2 Point
		if i == len(perim)-1 {
			p2 = perim[0]
		} else {
			p2 = perim[i+1]
		}

		x1y2 := p1.x * p2.y
		x2y1 := p2.x * p1.y

		area += x1y2 - x2y1
	}
	return area / 2
}

func getPerimeter(instructions []string) []Point {
	// Calculate perimeter size
	size := 0
	for _, instruction := range instructions {
		count, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		size += count
	}

	perimeter := make([]Point, size)
	currentPoint := Point{x: 0, y: 0}
	index := 0
	for _, instruction := range instructions {
		instructionSplit := strings.Split(instruction, " ")
		dir := instructionSplit[0]
		n, _ := strconv.Atoi(instructionSplit[1])

		for step := 0; step < n; step++ {
			dy, dx := getDyDx(dir)
			newX := currentPoint.x + dx
			newY := currentPoint.y + dy
			currentPoint = Point{x: newX, y: newY}
			perimeter[index] = currentPoint
			index++
		}
	}
	return perimeter
}

func updateInstructions(in []string) []string {
	var updatedInstructions []string
	for _, instruction := range in {
		hex := strings.Split(instruction, " ")[2]
		// Remove brackets
		hex = hex[1 : len(hex)-1]

		var dir string
		if hex[len(hex)-1] == '0' {
			dir = "R"
		} else if hex[len(hex)-1] == '1' {
			dir = "D"
		} else if hex[len(hex)-1] == '2' {
			dir = "L"
		} else if hex[len(hex)-1] == '3' {
			dir = "U"
		}
		count := new(big.Int)
		count.SetString(hex[1:len(hex)-1], 16)

		updatedInstructions = append(updatedInstructions, dir+" "+count.String()+" #1")

	}
	return updatedInstructions
}

type Point struct {
	x int
	y int
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

// func processPlan(plan []string) [][]string {
// 	holes := make(map[Point]string)
// 	currentPoint := Point{x: 0, y: 0}
// 	holes[currentPoint] = "#FFFFFF"

// 	// Vars to hold the edges of the trenches so we can normalize the possitions to 0,0.
// 	// maxY is set to highest int as visually a lower Y is higher up the screen (0)
// 	minX, maxY := math.MaxInt, math.MaxInt
// 	maxX, minY := math.MinInt, math.MinInt
// 	// Load points into a map based on relative positions (x & y can go into negative values)
// 	for _, instruction := range plan {
// 		dir := strings.Split(instruction, " ")[0]
// 		n, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
// 		col := strings.Split(instruction, " ")[2]
// 		col = col[1 : len(col)-1]

// 		for steps := 0; steps < n; steps++ {
// 			dy, dx := getDyDx(dir)
// 			newX := currentPoint.x + dx
// 			newY := currentPoint.y + dy
// 			if newX < minX {
// 				minX = newX
// 			} else if newX > maxX {
// 				maxX = newX
// 			}
// 			if newY < maxY {
// 				maxY = newY
// 			} else if newY > minY {
// 				minY = newY
// 			}
// 			currentPoint = Point{x: newX, y: newY}
// 			holes[currentPoint] = col
// 		}
// 	}

// 	xSize := maxX - minX + 1
// 	ySize := minY - maxY + 1

// 	trenches := make([][]string, ySize)
// 	for row := 0; row < ySize; row++ {
// 		newRow := make([]string, xSize)
// 		trenches[row] = newRow
// 	}
// 	for pos, col := range holes {
// 		trenches[pos.y-maxY][pos.x-minX] = col
// 	}

// 	return trenches
// }

// func printTrenches(trenches [][]string) {
// 	for _, row := range trenches {
// 		for _, val := range row {
// 			if len(val) > 0 {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func countCubicMeters(trenches [][]string) int {
// 	totalArea := len(trenches) * len(trenches[0])

// 	// Keep track of already visited points
// 	seen := make(map[Point]bool)
// 	onEdge := make(map[Point]bool)
// 	var toProcess []Point

// 	// Add all the edge nodes to the toProcess list.
// 	// toProcess will then be used for BFS search.
// 	for x := 0; x < len(trenches[0]); x++ {
// 		toProcess = append(toProcess, Point{x: x, y: 0})
// 		toProcess = append(toProcess, Point{x: x, y: len(trenches) - 1})
// 	}
// 	for y := 0; y < len(trenches); y++ {
// 		toProcess = append(toProcess, Point{x: 0, y: y})
// 		toProcess = append(toProcess, Point{x: len(trenches[0]) - 1, y: y})
// 	}

// 	for len(toProcess) > 0 {
// 		currPoint := toProcess[0]
// 		if seen[currPoint] {
// 			toProcess = toProcess[1:]
// 			continue
// 		}
// 		seen[currPoint] = true

// 		if len(trenches[currPoint.y][currPoint.x]) == 0 {
// 			onEdge[currPoint] = true
// 			for _, neighbour := range getNeighbours(trenches, currPoint) {
// 				toProcess = append(toProcess, neighbour)
// 			}
// 		}

// 		toProcess = toProcess[1:]
// 	}

// 	return totalArea - len(onEdge)
// }

// func getNeighbours(trenches [][]string, pos Point) []Point {
// 	var neighbours []Point
// 	for _, posOffset := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
// 		newX := pos.x + posOffset.x
// 		newY := pos.y + posOffset.y

// 		if newY >= 0 && newY < len(trenches) && newX >= 0 && newX < len(trenches[newY]) {
// 			neighbours = append(neighbours, Point{x: newX, y: newY})
// 		}
// 	}
// 	return neighbours
// }
