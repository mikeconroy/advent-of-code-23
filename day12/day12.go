package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day12/input")
	return part1(input), part2(input)
}

type Field struct {
	springs       []rune
	damagedGroups []int
	unknowns      int
}

func part1(input []string) string {
	fields := createFields(input)
	arrangementsCount := 0
	for _, field := range fields {
		arrangementsCount += countValidArrangements(field, 0)
	}
	return fmt.Sprint(arrangementsCount)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

func countValidArrangements(field Field, springIndex int) int {
	// TODO: Optimization to check validity as we go to reduce checking every possibility.
	// E.g. in some cases a # at the start will be invalid in which case there is no point verifying every
	// Combination after that.
	// We can also stop recursing once all unknowns are populated.

	// fmt.Println("Count Valid Arrangements:", getSpringsAsString(field.springs), springIndex)
	count := 0
	if springIndex >= len(field.springs) || field.unknowns == 0 {
		if isValidArrangement(field) {
			// fmt.Println("Valid Arrangement.")
			return 1
		} else {
			// fmt.Println("Invalid Arrangement.")
			return 0
		}
	}
	if field.springs[springIndex] == '?' {
		// Create a new field with the Unknown replaced with a .
		springsDot := make([]rune, len(field.springs))
		copy(springsDot, field.springs)
		springsDot[springIndex] = '.'
		fieldDot := Field{springs: springsDot, damagedGroups: field.damagedGroups, unknowns: field.unknowns - 1}

		// Create a new field with the Unknown replaced with a #
		springsHash := make([]rune, len(field.springs))
		copy(springsHash, field.springs)
		springsHash[springIndex] = '#'
		fieldHash := Field{springs: springsHash, damagedGroups: field.damagedGroups, unknowns: field.unknowns - 1}
		springIndex++
		// fmt.Println("springsDot", getSpringsAsString(springsDot), "springsHash", getSpringsAsString(springsHash))
		count += countValidArrangements(fieldDot, springIndex)
		count += countValidArrangements(fieldHash, springIndex)
	} else {
		// TODO: It may be more efficient here to find the next unknown and calling the function with that index.
		count += countValidArrangements(field, springIndex+1)
	}
	return count
}

func getSpringsAsString(springs []rune) string {
	result := ""
	for _, spring := range springs {
		result += string(spring)
	}
	return result
}

func slicesAreEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

func isValidArrangement(field Field) bool {
	arrangement := make([]int, len(field.damagedGroups))
	currentGroupIndex := 0
	inGroupOfDamagedSprings := false
	// fmt.Println(field.springs, field.damagedGroups)
	for _, spring := range field.springs {
		if spring == '#' {
			if currentGroupIndex >= len(arrangement) {
				return false
			}
			arrangement[currentGroupIndex] = arrangement[currentGroupIndex] + 1
			inGroupOfDamagedSprings = true
		} else {
			if inGroupOfDamagedSprings {
				if field.damagedGroups[currentGroupIndex] != arrangement[currentGroupIndex] {
					return false
				}
				currentGroupIndex++
			}
			inGroupOfDamagedSprings = false
		}
	}

	if len(arrangement) != len(field.damagedGroups) {
		return false
	} else {
		// Check the last value of the groupings are equal
		if arrangement[len(arrangement)-1] != field.damagedGroups[len(field.damagedGroups)-1] {
			return false
		}
	}
	return true
}

func printCombinations(combos [][]rune) {
	for _, line := range combos {
		for _, val := range line {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func createFields(in []string) (fields []Field) {
	for _, line := range in {
		springsRaw := strings.Split(line, " ")[0]
		groupsRaw := strings.Split(line, " ")[1]
		unknowns := 0
		var springs []rune
		for _, spring := range springsRaw {
			springs = append(springs, spring)
			if spring == '?' {
				unknowns++
			}
		}
		var groups []int
		for _, groupRaw := range strings.Split(groupsRaw, ",") {
			group, _ := strconv.Atoi(string(groupRaw))
			groups = append(groups, group)
		}
		fields = append(fields, Field{springs: springs, damagedGroups: groups, unknowns: unknowns})
	}
	return fields
}
