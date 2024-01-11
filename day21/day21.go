package day21

import (
	"fmt"
	"math"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day21/input")
	return part1(input, 64, 131), part2(input, 131)
}

type Point struct {
	x, y int
}

type Plot int

const (
	GARDEN Plot = iota
	ROCK
)

func part1(input []string, steps int, length int) string {
	plots, start := loadMap(input)

	plotsReached := make(map[Point]bool)
	plotsReached[start] = true
	for step := 0; step < steps; step++ {
		plotsReached = processStep(plotsReached, plots, length)
	}
	return fmt.Sprint(len(plotsReached))
}

func processStep(positions map[Point]bool, plots map[Point]Plot, length int) map[Point]bool {
	newPositions := make(map[Point]bool)
	for pos, _ := range positions {
		neighbours := []Point{
			{x: pos.x, y: pos.y - 1},
			{x: pos.x, y: pos.y + 1},
			{x: pos.x - 1, y: pos.y},
			{x: pos.x + 1, y: pos.y},
		}
		for _, neighbour := range neighbours {
			neighbourOffset := Point{x: neighbour.x % length, y: neighbour.y % length}
			if neighbourOffset.x < 0 {
				neighbourOffset.x = length + neighbourOffset.x
			}

			if neighbourOffset.y < 0 {
				neighbourOffset.y = length + neighbourOffset.y
			}

			// if (len(positions) == 16 || len(positions) == 22) && neighbour != neighbourOffset {
			// 	fmt.Println("Original:", neighbour, "Translated:", neighbourOffset, "Type:", plots[neighbourOffset])
			// }
			if plot, ok := plots[neighbourOffset]; ok && plot == GARDEN && !newPositions[neighbour] {
				newPositions[neighbour] = true
			}
		}
	}
	return newPositions
}

// 26_501_365 steps with the map repeating in all directions infinitely.
// https://www.reddit.com/r/adventofcode/comments/18orn0s
// https://www.reddit.com/r/adventofcode/comments/18oh5f7/comment/keh27rk
// https://www.dcode.fr/lagrange-interpolating-polynomial
// (65, 3867) (196, 34253) (327, 94909) -> 619407349431167
func part2(input []string, length int) string {
	plots, start := loadMap(input)
	plotsReached := make(map[Point]bool)
	plotsReached[start] = true
	var y0, y1, y2 int
	for step := 1; step <= 327; step++ {
		plotsReached = processStep(plotsReached, plots, length)
		if step == 65 {
			y0 = len(plotsReached)
		} else if step == (131 + 65) {
			y1 = len(plotsReached)
		} else if step == ((2 * 131) + 65) {
			y2 = len(plotsReached)
		}
	}

	// y0 := 3867
	// y1 := 34253
	// y2 := 94909

	x := (26_501_365 - 65) / 131

	res := ((y0 * (int(math.Pow(float64(x), 2)) - (3 * x) + 2)) / 2) - (y1 * (int(math.Pow(float64(x), 2)) - (2 * x))) + (y2 * (int(math.Pow(float64(x), 2)) - x) / 2)
	return fmt.Sprint(res)
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
