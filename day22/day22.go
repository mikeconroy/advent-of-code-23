package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
	"golang.org/x/exp/maps"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day22/input")
	return part1(input), part2(input)
}

type Point struct {
	x, y, z int
}

type Brick struct {
	id            int
	points        []Point
	axisDir       string
	isSupportedBy map[int]bool
}

func part1(input []string) string {
	bricks, grid := parseBricks(input)
	// Let the bricks fall down and update the map with the new positions of the bricks.
	bricks, grid, _ = dropBricks(bricks, grid)
	// Count how many bricks solely support other bricks - these bricks cannot be deleted.
	// Result = Total Bricks - Number of bricks that cannot be deleted.
	bricksToNotBeDisintegrated := calculateBricksNotToDisintegrate(bricks, grid)
	result := len(bricks) - len(bricksToNotBeDisintegrated)
	return fmt.Sprint(result)
}

func calculateBricksNotToDisintegrate(bricks map[int]Brick, grid map[Point]int) map[int]bool {
	importantBricks := make(map[int]bool)
	for point, id := range grid {
		if point.z == 1 {
			continue
		}
		currBrick := bricks[id]
		pointBelow := point
		pointBelow.z = point.z - 1
		idBelow := grid[pointBelow]
		if idBelow != 0 && idBelow != currBrick.id {
			currBrick.isSupportedBy[idBelow] = true
		}
	}
	for _, brick := range bricks {
		if len(brick.isSupportedBy) == 1 {
			importantBricks[maps.Keys(brick.isSupportedBy)[0]] = true
		}
	}
	return importantBricks
}

func part2(input []string) string {
	bricks, grid := parseBricks(input)
	bricks, grid, _ = dropBricks(bricks, grid)
	bricksToNotBeDisintegrated := calculateBricksNotToDisintegrate(bricks, grid)

	totalBricksMoved := 0
	for disintegrateId := range bricksToNotBeDisintegrated {
		bricksDropped := make(map[int]bool)
		bricksDropped[disintegrateId] = true
		bricksStillDropping := true
		for bricksStillDropping {
			bricksStillDropping = false
			for id, brick := range bricks {
				// Don't check ground layer. Don't check bricks that have already been marked as dropped.
				if brick.points[0].z == 1 || bricksDropped[id] {
					continue
				}
				brickIsSupported := false
				for supportedById := range brick.isSupportedBy {
					if !bricksDropped[supportedById] {
						brickIsSupported = true
						break
					}
				}
				if !brickIsSupported {
					bricksDropped[id] = true
					bricksStillDropping = true
				}
			}
		}
		totalBricksMoved += len(bricksDropped) - 1
	}

	return fmt.Sprint(totalBricksMoved)
}

func dropBricks(bricks map[int]Brick, grid map[Point]int) (map[int]Brick, map[Point]int, int) {
	brickMoved := true
	bricksMoved := make(map[int]bool)

	// Loop until all bricks have dropped as far as possible
	for brickMoved {
		brickMoved = false

		for _, currBrick := range bricks {
			// Brick is already on the ground so can't fall futher
			if currBrick.points[0].z == 1 {
				continue
			}

			currLayer := currBrick.points[0].z
			newLayer := currLayer
			keepDropping := true
			for keepDropping && newLayer > 1 {
				newLayer--
				for _, point := range currBrick.points {
					point.z = newLayer
					if grid[point] != 0 {
						keepDropping = false
						newLayer++
						break
					}
				}
			}
			if newLayer != currLayer {
				brickMoved = true
				newBrick := currBrick
				// Update the grid to remove the old points
				for _, point := range currBrick.points {
					delete(grid, point)
				}
				layerDiff := newBrick.points[0].z - newLayer
				for i, _ := range newBrick.points {
					newBrick.points[i].z -= layerDiff

					grid[newBrick.points[i]] = newBrick.id
				}
				bricks[newBrick.id] = newBrick
				bricksMoved[newBrick.id] = true
			}

		}
	}
	return bricks, grid, len(bricksMoved)
}

func parseBricks(in []string) (map[int]Brick, map[Point]int) {
	bricksMap := make(map[int]Brick)
	grid := make(map[Point]int)
	for i, line := range in {
		id := i + 1
		ends := strings.Split(line, "~")
		end1 := strings.Split(ends[0], ",")
		end2 := strings.Split(ends[1], ",")
		point1 := Point{toInt(end1[0]), toInt(end1[1]), toInt(end1[2])}
		point2 := Point{toInt(end2[0]), toInt(end2[1]), toInt(end2[2])}

		var axisDirection string
		var start, end int

		if point1.x != point2.x {
			axisDirection = "x"
			start = point1.x
			end = point2.x
		} else if point1.y != point2.y {
			axisDirection = "y"
			start = point1.y
			end = point2.y
		} else if point1.z != point2.z {
			axisDirection = "z"
			start = point1.z
			end = point2.z
		}

		grid[point1] = id
		grid[point2] = id

		currentPoint := point1
		brickPoints := make([]Point, end-start+1)
		brickPoints[0] = point1
		for j := start + 1; j < end; j++ {
			if axisDirection == "x" {
				currentPoint.x += 1
			} else if axisDirection == "y" {
				currentPoint.y += 1
			} else if axisDirection == "z" {
				currentPoint.z += 1
			}
			brickPoints[j-start] = currentPoint
			grid[currentPoint] = id
		}
		brickPoints[len(brickPoints)-1] = point2
		newBrick := Brick{id: id, points: brickPoints, axisDir: axisDirection, isSupportedBy: make(map[int]bool)}
		bricksMap[newBrick.id] = newBrick

	}
	return bricksMap, grid
}

func toInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
