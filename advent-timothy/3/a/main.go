package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("3/a/input")
	scoreindexer, err := os.Open("3/a/scoreindex")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scoreindex := bufio.NewScanner(scoreindexer)

	var i = 1
	var a = 0
	var compartments [2]string

	var sameGoods string
	for scanner.Scan() {
		for index, v := range scanner.Text() {
			compartments[a] = compartments[a] + string(v)
			if len(scanner.Text())/2 == index+1 {
				i = 1
				a = 1
			}
			i++
		}
		a = 0

		var check = findSameGoods(compartments)

		if check != "" {
			sameGoods = sameGoods + check
		}

		compartments[0] = ""
		compartments[1] = ""
	}
	var score = 0
	for scoreindex.Scan() {
		for _, goods := range sameGoods {
			for index, value := range scoreindex.Text() {
				if string(value) == string(goods) {
					score += index + 1
				}
			}
		}
	}
	fmt.Println(score)
}

func findSameGoods(compartments [2]string) string {
	for _, value1 := range compartments[0] {
		for _, value2 := range compartments[1] {
			if value1 == value2 {
				return string(value1)
			}
		}
	}
	return ""
}
