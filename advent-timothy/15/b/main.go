package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

var pairs []interface{}

var abbysRightY = 0
var abbysLeftY = 1000000000000
var abbysX = 0

var cave map[image.Point]rune
var points []image.Point

func main() {
	input, _ := os.ReadFile("14/a/input")
	// Split input on newline
	lines := strings.Split(string(input), "\n")
	cave = map[image.Point]rune{}
	cave[image.Point{X: 0, Y: 500}] = '+'
	for _, line := range lines {
		points = []image.Point{}
		rocks := strings.Split(line, " -> ")
		for _, rock := range rocks {
			parsed := strings.Split(rock, ",")
			x, _ := strconv.Atoi(parsed[1])
			y, _ := strconv.Atoi(parsed[0])
			points = append(points, image.Point{X: x, Y: y})
		}
		for i := 0; i < len(points)-1; i++ {
			addRocks(points[i], points[i+1])
		}

	}
	var noVoid = true
	for noVoid {
		sand := image.Point{X: 0, Y: 500}
		moves := []image.Point{{1, 0}, {1, -1}, {1, 1}}
		index := 0
		for index < len(moves) && noVoid {
			var possible = true
			for possible {
				if sand.X == abbysX {
					noVoid = false
					break
				}
				sand = sand.Add(moves[index])
				if _, ok := cave[sand]; ok {
					possible = false
					sand = sand.Sub(moves[index])
					index++
					break
				}
				index = 0
			}
		}
		if noVoid {
			cave[sand] = '.'
		}
	}
	var sand = 0
	for i, v := range cave {
		if v == '.' {
			fmt.Printf("%+v\n", i)
			sand++
		}
	}
	println(sand)
}

func checkVoid(rock image.Point) {
	if rock.Y+1 > abbysRightY {
		abbysRightY = rock.Y + 1
	}
	if rock.Y < abbysLeftY {
		abbysLeftY = rock.Y - 1
	}
	if rock.X > abbysX {
		abbysX = rock.X
	}
}

func addRocks(from, to image.Point) {
	for y := min(from.Y, to.Y); y <= max(from.Y, to.Y); y++ {

		for x := min(from.X, to.X); x <= max(from.X, to.X); x++ {
			rock := image.Point{x, y}
			checkVoid(rock)
			cave[rock] = '#'
		}

	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
