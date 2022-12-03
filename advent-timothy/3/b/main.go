package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("3/b/input")
	scoreindexer, err := os.Open("3/b/scoreindex")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scoreindex := bufio.NewScanner(scoreindexer)

	var i = 0
	var compartments [3]string

	var sameGoods string
	var check string
	for scanner.Scan() {

		if i == 3 {
			i = 0
		}

		compartments[i] = scanner.Text()
		i++
		if i == 3 {
			check = checkBedge(compartments)
		}

		if check != "" && i == 3 {
			sameGoods = sameGoods + check
		}

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

func checkBedge(compartments [3]string) string {
	var bedge = ""
	for _, value1 := range compartments[0] {
		for _, value2 := range compartments[1] {
			if value1 == value2 {

				for _, value3 := range compartments[2] {
					if value3 == value2 {
						return string(value3)
					}
				}
			}
		}
	}
	return bedge
}
