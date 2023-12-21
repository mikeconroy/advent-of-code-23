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
	fields := createFields(input)
	return "0"
	arrangementsCount := 0
	for _, field := range fields {
		ogSprings := make([]rune, len(field.springs))
		ogDamagedGroups := make([]int, len(field.damagedGroups))
		copy(ogSprings, field.springs)
		copy(ogDamagedGroups, field.damagedGroups)
		for i := 0; i < 4; i++ {
			field.springs = append(field.springs, '?')
			field.springs = append(field.springs, ogSprings...)
			field.damagedGroups = append(field.damagedGroups, ogDamagedGroups...)
		}
		field.unknowns = (field.unknowns * 5) + 4
		fmt.Println("Field:", field, getSpringsAsString(field.springs))
		arrangementsCount += countValidArrangements(field, 0)
	}
	return fmt.Sprint(arrangementsCount)
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

		// We only need to check for possibly valid for the field we have added a # to.
		// As adding a . should not affect the possibility of it being valid.
		// Possible check for validity =
		// 		len(actual groups) < len(expectedDamagedGroups)
		// 		&& Groups created so far match the expected values.
		// This could be further improved as it's not efficient repeatedly checking the same groups...
		if isPossiblyValid(fieldHash) {
			count += countValidArrangements(fieldHash, springIndex)
		}

	} else {
		// TODO: It may be more efficient here to find the next unknown and calling the function with that index.
		count += countValidArrangements(field, springIndex+1)
	}
	return count
}

func isPossiblyValid(field Field) bool {
	count := 0
	currentIndex := 0

	for _, spring := range field.springs {
		if spring == '#' {
			// if len(actualGroups) >= len(field.damagedGroups) && count == 0 {
			// 	fmt.Println("NOT POSSIBLY VALID too many groups.", getSpringsAsString(field.springs), field.damagedGroups)
			// 	return false
			// }
			count++
			// fmt.Println(count, currentIndex, getSpringsAsString(field.springs))
			if currentIndex >= len(field.damagedGroups) || count > field.damagedGroups[currentIndex] {
				// fmt.Println("NOT POSSIBLY VALID group count exceeded", currentIndex, count, getSpringsAsString(field.springs), field.damagedGroups)
				return false
			}
		} else if spring == '.' {
			if count > 0 {
				if count != field.damagedGroups[currentIndex] {
					return false
				}
				currentIndex++
				count = 0
			}
		} else {
			// When we hit an unknown spring return true as it is hard to reason about validity at this point.
			return true
		}
	}
	return true
}

func getSpringsAsString(springs []rune) string {
	result := ""
	for _, spring := range springs {
		result += string(spring)
	}
	return result
}

// Unused - slices are compared inline of checking validity
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
