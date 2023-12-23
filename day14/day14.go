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

// 1 Cycle = Roll North, West, South, East
// We can't iterate 1 Billion times so instead we find the a loop
// and skip iterating the loop.
func part2(input []string) string {

	plat := loadPlatform(input)

	// Holds the first iteration the value was found at.
	cycles := make(map[string]int)
	cycleFound := false
	for x := 1; x <= 1_000_000_000; x++ {
		plat.TiltNorth()
		plat.TiltWest()
		plat.TiltSouth()
		plat.TiltEast()

		if !cycleFound {
			key := plat.Key()
			if iter, ok := cycles[key]; ok {
				cycleLen := x - iter
				remainingCycles := (1_000_000_000 - x) % cycleLen
				x = 1_000_000_000 - remainingCycles
				cycleFound = true
			} else {
				cycles[key] = x
			}
		}
	}
	return fmt.Sprint(plat.CalculateTotalLoad())
}

type Platform [][]rune

func (p Platform) Key() string {
	var key string
	for _, row := range p {
		for _, val := range row {
			key += string(val)
		}
	}
	return key
}

func (p Platform) TiltNorth() {
	for y, row := range p {
		for x, val := range row {
			if val == '.' {
				// Search for any Os in the column (above a #)
				// And move to this location if so.
				for rowIdx := y + 1; rowIdx < len(p); rowIdx++ {
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

func (p Platform) TiltSouth() {
	for y := len(p) - 1; y >= 0; y-- {
		for x, val := range p[y] {
			if val == '.' {
				for rowIdx := y - 1; rowIdx >= 0; rowIdx-- {
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

func (p Platform) TiltEast() {
	for x := len(p[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(p); y++ {
			if p[y][x] == '.' {
				for colIdx := x - 1; colIdx >= 0; colIdx-- {
					if p[y][colIdx] == '#' {
						break
					} else if p[y][colIdx] == 'O' {
						p[y][colIdx] = '.'
						p[y][x] = 'O'
						break
					}
				}
			}
		}
	}
}

func (p Platform) TiltWest() {
	for x := 0; x < len(p[0]); x++ {
		for y := 0; y < len(p); y++ {
			if p[y][x] == '.' {
				for colIdx := x + 1; colIdx < len(p[0]); colIdx++ {
					if p[y][colIdx] == '#' {
						break
					} else if p[y][colIdx] == 'O' {
						p[y][colIdx] = '.'
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
