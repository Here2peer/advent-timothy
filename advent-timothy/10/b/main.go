package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var addxRegister []int
var screen [6][40]string
var sum = 0
var yAxis = 0
var cycle = 0
var counter = 0

func main() {

	input, err := os.Open("10/b/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	var register = 1

	addxRegister = append(addxRegister, 1)
	for scanner.Scan() {
		if scanner.Text() == "addx" {
			scanner.Scan()
			var value, _ = strconv.Atoi(scanner.Text())
			register = addxV(register, value)
			addxRegister = append(addxRegister, value)
		} else {
			cycle++
			sum += checkSignal(register)
		}
	}

	drawScreen()
	println(cycle, sum)
}

func drawScreen() {
	for _, inner := range screen {
		for _, value := range inner {
			print(value)
		}
		println("")
	}
}

func addxV(register int, value int) int {
	for i := 1; i <= 2; i++ {
		cycle++
		sum += checkSignal(register)
	}
	return register + value
}

func checkSignal(register int) int {
	var multi = 0
	draw(register)
	if cycle == 20 || (cycle > 20 && (cycle-20)%40 == 0) {
		for _, value := range addxRegister {
			multi = multi + value
		}
		return cycle * multi
	}
	return 0
}

func draw(register int) {
	println("register ", register, counter, cycle-1)

	if counter%40 == 0 && counter != 0 {
		yAxis++
		counter = 0
	}

	if screen[yAxis][counter] == "" {
		if register+1 == counter || register-1 == counter || register == counter {
			screen[yAxis][counter] = "#"
		} else {
			screen[yAxis][counter] = "."
		}
	}
	counter++
}
