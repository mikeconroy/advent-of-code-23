package day14

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day14/input")
	return part1(input), part2(input)
}

func part1(input []string) string {
	plat := loadPlatform(input)
	plat.TiltNorth()
	return fmt.Sprint(plat.CalculateTotalLoad())
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Platform [][]rune

func (p Platform) TiltNorth() {
	for y, row := range p {
		for x, val := range row {
			if val == '.' {
				// Search for any Os in the column (above a #)
				// And move to this location if so.
				for rowIdx := y; rowIdx < len(p); rowIdx++ {
					if p[rowIdx][x] == '#' {
						break
					} else if p[rowIdx][x] == 'O' {
						p[rowIdx][x] = '.'
						p[y][x] = 'O'
						break
					}
				}
			}
		}
	}
}

func (p Platform) CalculateTotalLoad() int {
	var totalLoad int
	for y, row := range p {
		for _, val := range row {
			if val == 'O' {
				totalLoad += len(p) - y
			}
		}
	}
	return totalLoad
}

func (p Platform) Print() {
	for _, row := range p {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
	fmt.Println()
}
func loadPlatform(in []string) (plat Platform) {
	for _, line := range in {
		var row []rune
		for _, val := range line {
			row = append(row, val)
		}
		plat = append(plat, row)
	}
	return plat
}
