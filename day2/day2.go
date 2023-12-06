package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day2/input")
	return part1(input), part2(input)
}

// 3117 is too high
func part1(input []string) string {
	red, green, blue := 12, 13, 14

	sumValidGames := 0

	for id, game := range input {
		if isValidGame(game, red, green, blue) {
			sumValidGames += id + 1
		}
	}

	return fmt.Sprint(sumValidGames)
}

func part2(input []string) string {
	totalPower := 0
	for _, game := range input {
		minRed, minBlue, minGreen := getMinColours(game)
		totalPower += minRed * minBlue * minGreen
	}
	return fmt.Sprint(totalPower)
}

func getMinColours(game string) (int, int, int) {
	minRed, minBlue, minGreen := 0, 0, 0
	game = strings.Split(game, ":")[1]
	for _, cubes := range strings.Split(game, ";") {
		for _, cubeStr := range strings.Split(cubes[1:], ", ") {
			cube := strings.Split(cubeStr, " ")
			count, _ := strconv.Atoi(cube[0])
			colour := cube[1]
			switch colour {
			case "red":
				if count > minRed {
					minRed = count
				}
			case "green":
				if count > minGreen {
					minGreen = count
				}
			case "blue":
				if count > minBlue {
					minBlue = count
				}
			}

		}
	}

	return minRed, minBlue, minGreen
}

func isValidGame(game string, red, green, blue int) bool {
	game = strings.Split(game, ":")[1]
	for _, cubes := range strings.Split(game, ";") {
		for _, cubeStr := range strings.Split(cubes[1:], ", ") {
			cube := strings.Split(cubeStr, " ")
			count, _ := strconv.Atoi(cube[0])

			colour := cube[1]
			switch colour {
			case "red":
				if count > red {
					return false
				}
			case "green":
				if count > green {
					return false
				}
			case "blue":
				if count > blue {
					return false
				}
			}
		}
	}
	return true
}
