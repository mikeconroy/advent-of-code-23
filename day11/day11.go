package day11

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day11/input")
	return part1(input), part2(input, 1_000_000)
}

/*
 * Input is an 'image'
 *	. - Empty Space
 * 	# - Galaxies
 * Sum of shortest paths between every pair of galaxies.
 * Any rows or columns that contain no galaxies should be twice as big.
 */
func part1(input []string) string {
	expandedImage := createExpandedImage(createImage(input))
	// printImage(expandedImage)
	galaxies := getGalaxies(expandedImage)

	return fmt.Sprint(getTotalDistance(galaxies))
}

func part2(input []string, expansionSize int) string {
	image := createImage(input)
	galaxies := getGalaxies(image)
	emptyRows := getEmptyRows(image)
	emptyCols := getEmptyCols(image)

	galaxies = expandGalaxies(galaxies, emptyRows, emptyCols, expansionSize)
	return fmt.Sprint(getTotalDistance(galaxies))
}

func getTotalDistance(galaxies []Point) int {
	totalDistance := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += getDistanceBetweenPoints(galaxies[i], galaxies[j])
		}
	}
	return totalDistance
}

func expandGalaxies(galaxies []Point, emptyRows []int, emptyCols []int, expansionSize int) []Point {
	var newGalaxies []Point
	for _, galaxy := range galaxies {
		expandedGalaxy := Point{x: galaxy.x, y: galaxy.y}
		for _, emptyRow := range emptyRows {
			if emptyRow < galaxy.y {
				expandedGalaxy.y = expandedGalaxy.y + expansionSize - 1
			}
		}

		for _, emptyCol := range emptyCols {
			if emptyCol < galaxy.x {
				expandedGalaxy.x = expandedGalaxy.x + expansionSize - 1
			}
		}
		newGalaxies = append(newGalaxies, expandedGalaxy)
	}
	return newGalaxies
}

func getEmptyRows(image [][]rune) (rows []int) {
	for y, row := range image {
		galaxiesInRow := false
		for _, val := range row {
			if val != '.' {
				galaxiesInRow = true
			}
		}
		if !galaxiesInRow {
			rows = append(rows, y)
		}
	}
	return rows
}

func getEmptyCols(image [][]rune) (cols []int) {
	for x := 0; x < len(image[0]); x++ {
		galaxiesInCol := false
		for _, row := range image {
			if row[x] != '.' {
				galaxiesInCol = true
			}
		}
		if !galaxiesInCol {
			cols = append(cols, x)
		}
	}
	return cols
}

type Point struct {
	x int
	y int
}

func getDistanceBetweenPoints(a Point, b Point) int {
	distance := 0
	if a.x > b.x {
		distance += a.x - b.x
	} else {
		distance += b.x - a.x
	}

	if a.y > b.y {
		distance += a.y - b.y
	} else {
		distance += b.y - a.y
	}

	return distance
}

func getGalaxies(image [][]rune) (galaxies []Point) {
	for y, row := range image {
		for x, val := range row {
			if val == '#' {
				galaxies = append(galaxies, Point{x: x, y: y})
			}
		}
	}
	return galaxies
}

func createImage(in []string) (image [][]rune) {
	for _, line := range in {
		var row []rune
		for _, val := range line {
			row = append(row, val)
		}
		image = append(image, row)
	}
	return image
}

func createExpandedImage(image [][]rune) (expandedImage [][]rune) {
	appendToColumn := make(map[int]bool)
	for x := 0; x < len(image[0]); x++ {
		galaxiesInCol := false
		for _, row := range image {
			if row[x] != '.' {
				galaxiesInCol = true
			}
		}
		if !galaxiesInCol {
			appendToColumn[x] = true
		}
	}
	for _, y := range image {
		var row []rune
		galaxiesInRow := false
		for x, val := range y {
			row = append(row, val)
			if appendToColumn[x] == true {
				row = append(row, '.')
			}
			if val == '#' {
				galaxiesInRow = true
			}
		}
		expandedImage = append(expandedImage, row)
		if !galaxiesInRow {
			expandedImage = append(expandedImage, row)
		}
	}

	return expandedImage
}

func printImage(image [][]rune) {
	for _, y := range image {
		for _, val := range y {
			fmt.Print(string(val))
		}
	}
}
