package day10

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day10/input")
	return part1(input), part2(input)
}

const (
	UP = iota
	DOWN
	RIGHT
	LEFT
)

func part1(input []string) string {
	area, start := parseInput(input)
	steps := 0

	pipe, direction := findFirstConnectingPipe(area, start)
	steps++
	for pipe.val != 'S' {
		pipe, direction = findNextPipe(area, pipe, direction)
		steps++
	}

	return fmt.Sprint(steps / 2)
}

func part2(input []string) string {
	// Same setup as Part 1 to first find the loop.
	area, start := parseInput(input)
	pipe, direction := findFirstConnectingPipe(area, start)
	loop := map[Position]bool{pipe: true}
	for pipe.val != 'S' {
		pipe, direction = findNextPipe(area, pipe, direction)
		loop[pipe] = true
	}

	// FOR DEBUGGING PURPOSES
	fmt.Println()
	for y := range area {
		for x, val := range area[y] {
			pipe := Position{x: x, y: y, val: val}
			if loop[pipe] {
				fmt.Print(string(val))
				// fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	/* Need to leak down in between pipes...
	 * No leak:		Leak:
	 * 	 F7			 7F
	 *   ||			 ||
	 * Flood fill the 'gaps' between pipes to find which nodes are connected to the outside.
	 * Detect gaps by recognising pipe shapes that create them.
	 * 	7F	7L	JF	JL	||	=
	 * Each step check surroundings for a node not part of the pipe - if found it's part of results.
	 *
	 * May be easier to double the number of rows & columns.
	 * Filling the gaps with '.' and connecting pipes with - & |.
	 * 	So Connections like:
	 * 		--		F-		L-		-J		-7		F7		LJ		FJ		7		F		|		|		|
	 * 					  													|		|		|		J		L
	 *  Become:
	 * 		---		F-7		l--		--J		--7		F-7		L-J		F-J		7		F		|		|		|
	 * 																		|		|		|		|		|
	 * 																		|		|		|		J		L
	 *
	 *	..........				....................
	 *	.S------7.				....................
	 *	.|F----7|.				..S--------------7..
	 *	.||....||.				..|..............|..
	 *	.||....||.		---->	..|.F----------7.|..
	 *	.|L-7F-J|.				..|.|..........|.|..
	 *	.|..||..|.				..|.|..........|.|..
	 *	.L--JL--J.				..|.|..........|.|..
	 *	..........				..|.|..........|.|..
	 *							..|.|..........|.|..
	 * 							..|.L---7.F---J|.|..
	 *							..|.....|.|....|.|..
	 *							..|.....|.|....|.|..
	 *							..|.....|.|....|.|..
	 *							..L-----J.|----J.|..
	 *							....................
	 *							....................
	 */
	return fmt.Sprint(0)
}

type Position struct {
	x   int
	y   int
	val rune
}

func parseInput(in []string) ([][]rune, Position) {
	var area [][]rune
	var start Position
	for y, line := range in {
		var row []rune
		for x, val := range line {
			row = append(row, val)
			if val == 'S' {
				start = Position{x: x, y: y, val: 'S'}
			}
		}
		area = append(area, row)
	}
	return area, start
}

func findFirstConnectingPipe(area [][]rune, start Position) (pipe Position, direction int) {
	// Find the first pipe connecting to the Start by checking all directions
	if start.y != 0 {
		pipeAbove := Position{x: start.x, y: start.y - 1, val: area[start.y-1][start.x]}

		if pipeAbove.val == '|' || pipeAbove.val == 'F' || pipeAbove.val == '7' {
			direction = UP
			pipe = pipeAbove
		}
	}
	if start.y != len(area)-1 {
		pipeBelow := Position{x: start.x, y: start.y + 1, val: area[start.y+1][start.x]}

		if pipeBelow.val == '|' || pipeBelow.val == 'J' || pipeBelow.val == 'L' {
			direction = DOWN
			pipe = pipeBelow
		}
	}
	if start.x != len(area[start.y]) {
		pipeRight := Position{x: start.x + 1, y: start.y, val: area[start.y][start.x+1]}
		if pipeRight.val == 'J' || pipeRight.val == '-' || pipeRight.val == '7' {
			direction = RIGHT
			pipe = pipeRight
		}
	}
	if start.x != 0 {
		pipeLeft := Position{x: start.x - 1, y: start.y, val: area[start.y][start.x-1]}
		if pipeLeft.val == 'F' || pipeLeft.val == '-' || pipeLeft.val == 'L' {
			direction = LEFT
			pipe = pipeLeft
		}
	}

	return pipe, direction
}

// Maps a Pipe Type & direction to the output direction.
var directionMap = map[rune]map[int]int{
	'L': {
		LEFT: UP,
		DOWN: RIGHT,
	},
	'J': {
		DOWN:  LEFT,
		RIGHT: UP,
	},
	'F': {
		LEFT: DOWN,
		UP:   RIGHT,
	},
	'7': {
		RIGHT: DOWN,
		UP:    LEFT,
	},
	'-': {
		LEFT:  LEFT,
		RIGHT: RIGHT,
	},
	'|': {
		UP:   UP,
		DOWN: DOWN,
	},
}

func findNextPipe(area [][]rune, pipe Position, direction int) (nextPipe Position, nextDirection int) {
	nextDirection = directionMap[pipe.val][direction]

	// No boundary checks as our input can never take us off the edge (at least for P1).
	if nextDirection == UP {
		nextPipe = Position{x: pipe.x, y: pipe.y - 1, val: area[pipe.y-1][pipe.x]}
	} else if nextDirection == DOWN {
		nextPipe = Position{x: pipe.x, y: pipe.y + 1, val: area[pipe.y+1][pipe.x]}
	} else if nextDirection == RIGHT {
		nextPipe = Position{x: pipe.x + 1, y: pipe.y, val: area[pipe.y][pipe.x+1]}
	} else if nextDirection == LEFT {
		nextPipe = Position{x: pipe.x - 1, y: pipe.y, val: area[pipe.y][pipe.x-1]}
	}

	return nextPipe, nextDirection
}

func getDirection(d int) string {
	if d == UP {
		return "UP"
	} else if d == DOWN {
		return "DOWN"
	} else if d == RIGHT {
		return "RIGHT"
	} else if d == LEFT {
		return "LEFT"
	}
	return "?"
}
