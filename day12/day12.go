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
	// For each field:
	// 		Create combinations of operational & non-operational springs for all the unknowns.
	// 		Check whether the replacement unknowns match the groupings or not
	// This is probably quite a naive approach as we are creating groupings that cannot exists.
	// May be a better algorithm using recursion to only create possible values for the unknowns.
	arrangementsCount := 0
	for _, field := range fields {
		combos := generateCombinations(field.unknowns)
		for _, combo := range combos {
			newField := replaceUnknownSprings(field, combo)
			if isValidArrangement(newField) {
				arrangementsCount++
			}
		}
	}
	return fmt.Sprint(arrangementsCount)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

var combosCache = map[int][][]rune{
	1: {[]rune{'.'}, []rune{'#'}},
}

// Seems suitable for up to a count of 20.
func generateCombinations(count int) [][]rune {
	var combos [][]rune
	if cachedCombos, ok := combosCache[count]; ok {
		return cachedCombos
	} else {
		for _, generatedCombo := range generateCombinations(count - 1) {
			newComboDot := append([]rune{'.'}, generatedCombo...)
			combos = append(combos, newComboDot)
			newComboHash := append([]rune{'#'}, generatedCombo...)
			combos = append(combos, newComboHash)
		}
	}
	combosCache[count] = combos
	return combos
}

func replaceUnknownSprings(field Field, unknownsReplacement []rune) Field {
	var newSprings []rune
	currentUnknown := 0
	for _, spring := range field.springs {
		if spring == '?' {
			newSprings = append(newSprings, unknownsReplacement[currentUnknown])
			currentUnknown++
		} else {
			newSprings = append(newSprings, spring)
		}
	}
	return Field{springs: newSprings, damagedGroups: field.damagedGroups, unknowns: 0}
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
				currentGroupIndex++
			}
			inGroupOfDamagedSprings = false
		}
	}

	// fmt.Println("COMPARING ARRANGEMENT WITH EXPECTED", arrangement, field.damagedGroups)

	if slicesAreEqual(arrangement, field.damagedGroups) {
		// fmt.Println("SLICES ARE EQUAL, RETURNING TRUE")
		return true
	} else {
		return false
	}

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
