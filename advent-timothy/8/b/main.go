package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	input, err := os.Open("8/b/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	var y = 0
	var x = 0

	var grid [99][99]int
	var visible [99][99]bool

	for scanner.Scan() {
		for _, value := range scanner.Text() {
			grid[y][x] = int(value - '0')
			x++
		}
		y++
		x = 0
	}

	var scenic = 0

	for rowIndex, row := range grid {
		print(rowIndex, "|")
		for colIndex, _ := range row {
			var score = 0

			if isEdge(rowIndex, colIndex) {
				visible[rowIndex][colIndex] = true
			} else {

				var up, score1 = lookUp(rowIndex, colIndex, grid)
				var down, score2 = lookDown(rowIndex, colIndex, grid)
				var left, score3 = lookLeft(rowIndex, colIndex, grid)
				var right, score98 = lookRight(rowIndex, colIndex, grid)

				score = (score1 * score2) * (score3 * score98)

				if up {
					visible[rowIndex][colIndex] = true
				}
				if down {
					visible[rowIndex][colIndex] = true
				}
				if left {
					visible[rowIndex][colIndex] = true
				}
				if right {
					visible[rowIndex][colIndex] = true
				}

				if colIndex == 2 && rowIndex == 3 {
					println("")
					println("score: ", score1, score2, score3, score98)
				}
			}
			if score > scenic {
				scenic = score
			}

			print(row[colIndex])
		}
		println("")
	}
	var count = 0
	println("")
	for rowIndex, row := range visible {
		print(rowIndex, "|")
		for colIndex, _ := range row {
			if row[colIndex] {
				print("1")
			} else {
				print("0")
			}
			print(" ")
			if row[colIndex] {
				count++
			}
		}
		println("")
	}
	println("Trees: ", count)
	println("Scenic score: ", scenic)
}

func lookUp(row int, col int, grid [99][99]int) (bool, int) {
	var count = 0
	for i := row - 1; i >= 0; i-- {
		count++
		if grid[i][col] >= grid[row][col] {
			return false, count
		}
	}
	return true, count
}

func lookDown(row int, col int, grid [99][99]int) (bool, int) {
	var count = 0
	for i := row + 1; i <= 98; i++ {
		count++
		if grid[i][col] >= grid[row][col] {
			if col == 2 && row == 3 {
				println("row: ", row, i)
			}
			return false, count
		}
	}
	return true, count
}

func lookLeft(row int, col int, grid [99][99]int) (bool, int) {
	var count = 0
	for i := col - 1; i >= 0; i-- {
		count++
		if grid[row][i] >= grid[row][col] {
			return false, count
		}
	}
	return true, count
}

func lookRight(row int, col int, grid [99][99]int) (bool, int) {
	var count = 0
	for i := col + 1; i <= 98; i++ {
		count++
		if grid[row][i] >= grid[row][col] {
			return false, count
		}
	}
	return true, count
}

func isEdge(row int, col int) bool {
	if row == 0 || col == 0 {
		return true
	}
	if col == 98 || row == 98 {
		return true
	}
	return false
}
