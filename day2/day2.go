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
	return "Part 2"
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
