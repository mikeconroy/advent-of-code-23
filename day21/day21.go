package day21

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day21/input")
	return part1(input, 64), part2(input)
}

type Point struct {
	x, y int
}

type Plot int

const (
	GARDEN Plot = iota
	ROCK
)

func part1(input []string, steps int) string {
	plots, start := loadMap(input)

	plotsReached := make(map[Point]bool)
	plotsReached[start] = true
	for step := 0; step < steps; step++ {
		plotsReached = processStep(plotsReached, plots)
	}
	return fmt.Sprint(len(plotsReached))
}

func processStep(positions map[Point]bool, plots map[Point]Plot) map[Point]bool {
	newPositions := make(map[Point]bool)
	for pos, _ := range positions {
		neighbours := []Point{
			{x: pos.x, y: pos.y - 1},
			{x: pos.x, y: pos.y + 1},
			{x: pos.x - 1, y: pos.y},
			{x: pos.x + 1, y: pos.y},
		}
		for _, neighbour := range neighbours {
			if plot, ok := plots[neighbour]; ok && plot == GARDEN && !newPositions[neighbour] {
				newPositions[neighbour] = true
			}
		}
	}
	return newPositions
}

// 26501365 steps with the map repeating in all directions infinitely.
func part2(input []string) string {
	fmt.Println(part1(input, 200))
	return fmt.Sprint(0)
}

func loadMap(in []string) (plots map[Point]Plot, start Point) {
	plots = make(map[Point]Plot)
	for y, line := range in {
		for x, val := range line {
			if val == '.' {
				plots[Point{x, y}] = GARDEN
			} else if val == '#' {
				plots[Point{x, y}] = ROCK
			} else if val == 'S' {
				start = Point{x, y}
				plots[start] = GARDEN
			}
		}
	}
	return plots, start
}
