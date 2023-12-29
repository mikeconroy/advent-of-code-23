package day16

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day16/input")
	return part1(input), part2(input)
}

/*
 * . Empty Space
 * / \ Mirrors
 * | - Splitters
 * Beam enters in the top-left corner heading to the right.
 * Empty Space - Continue
 * Mirror - Reflected 90 degrees depending on the angle of the mirror.
 *			Right-ward moving beam encountering a / would move upwards.
 * 			Right-ward moving beam encountering a \ would move downards.
 * Splitter - If the beam hits a pointy end then it continues as if it were empty space.
 *			  Otherwise the beam splits into 2 beams going in the directions of the pointy ends.
 * Beams do not interact with other beams. A tile can have multiple beams passing through it.
 */
func part1(input []string) string {
	tiles := loadTiles(input)
	energizedTiles := make(map[Point]int)
	beams := []Beam{
		{
			pos: Point{x: 0, y: 0},
			dir: RIGHT,
		},
	}

	// Cache of beams already processed to stop when a loop is found.
	cache := make(map[Beam]bool)
	for len(beams) > 0 {
		beam := beams[0]

		canMove := true
		for canMove {
			if cache[beam] {
				break
			}
			cache[beam] = true
			if beam.pos.y < 0 || beam.pos.x < 0 || beam.pos.y >= len(tiles) || beam.pos.x >= len(tiles[beam.pos.y]) {
				canMove = false
				break
			}

			currTile := tiles[beam.pos.y][beam.pos.x]
			energizedTiles[beam.pos] = energizedTiles[beam.pos] + 1
			if currTile == '-' && (beam.dir == UP || beam.dir == DOWN) {
				// Split Beam into horizontals
				leftBeam := Beam{
					dir: LEFT,
					pos: Point{
						x: beam.pos.x - 1,
						y: beam.pos.y,
					},
				}
				beams = append(beams, leftBeam)
				rightBeam := Beam{
					dir: RIGHT,
					pos: Point{
						x: beam.pos.x + 1,
						y: beam.pos.y,
					},
				}
				beams = append(beams, rightBeam)
				canMove = false
				break
			} else if currTile == '|' && (beam.dir == RIGHT || beam.dir == LEFT) {
				// Split Beam into verticals
				upBeam := Beam{
					dir: UP,
					pos: Point{
						x: beam.pos.x,
						y: beam.pos.y - 1,
					},
				}
				beams = append(beams, upBeam)
				downBeam := Beam{
					dir: DOWN,
					pos: Point{
						x: beam.pos.x,
						y: beam.pos.y + 1,
					},
				}
				beams = append(beams, downBeam)
				canMove = false
				break
			} else if currTile == '/' {
				if beam.dir == RIGHT {
					beam.dir = UP
				} else if beam.dir == DOWN {
					beam.dir = LEFT
				} else if beam.dir == UP {
					beam.dir = RIGHT
				} else if beam.dir == LEFT {
					beam.dir = DOWN
				}
			} else if currTile == '\\' {
				if beam.dir == RIGHT {
					beam.dir = DOWN
				} else if beam.dir == DOWN {
					beam.dir = RIGHT
				} else if beam.dir == UP {
					beam.dir = LEFT
				} else if beam.dir == LEFT {
					beam.dir = UP
				}
			}
			beam.move()
		}

		beams = beams[1:]
	}

	return fmt.Sprint(len(energizedTiles))
}

func (b *Beam) move() {
	if b.dir == UP {
		b.pos.y -= 1
	} else if b.dir == DOWN {
		b.pos.y += 1
	} else if b.dir == RIGHT {
		b.pos.x += 1
	} else if b.dir == LEFT {
		b.pos.x -= 1
	}
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Point struct {
	x int
	y int
}

type Beam struct {
	pos Point
	dir Direction
}
type Direction int

const (
	UP = iota
	DOWN
	RIGHT
	LEFT
)

func printTiles(tiles [][]rune) {
	fmt.Println()
	for _, row := range tiles {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
	fmt.Println()
}

func loadTiles(in []string) (tiles [][]rune) {
	for _, row := range in {
		newRow := make([]rune, len(row))
		for x, val := range row {
			newRow[x] = val
		}
		tiles = append(tiles, newRow)
	}
	return tiles
}

func printEnergizedBeams(tiles map[Point]int) {
	maxX, maxY := 0, 0
	for keys, _ := range tiles {
		if keys.x > maxX {
			maxX = keys.x
		}
		if keys.y > maxY {
			maxY = keys.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			val := tiles[Point{x: x, y: y}]
			if val == 0 {
				fmt.Print(".")
			} else if val < 10 {
				fmt.Print(val)
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
