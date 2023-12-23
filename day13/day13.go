package day13

import (
	"fmt"
	"reflect"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day13/input")
	return part1(input), part2(input)
}

/*
 *	Find reflection point - horizontal or vertical.
 *	Columns or rows on one side can be ignored if "run out of space".
 *	Result is sum of:
 *		Number of columns to the left of each vertical line.
 *		Number of rows above each horizontal line * 100.
 */
func part1(input []string) string {
	ters := loadTerrain(input)
	var result int
	for _, ter := range ters {
		result += ter.FindHorizontalReflection() * 100
		result += ter.FindVerticalReflection()
	}
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Terrain [][]rune

// Finds if a horizontal line reflection exists and returns the number of rows above the reflection.
// Or 0 if no reflection is found.
func (t Terrain) FindHorizontalReflection() int {
	// Reflection Point is the horizontal line we are checking.
	// 1 means we check the first 2 lines for reflection.
	// Then move the reflectionPoint to 2 and check lines 2 & 3 + 1 & 4 (if 2&3 match)
	for reflectionPoint := 1; reflectionPoint < len(t); reflectionPoint++ {
		rowsMatch := true
		for offset := 0; offset+reflectionPoint < len(t) && reflectionPoint-offset-1 >= 0; offset++ {
			rowA := t[offset+reflectionPoint]
			rowB := t[reflectionPoint-offset-1]

			if !reflect.DeepEqual(rowA, rowB) {
				rowsMatch = false
				break
			}
		}
		if rowsMatch {
			return reflectionPoint
		}
	}
	return 0
}

// Finds if a vertial line reflection exists and returns the number of columns to the left of the reflection.
// Or 0 if no reflection is found.
func (t Terrain) FindVerticalReflection() int {
	// Transposing the terrain and finding the horizontal reflection is the same result.
	var transposedT Terrain
	for x := 0; x < len(t[0]); x++ {
		var newRow []rune
		for _, row := range t {
			newRow = append(newRow, row[x])
		}
		transposedT = append(transposedT, newRow)
	}
	return transposedT.FindHorizontalReflection()
}

func (t Terrain) Print() {
	for _, row := range t {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
	fmt.Println()
}

func loadTerrain(in []string) (terrains []Terrain) {
	var terrain Terrain
	for _, line := range in {
		var row []rune
		if len(line) > 1 {
			for _, val := range line {
				row = append(row, val)
			}
			terrain = append(terrain, row)
		} else {
			terrains = append(terrains, terrain)
			terrain = Terrain{}
		}
	}
	terrains = append(terrains, terrain)
	return terrains
}
