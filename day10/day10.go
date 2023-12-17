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

// This code could definitely be cleaner...
func part2(input []string) string {
	// Same setup as Part 1 to first find the loop.
	area, start := parseInput(input)
	pipe, direction := findFirstConnectingPipe(area, start)
	loop := map[Position]bool{pipe: true}
	for pipe.val != 'S' {
		pipe, direction = findNextPipe(area, pipe, direction)
		loop[pipe] = true
	}

	pipesArea := area
	for y, row := range area {
		for x, val := range row {
			pipe := Position{
				x:   x,
				y:   y,
				val: val,
			}
			if loop[pipe] {
				pipesArea[y][x] = val
			} else {
				pipesArea[y][x] = '.'
			}
		}
	}

	pipesArea[start.y][start.x] = getStartPipeType(pipesArea, start)

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
	 * We can then flood fill tiles connected to borders with 0s.
	 * And finally workout which tiles are enclosed by finding only checking original tiles.
	 * (Every even Row & Even Column can be ignored as they are the expanded values).
	 * Count the remaining .s.
	 */
	expandedArea := expandArea(pipesArea)
	enclosedNodesArea := markUnenclosedNodes(expandedArea)
	count := 0
	for y := 0; y < len(enclosedNodesArea); y += 2 {
		for x := 0; x < len(enclosedNodesArea[y]); x += 2 {
			if enclosedNodesArea[y][x] == '.' {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}
func markUnenclosedNodes(area [][]rune) [][]rune {
	var toProcess []Position
	for x, val := range area[0] {
		tile := Position{
			x:   x,
			y:   0,
			val: val,
		}
		toProcess = append(toProcess, tile)
		y := len(area) - 1
		tile = Position{
			x:   x,
			y:   y,
			val: area[y][x],
		}
		toProcess = append(toProcess, tile)
	}

	for y, _ := range area {
		tile := Position{
			y:   y,
			x:   0,
			val: area[y][0],
		}
		toProcess = append(toProcess, tile)
		x := len(area[y]) - 1
		tile = Position{
			y:   y,
			x:   x,
			val: area[y][x],
		}
		toProcess = append(toProcess, tile)
	}
	for len(toProcess) > 0 {
		pos := Position{
			x:   toProcess[0].x,
			y:   toProcess[0].y,
			val: area[toProcess[0].y][toProcess[0].x],
		}
		toProcess = append(toProcess[:0], toProcess[1:]...)
		if pos.val == '.' {
			area[pos.y][pos.x] = '0'
			if pos.x < len(area[pos.y])-2 {
				right := Position{x: pos.x + 1, y: pos.y}
				if area[right.y][right.x] == '.' {
					toProcess = append(toProcess, right)
				}
			}
			if pos.x > 0 {
				left := Position{x: pos.x - 1, y: pos.y}
				if area[left.y][left.x] == '.' {
					toProcess = append(toProcess, left)
				}
			}
			if pos.y > 0 {
				above := Position{x: pos.x, y: pos.y - 1}
				if area[above.y][above.x] == '.' {
					toProcess = append(toProcess, above)
				}
			}
			if pos.y < len(area)-2 {
				below := Position{x: pos.x, y: pos.y + 1}
				if area[below.y][below.x] == '.' {
					toProcess = append(toProcess, below)
				}
			}
		}
	}

	return area
}

func expandArea(area [][]rune) [][]rune {
	var expandedArea [][]rune
	for y, row := range area {
		expandedRow := make([]rune, len(area[y])*2)
		insertedRow := make([]rune, len(area[y])*2)
		for x, val := range row {
			newX := x * 2
			// newY := y * 2
			expandedRow[newX] = val
			if val == '-' || val == 'F' || val == 'L' {
				expandedRow[newX+1] = '-'
			} else {
				expandedRow[newX+1] = '.'
			}

			if val == '|' || val == 'F' || val == '7' {
				insertedRow[newX] = '|'
			} else {
				insertedRow[newX] = '.'
			}
			insertedRow[newX+1] = '.'
		}
		expandedArea = append(expandedArea, expandedRow)
		expandedArea = append(expandedArea, insertedRow)
	}
	return expandedArea
}

func printArea(area [][]rune) {
	// FOR DEBUGGING PURPOSES
	fmt.Println()
	for y := range area {
		for _, val := range area[y] {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
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

func getStartPipeType(area [][]rune, start Position) rune {
	rightConnectors := map[rune]bool{
		'-': true,
		'J': true,
		'7': true,
	}
	leftConnectors := map[rune]bool{
		'-': true,
		'F': true,
		'L': true,
	}
	aboveConnectors := map[rune]bool{
		'|': true,
		'F': true,
		'7': true,
	}
	belowConnectors := map[rune]bool{
		'|': true,
		'J': true,
		'L': true,
	}

	var right bool
	var left bool
	var above bool
	var below bool
	if start.x < len(area[start.y])-1 {
		right = rightConnectors[area[start.y][start.x+1]]
	}
	if start.x > 0 {
		left = leftConnectors[area[start.y][start.x-1]]
	}
	if start.y > 0 {
		above = aboveConnectors[area[start.y-1][start.x]]
	}
	if start.y < len(area)-1 {
		below = belowConnectors[area[start.y+1][start.x]]
	}

	if right {
		if left {
			return '-'
		} else if above {
			return 'L'
		} else if below {
			return 'F'
		}
	} else if left {
		if above {
			return 'J'
		} else if below {
			return '7'
		}
	} else if above {
		if below {
			return '|'
		}
	}
	return '?'
}
