package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	input, err := os.Open("8/a/input")

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

	for rowIndex, row := range grid {
		print(rowIndex, "|")
		for colIndex, _ := range row {

			if isEdge(rowIndex, colIndex) {
				visible[rowIndex][colIndex] = true
			} else {
				if lookUp(rowIndex, colIndex, grid) {
					visible[rowIndex][colIndex] = true
				}
				if lookDown(rowIndex, colIndex, grid) {
					visible[rowIndex][colIndex] = true
				}
				if lookLeft(rowIndex, colIndex, grid) {
					visible[rowIndex][colIndex] = true
				}
				if lookRight(rowIndex, colIndex, grid) {
					visible[rowIndex][colIndex] = true
				}
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
}

func lookUp(row int, col int, grid [99][99]int) bool {
	for i := row - 1; i >= 0; i-- {
		if grid[i][col] >= grid[row][col] {
			return false
		}
	}
	return true
}

func lookDown(row int, col int, grid [99][99]int) bool {
	for i := row + 1; i <= 98; i++ {
		if grid[i][col] >= grid[row][col] {
			return false
		}
	}
	return true
}

func lookLeft(row int, col int, grid [99][99]int) bool {
	for i := col - 1; i >= 0; i-- {
		if grid[row][i] >= grid[row][col] {
			return false
		}
	}
	return true
}

func lookRight(row int, col int, grid [99][99]int) bool {
	for i := col + 1; i <= 98; i++ {
		if grid[row][i] >= grid[row][col] {
			return false
		}
	}
	return true
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
