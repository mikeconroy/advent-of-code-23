package day14

import (
	"reflect"
	"testing"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func TestDay14Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part1(input) != "136" {
		t.Fatal("Day 14 - Part 1 output should be 136")
	}
}

func TestDay14Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "64" {
		t.Fatal("Day 14 - Part 2 output should be 64")
	}
}

func TestCalculateTotalLoad(t *testing.T) {
	plat := Platform{
		[]rune("OOOO.#.O.."),
		[]rune("OO..#....#"),
		[]rune("OO..O##..O"),
		[]rune("O..#.OO..."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("..O..#.O.O"),
		[]rune("..O......."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}
	actual := plat.CalculateTotalLoad()

	if actual != 136 {
		t.Fatal("Day 14 - Total Load is wrong. Should be 136. Got:", actual)
	}
}

func TestTiltSouth(t *testing.T) {
	plat := Platform{
		[]rune("OOOO.#.O.."),
		[]rune("OO..#....#"),
		[]rune("OO..O##..O"),
		[]rune("O..#.OO..."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("..O..#.O.O"),
		[]rune("..O......."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}
	plat.TiltSouth()

	expected := Platform{
		[]rune(".....#...."),
		[]rune("....#....#"),
		[]rune("...O.##..."),
		[]rune("...#......"),
		[]rune("O.O....O#O"),
		[]rune("O.#..O.#.#"),
		[]rune("O....#...."),
		[]rune("OO....OO.."),
		[]rune("#OO..###.."),
		[]rune("#OO.O#...O"),
	}

	if !reflect.DeepEqual(plat, expected) {
		plat.Print()
		expected.Print()
		t.Fatal("Day 14 - Tilt South is not correct.")
	}
}

func TestTiltEast(t *testing.T) {
	plat := Platform{
		[]rune("OOOO.#.O.."),
		[]rune("OO..#....#"),
		[]rune("OO..O##..O"),
		[]rune("O..#.OO..."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("..O..#.O.O"),
		[]rune("..O......."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}
	plat.TiltEast()

	expected := Platform{
		[]rune(".OOOO#...O"),
		[]rune("..OO#....#"),
		[]rune("..OOO##..O"),
		[]rune("..O#....OO"),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("....O#..OO"),
		[]rune(".........O"),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}

	if !reflect.DeepEqual(plat, expected) {
		plat.Print()
		expected.Print()
		t.Fatal("Day 14 - Tilt East is not correct.")
	}
}
func TestTiltWest(t *testing.T) {
	plat := Platform{
		[]rune("OOOO.#.O.."),
		[]rune("OO..#....#"),
		[]rune("OO..O##..O"),
		[]rune("O..#.OO..."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("..O..#.O.O"),
		[]rune("..O......."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}
	plat.TiltWest()

	expected := Platform{
		[]rune("OOOO.#O..."),
		[]rune("OO..#....#"),
		[]rune("OOO..##O.."),
		[]rune("O..#OO...."),
		[]rune("........#."),
		[]rune("..#....#.#"),
		[]rune("O....#OO.."),
		[]rune("O........."),
		[]rune("#....###.."),
		[]rune("#....#...."),
	}

	if !reflect.DeepEqual(plat, expected) {
		plat.Print()
		expected.Print()
		t.Fatal("Day 14 - Tilt West is not correct.")
	}
}
