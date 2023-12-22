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
		// fmt.Println("Field:", field, getSpringsAsString(field.springs))
		arrangementsCount += countDp(field.springs, field.damagedGroups)
	}
	return fmt.Sprint(arrangementsCount)
}

/*
 * Based on:
 *	https://www.youtube.com/watch?v=g3Ms5e7Jdqo
 *	https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
 */
var cache = make(map[string]int)

func countDp(springs []rune, groups []int) int {
	key := makeKey(springs, groups)
	if cacheResult, ok := cache[key]; ok {
		return cacheResult
	}
	count := 0
	// Check if we still have springs left to check.
	if len(springs) == 0 {
		// If we don't have any springs left then check we also don't have any groups left.
		// If we don't have groups left then this is valid so return 1.
		// If we have groups left then this is invalid so return 0.
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}

	// If we have 0 groups left but still have a # in our springs then return 0 as this is invalid.
	// If we have no more springs then return 1 (all question marks will be .) which is a single arrangement.
	if len(groups) == 0 {
		if sliceContains(springs, '#') {
			return 0
		} else {
			return 1
		}
	}

	// If the first spring is . or ? then continue with the next spring.
	// This treats the ? as a .
	if springs[0] == '.' || springs[0] == '?' {
		count += countDp(springs[1:], groups)
	}

	// Handle the case where the spring is a # - treating the ? as a #.
	// This combined with the above if statement allows us to handle the ? as both types.
	if springs[0] == '#' || springs[0] == '?' {
		// Possibilities:
		// 		This group matches the expected group - in which case continue
		// 		This group matches but has an extra # on the end - invalid
		// 		This group does not match - not enough # or not enough length
		if groups[0] <= len(springs) {
			if !sliceContains(springs[:groups[0]], '.') {
				if groups[0] == len(springs) || springs[groups[0]] != '#' {
					if len(springs) == groups[0] {
						if len(groups) == 1 {
							return 1
						}
					} else {
						count += countDp(springs[groups[0]+1:], groups[1:])
					}
				}
			}
		}
	}
	cache[key] = count
	return count
}

func makeKey(springs []rune, groups []int) string {
	return fmt.Sprintf("%v_%v", springs, groups)
}

func sliceContains(slice []rune, search rune) bool {
	for _, val := range slice {
		if val == search {
			return true
		}
	}
	return false
}

// This structure for the recursion doesn't allow us to memoise (cache) results to speed up part 2.
// Kept in the code and used by Part 1 for posterity.
func countValidArrangements(field Field, springIndex int) int {
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
