package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day22/input")
	return part1(input), part2(input)
}

type Point struct {
	x, y, z int
}

type Brick struct {
	points  []Point
	axisDir string
}

func part1(input []string) string {
	// Load the bricks into a map based on layer [Layer/Z] = []Brick
	bricksByLayer := parseBricks(input)
	// Let the bricks fall down and update the map with the new positions of the bricks.
	bricksByLayer = dropBricks(bricksByLayer)
	fmt.Println(bricksByLayer)
	// Count how many bricks solely support other bricks - these bricks cannot be deleted.
	// Result = Total Bricks - Number of bricks that cannot be deleted.
	return fmt.Sprint(len(bricksByLayer))
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

func dropBricks(bricksByLayer map[int][]Brick) map[int][]Brick {
	totalLayers := len(bricksByLayer) - 1
	layersProcessed := 0

	droppedBricksByLayer := make(map[int][]Brick, len(bricksByLayer))
	droppedBricksByLayer[1] = bricksByLayer[1]
	// Start at layer 2 - as bricks at layer 1 are already as low as they can go.
	for layer := 2; layersProcessed < totalLayers; layer++ {
		bricksAtLayer := bricksByLayer[layer]
		fmt.Println(layer, bricksAtLayer, layersProcessed, totalLayers)
		if len(bricksAtLayer) > 0 {
			layersProcessed++
			for _, brick := range bricksAtLayer {
				newBrick := dropBrick(brick, droppedBricksByLayer)
				newBrickLayer := newBrick.points[0].z
				droppedBricksByLayer[newBrickLayer] = append(droppedBricksByLayer[newBrickLayer], newBrick)
			}
		}
	}

	return droppedBricksByLayer
}

// Check the layers below the brick for any bricks already occupying a point in the brick's path.
// Once a layer is found with the collision - the new brick is the same but z = collision layer + 1
func dropBrick(brick Brick, bricksByLayer map[int][]Brick) Brick {
	// Work down each layer until a collision is found.
	newLayer := brick.points[0].z
	for currLayer := newLayer; currLayer > 0; currLayer-- {
		newLayer = currLayer - 1
		collision := isCollisionAtLayer(brick, bricksByLayer[newLayer])
		if collision {
			break
		}
	}
	return updateBrickLayer(brick, newLayer+1)
}

func updateBrickLayer(brick Brick, newLayer int) Brick {
	newBrick := brick
	layerDiff := newBrick.points[0].z - newLayer
	for i, _ := range newBrick.points {
		newBrick.points[i].z -= layerDiff
	}
	return newBrick
}

func isCollisionAtLayer(brick Brick, bricksAtLayer []Brick) bool {
	// At each layer check all the bricks at that layer.
	for _, brickAtLayer := range bricksAtLayer {
		// Compare the x & Y coords of the brick falling and the layer's bricks coords.
		for _, point := range brick.points {
			for _, brickAtNewLayerPoints := range brickAtLayer.points {
				if point.x == brickAtNewLayerPoints.x && point.y == brickAtNewLayerPoints.y {
					return true
				}
			}
		}
	}
	return false
}

func parseBricks(in []string) map[int][]Brick {
	bricksByLayer := make(map[int][]Brick)
	for _, line := range in {
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
		}
		brickPoints[len(brickPoints)-1] = point2

		bricksByLayer[point1.z] = append(bricksByLayer[point1.z], Brick{points: brickPoints, axisDir: axisDirection})
	}
	return bricksByLayer
}

func toInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
