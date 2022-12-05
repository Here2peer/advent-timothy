package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	crates, err := os.Open("5/b/crates")

	input, err := os.Open("5/b/input")

	if err != nil {
		log.Fatal(err)
	}
	defer crates.Close()

	scanner := bufio.NewScanner(crates)

	scanner.Split(bufio.ScanLines)

	moves := bufio.NewScanner(input)

	moves.Split(bufio.ScanLines)

	var emptyCount = 0
	var x = 0
	var y = 0
	var crateBool = false

	var cargoHold [10][999]string

	for scanner.Scan() {
		for _, value := range scanner.Text() {
			if crateBool == true {
				cargoHold[x][y] = string(value)
				crateBool = false
			}

			if value == 32 {
				emptyCount++
			} else {
				emptyCount = 0
			}
			if emptyCount == 4 {
				x++
				emptyCount = 0
			}
			if value == 91 {
				crateBool = true
			}
			if value == 93 {
				x++
			}

		}
		x = 0
		y++
	}
	cargoHold = flipArray(cargoHold)

	var moveBool = false
	var fromBool = false
	var toBool = false

	for moves.Scan() {
		var moveAmount = 0
		var from = 0
		var to = 0
		for _, value := range strings.Split(moves.Text(), " ") {
			if moveBool {
				moveBool = false
				var parsed, _ = strconv.Atoi(value)
				moveAmount = parsed
			}

			if fromBool {
				fromBool = false
				var parsed, _ = strconv.Atoi(value)
				from = from + (parsed - 1)
			}

			if toBool {
				toBool = false
				var parsed, _ = strconv.Atoi(value)
				to = to + (parsed - 1)
			}

			if value == "move" {
				moveBool = true
			}

			if value == "from" {
				fromBool = true
			}
			if value == "to" {
				toBool = true
			}
		}
		println("Move: ", moveAmount, "from: ", from, "to: ", to)
		cargoHold = doMove(moveAmount, from, to, cargoHold)
	}

	for index, cargo := range cargoHold {
		if cargoHold[index][0] != "" {
			var check = findFirst(cargoHold, index, 1)
			print(cargo[check])
		}
	}
}

func flipArray(array [10][999]string) [10][999]string {
	var outerArray [10][999]string
	var newY = 0
	for _, entry := range array {
		var innerArray [999]string
		var newX = 0
		for x := 10; x != 0; x-- {
			if entry[x-1] != "" {
				innerArray[newX] = entry[x-1]
				newX++
			}
		}
		outerArray[newY] = innerArray
		newY++
	}
	return outerArray
}

func doMove(amount int, from int, to int, cargoHold [10][999]string) [10][999]string {
	var index = findFirst(cargoHold, from, 1)
	if amount == 0 {
		return cargoHold
	}

	if amount == 1 {
		amount = amount - 1
		cargoHold = execute(index, to, from, cargoHold)
		println(1)
	} else {
		var inventory, cargo = crateMover9001(cargoHold[from], index, amount)
		cargoHold[from] = cargo
		if len(inventory) != 0 {
			cargoHold = addCrates(inventory, cargoHold, to)
		}
	}

	return cargoHold
}

func execute(index int, to int, from int, cargoHold [10][999]string) [10][999]string {
	var check = findFirst(cargoHold, to, 0)
	cargoHold[to][check] = cargoHold[from][index]
	cargoHold[from][index] = ""
	return cargoHold
}

func addCrates(inventory []string, cargoHold [10][999]string, to int) [10][999]string {
	var firstVal = findFirst(cargoHold, to, 0)
	for i := len(inventory); i != 0; i-- {
		cargoHold[to][firstVal] = inventory[i-1]
		firstVal++
	}
	return cargoHold
}

func crateMover9001(cargo [999]string, index int, amount int) ([]string, [999]string) {
	var inventory []string
	for i := index; i > index-amount; i-- {
		if i < 0 {
			break
		}
		inventory = append(inventory, cargo[i])
		cargo[i] = ""
	}
	return inventory, cargo
}

func findFirst(cargoHold [10][999]string, from int, offset int) int {
	for x := 0; x < len(cargoHold[from]); x++ {
		if cargoHold[from][x] == "" {
			return x - offset
		}
	}
	return 2000
}
