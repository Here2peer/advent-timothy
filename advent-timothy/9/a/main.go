package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	input, err := os.Open("9/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var grid [1000][1000]string
	var visited [1000][1000]bool
	//grid[0][98] = "H"
	visited[100][100] = true
	grid[100][100] = "H"
	//var tailX = 0
	//var tailY = 98
	var frontX = 100
	var frontY = 100
	var init = false

	for scanner.Scan() {
		var v, _ = strconv.Atoi(string(scanner.Text()[2]))
		if scanner.Text()[0] == 'R' {
			var g, checked, x, y = right(v, frontX, frontY, grid, visited)
			grid = g
			visited = checked
			frontX = x
			frontY = y
			println("After right: ", y, x)
		}
		if scanner.Text()[0] == 'L' {
			var g, checked, x, y = left(v, frontX, frontY, grid, visited, init)
			grid = g
			visited = checked
			frontX = x
			frontY = y
			init = true
			println("After left: ", y, x)
		}
		if scanner.Text()[0] == 'U' {
			var g, checked, x, y = up(v, frontX, frontY, grid, visited)
			grid = g
			visited = checked
			frontX = x
			frontY = y

			println("After up: ", y, x)
		}
		if scanner.Text()[0] == 'D' {
			var g, checked, x, y = down(v, frontX, frontY, grid, visited)
			grid = g
			visited = checked
			frontX = x
			frontY = y
			println("After down: ", y, x)
		}
		var count = 0
		for testy, inner := range grid {
			for testx, value := range inner {
				if value == "T" {
					count++
					println("Position of T: ", testy, testx)
				}
			}
		}
		if count > 1 {
			println(string(scanner.Text()[0]))
			os.Exit(1)
		}
	}
	var score = 0
	for _, inner := range visited {
		for _, value := range inner {
			if value {
				score++
			}
		}
	}
	println("Score: ", score)

}

func findT(grid [1000][1000]string) bool {

	for _, inner := range grid {
		for _, value := range inner {
			if value == "T" {
				return true
			}
		}
	}
	return false
}

func checkTail(x int, y int, grid [1000][1000]string) (bool, int, int) {
	if grid[y][x-1] == "T" {
		return true, y, x - 1
	}
	if grid[y][x+1] == "T" {
		return true, y, x + 1
	}
	if grid[y-1][x] == "T" {
		return true, y - 1, x
	}
	if grid[y+1][x] == "T" {
		return true, y + 1, x
	}
	if grid[y+1][x+1] == "T" {
		return true, y + 1, x + 1
	}
	if grid[y-1][x-1] == "T" {
		return true, y - 1, x - 1
	}
	if grid[y+1][x-1] == "T" {
		return true, y + 1, x - 1
	}
	if grid[y-1][x+1] == "T" {
		return true, y - 1, x + 1
	}
	return false, 1000, 1000
}

func right(amount int, x int, y int, grid [1000][1000]string, visited [1000][1000]bool) ([1000][1000]string, [1000][1000]bool, int, int) {
	//var isSet = false
	var overlap = !findT(grid)
	for i := amount; i > 0; i-- {
		if overlap {
			grid[y][x] = "T"
			visited[y][x] = true
		} else {
			grid[y][x] = ""
		}
		x = x + 1
		if grid[y][x] == "T" {
			overlap = true
		}
		grid[y][x] = "H"

		var isTail, _, _ = checkTail(x, y, grid)
		if !isTail || !overlap {
			var _, tailY, tailX = checkTail(x-1, y, grid)
			grid[y][x-1] = "T"
			visited[y][x-1] = true
			grid[tailY][tailX] = ""
		}

		overlap = !findT(grid)
	}
	return grid, visited, x, y
}
func left(amount int, x int, y int, grid [1000][1000]string, visited [1000][1000]bool, init bool) ([1000][1000]string, [1000][1000]bool, int, int) {

	var isTail = false
	var tailY = 0
	var tailX = 0
	for i := amount; i > 0; i-- {
		var overlap = !findT(grid)
		if overlap && init {
			grid[y][x] = "T"
		} else {
			grid[y][x] = ""
		}
		x = x - 1
		if grid[y][x] == "T" {
			overlap = true
		}
		grid[y][x] = "H"
		if !init {
			grid[100][100] = "T"
			isTail, _, _ = checkTail(x, y, grid)
		} else {
			isTail, _, _ = checkTail(x, y, grid)
		}
		if !isTail || !overlap {
			_, tailY, tailX = checkTail(x+1, y, grid)
			grid[y][x+1] = "T"
			visited[y][x+1] = true
			grid[tailY][tailX] = ""
		}
	}
	return grid, visited, x, y
}
func down(amount int, x int, y int, grid [1000][1000]string, visited [1000][1000]bool) ([1000][1000]string, [1000][1000]bool, int, int) {
	//var _, tailY, tailX = checkTail(x, y, grid)
	var overlap = !findT(grid)
	for i := amount; i > 0; i-- {

		if overlap {
			grid[y][x] = "T"
		} else {
			grid[y][x] = ""
		}
		y = y + 1
		if grid[y][x] == "T" {
			overlap = true
		}
		grid[y][x] = "H"

		var isTail, _, _ = checkTail(x, y, grid)
		if !isTail || !overlap {
			var _, tailY, tailX = checkTail(x, y-1, grid)
			grid[y-1][x] = "T"
			visited[y-1][x] = true
			grid[tailY][tailX] = ""
		}

		overlap = !findT(grid)
	}
	return grid, visited, x, y
}
func up(amount int, x int, y int, grid [1000][1000]string, visited [1000][1000]bool) ([1000][1000]string, [1000][1000]bool, int, int) {
	var overlap = !findT(grid)
	for i := amount; i > 0; i-- {
		if overlap {
			grid[y][x] = "T"
		} else {
			grid[y][x] = ""
		}
		y = y - 1
		if grid[y][x] == "T" {
			overlap = true
		}
		grid[y][x] = "H"

		var isTail, _, _ = checkTail(x, y, grid)
		if !isTail || !overlap {
			var _, tailY, tailX = checkTail(x, y+1, grid)
			grid[y+1][x] = "T"
			visited[y+1][x] = true
			grid[tailY][tailX] = ""
		}
		overlap = !findT(grid)
	}
	return grid, visited, x, y
}
