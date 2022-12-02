package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	elfs, err := os.Open("1/a/inventory")

	if err != nil {
		log.Fatal(err)
	}
	defer elfs.Close()

	scanner := bufio.NewScanner(elfs)

	var caloriesCurrent = 0
	var caloriesHighest = 0

	for scanner.Scan() {

		if len(scanner.Text()) == 0 {
			if caloriesCurrent > caloriesHighest {
				caloriesHighest = caloriesCurrent
			}
			caloriesCurrent = 0
		}
		conv, _ := strconv.Atoi(scanner.Text())
		caloriesCurrent += conv
	}

	fmt.Println("CalorieElf to ask: ", caloriesHighest)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
