package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var win = 6
	var draw = 3

	input, err := os.Open("2/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var i = 0

	var opponent string
	var myHand string
	var score = 0

	for scanner.Scan() {
		for _, v := range scanner.Text() {

			if i == 0 {
				opponent = string(v)

				i++
			} else if string(v) != " " {
				myHand = string(v)
				i = 0
			}
		}
		if len(opponent) != 0 && len(myHand) != 0 {
			if opponent == "A" {
				fmt.Println(opponent, myHand)
				if myHand == "X" {
					score += draw
				} else if myHand == "Y" {
					score += win
				}
			}
			if opponent == "B" {
				fmt.Println(opponent, myHand)
				if myHand == "Y" {
					score += draw
				} else if myHand == "Z" {
					score += win
				}
			}
			if opponent == "C" {
				fmt.Println(opponent, myHand)
				if myHand == "Z" {
					score += draw
				} else if myHand == "X" {
					score += win
				}
			}
			if myHand == "X" {
				score += 1
			} else if myHand == "Y" {
				score += 2
			} else {
				score += 3
			}
		}
	}
	fmt.Println(score)

}
