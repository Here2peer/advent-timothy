package main

import (
	"bufio"
	orderedmap2 "github.com/elliotchance/orderedmap/v2"
	"image"
	"log"
	"os"
)

var world [41][]int
var nodes *orderedmap2.OrderedMap[int, []int]

func main() {

	var currentNode *orderedmap2.OrderedMap[int, []int]
	input, err := os.Open("12/a/input")

	img := map[image.Point]rune{}

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	var y = 0
	var startY = 0
	var startX = 0
	//var start, end image.Point
	for scanner.Scan() {
		for i, value := range scanner.Text() {
			if value == 83 {
				startX = i
				startY = y
				//start = image.Point{y, i}
			} else if value == 'E' {
				//end = image.Point{y, i}
			}
			img[image.Point{y, i}] = value
			world[y] = append(world[y], int(value))
			if string(value) == "z" {
				println(value)
			}
		}
		y++
	}

	var count = 0
	var moves = generateMoves(startY, startX)
	currentNode = orderedmap2.NewOrderedMap[int, []int]()
	currentNode.Set(count, moves)
	println(aStar(startY, startX, currentNode, count))
	println(nodes.Len() + 1)

}

func aStar(fromY int, fromX int, currentNode1 *orderedmap2.OrderedMap[int, []int], count int) bool {
	//println(fromY, fromX, world[fromY][fromX])
	count++
	var test = string(world[fromY][fromX])
	var currentNode = currentNode1
	if test == "E" {
		nodes = currentNode
		return true
	}
	for _, moves := range findMoves(fromY, fromX) {
		if checkClimbable(world[moves[0]][moves[1]], fromY, fromX) {
			var visited = false
			for el := currentNode.Back(); el != nil; el = el.Prev() {
				if el.Value[0] == moves[0] && el.Value[1] == moves[1] {
					visited = true
					break
				}
			}
			if !visited {
				currentNode.Set(count, moves)
				if aStar(moves[0], moves[1], currentNode1, count) {
					return true
				}
			} else {
				currentNode.Delete(currentNode.Back().Key)
			}
		}
	}
	return false
}
func pop(i int, xs []int) []int {
	ys := append(xs[:i], xs[i+1:]...)
	return ys
}

func checkClimbable(value int, y int, x int) bool {
	if (value == world[y][x] || value == world[y][x]+1) || world[y][x] == 83 || (value == 69 && world[y][x] == 122) {
		return true
	}
	return false
}

func findMoves(y int, x int) map[string][]int {
	var moves = make(map[string][]int)
	if y != 0 {
		moves["up"] = generateMoves(y-1, x)
	}
	if y != len(world)-1 {
		moves["down"] = generateMoves(y+1, x)
	}
	if x != 0 {
		moves["left"] = generateMoves(y, x-1)
	}

	if x != len(world[y])-1 {
		moves["right"] = generateMoves(y, x+1)
	}
	return moves
}

func generateMoves(currentY int, currentX int) []int {
	var move []int
	move = append(move, currentY)
	move = append(move, currentX)
	return move
}
