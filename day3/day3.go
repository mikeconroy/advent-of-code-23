package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day3/input")
	eng := LoadSchematic(input)
	return part1(eng), part2(eng)
}

/*
 * Any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum.
 */
func part1(eng Engine) string {
	// This approach wasn't ideal - works by locating symbols and then surrounding numbers.
	// Better (quicker & easier) approach would be to find numbers and then check for adjacent symbols instead.
	symbols := eng.FindSymbols()
	validNums := make(map[adjNumber]bool)
	for _, symbol := range symbols {
		adjacents := eng.GetAdjacents(symbol)
		for _, adj := range adjacents {
			if unicode.IsDigit(adj.char) {
				validNums[eng.NumberAt(adj)] = true
			}
		}
	}

	var result int

	for key, _ := range validNums {
		result += key.num
	}

	return fmt.Sprint(result)
}

func part2(eng Engine) string {
	return fmt.Sprint(0)
}

type Engine struct {
	schematic [][]rune
}

func LoadSchematic(input []string) Engine {
	var schematic [][]rune
	for _, line := range input {
		schematic = append(schematic, []rune(line))
	}
	return Engine{
		schematic: schematic,
	}
}

func (eng Engine) Print() {
	for _, row := range eng.schematic {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

type position struct {
	char rune
	x, y int
}

type adjNumber struct {
	num          int
	xStart, xEnd int
	y            int
}

func (eng Engine) FindSymbols() []position {
	var symbols []position
	for y, row := range eng.schematic {
		for x, char := range row {
			// Check if it's a symbol.
			// A symbol is any char that's not a . and not a digit.
			// if char != '.' && (char <= 48 || char > 58) {
			if char != '.' && !unicode.IsDigit(char) {
				symbols = append(symbols, position{char, x, y})
			}
		}
	}
	return symbols
}

func (eng Engine) GetAdjacents(pos position) []position {
	var adjacents []position
	for yOffset := -1; yOffset < 2; yOffset++ {
		for xOffset := -1; xOffset < 2; xOffset++ {
			newX := pos.x + xOffset
			newY := pos.y + yOffset

			if newY >= 0 && newY < len(eng.schematic) &&
				newX >= 0 && newX < len(eng.schematic[newY]) {
				newPos := position{
					char: eng.schematic[newY][newX],
					x:    newX,
					y:    newY,
				}
				adjacents = append(adjacents, newPos)
			}
		}
	}
	return adjacents
}

func (eng Engine) NumberAt(pos position) adjNumber {
	var xStart, xEnd int
	y := pos.y

	for x := pos.x; x >= 0; x-- {
		if unicode.IsDigit(eng.schematic[y][x]) {
			xStart = x
		} else {
			break
		}
	}

	for x := pos.x; x < len(eng.schematic); x++ {
		if unicode.IsDigit(eng.schematic[y][x]) {
			xEnd = x
		} else {
			break
		}
	}

	var valStr string
	for x := xStart; x <= xEnd; x++ {
		valStr += string(eng.schematic[y][x])
	}
	val, _ := strconv.Atoi(valStr)

	return adjNumber{
		num:    val,
		xStart: xStart,
		xEnd:   xEnd,
		y:      y,
	}
}
