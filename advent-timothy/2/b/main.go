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
				if string(v) == "X" {
					myHand = lossHand(opponent)
				} else if string(v) == "Y" {
					myHand = drawHand(opponent)
				} else {
					myHand = winHand(opponent)
				}
				i = 0
			}
		}
		if len(opponent) != 0 && len(myHand) != 0 {
			if opponent == "A" {
				if myHand == "X" {
					score += draw
				} else if myHand == "Y" {
					score += win
				}
			}
			if opponent == "B" {
				if myHand == "Y" {
					score += draw
				} else if myHand == "Z" {
					score += win
				}
			}
			if opponent == "C" {
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

func winHand(opponent string) string {
	if opponent == "A" {
		return "Y"
	} else if opponent == "B" {
		return "Z"
	} else {
		return "X"
	}
}
func drawHand(opponent string) string {
	if opponent == "A" {
		return "X"
	} else if opponent == "B" {
		return "Y"
	} else {
		return "Z"
	}
}
func lossHand(opponent string) string {
	if opponent == "A" {
		return "Z"
	} else if opponent == "B" {
		return "X"
	} else {
		return "Y"
	}
}
