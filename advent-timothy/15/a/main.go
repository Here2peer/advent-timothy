package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var world map[image.Point]rune
var count = 0

func main() {
	input, _ := os.ReadFile("15/a/input")
	// Split input on newline
	lines := strings.Split(string(input), "\n")
	world = map[image.Point]rune{}

	re := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])
		world[image.Point{x1, y1}] = 's'
		world[image.Point{x2, y2}] = 'b'
	}
	row := 2000000
	for sensor, v := range world {
		if v == 's' {
			point := sensor
			closest := closestPoint(point, 'b')
			setBeaconlessLocations(point, closest, row)
		}
	}
	printMap(world, row, false, false)
	println(count)
}

func manhattanDistance(p1, p2 image.Point) int {
	return int(math.Abs(float64(p1.X-p2.X))) + int(math.Abs(float64(p1.Y-p2.Y)))
}

func setBeaconlessLocations(p1, p2 image.Point, row int) {
	maxDist := manhattanDistance(p1, p2) + 1
	minY := p1.Y - maxDist
	maxY := p1.Y + maxDist
	minX := p1.X - maxDist
	maxX := p1.X + maxDist

	for y := minY; y <= maxY; y++ {
		if y == row {
			for x := minX; x <= maxX; x++ {
				if _, ok := world[image.Point{x, y}]; ok {

				} else {
					if manhattanDistance(image.Point{x, y}, p1) < maxDist {
						world[image.Point{x, y}] = '#'
					}
				}
			}
		}
	}
}

func closestPoint(p image.Point, object rune) image.Point {
	var closest image.Point
	var minDistance int

	for point, v := range world {
		distance := manhattanDistance(point, p)
		if minDistance == 0 || distance < minDistance {
			if v == object {
				closest = point
				minDistance = distance
			}
		}
	}
	return closest
}

func printMap(points map[image.Point]rune, row int, fullMap bool, print bool) {
	var minX, minY, maxX, maxY int
	minX = math.MaxInt32
	minY = math.MaxInt32

	for point, _ := range points {
		if minX == 0 || point.X < minX {
			minX = point.X
		}
		if minY == 0 || point.Y < minY {
			minY = point.Y
		}
		if maxX == 0 || point.X > maxX {
			maxX = point.X
		}
		if maxY == 0 || point.Y > maxY {
			maxY = point.Y
		}
	}
	if !fullMap {
		minY = row - 1
		maxY = row + 1
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if point, ok := points[image.Point{x, y}]; ok {
				if points[image.Point{x, y}] == '#' && y == row {
					count++
				}
				if print {
					fmt.Print(string(point))
				}
			} else {
				if print {
					fmt.Print(".")
				}
			}
		}
		if print {
			fmt.Println()
		}
	}
}
