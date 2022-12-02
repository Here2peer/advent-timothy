package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	elfs, err := os.Open("1/b/inventory")

	if err != nil {
		log.Fatal(err)
	}
	defer elfs.Close()

	scanner := bufio.NewScanner(elfs)

	var caloriesCurrent = 0
	var caloriesTop = [3]int{0, 0, 0}

	for scanner.Scan() {

		if len(scanner.Text()) == 0 {
			for index, entry := range caloriesTop {
				if caloriesCurrent > entry {
					caloriesTop[index] = caloriesCurrent
					break
				}
			}

			caloriesCurrent = 0
		}
		conv, _ := strconv.Atoi(scanner.Text())
		caloriesCurrent += conv
	}
	var totalCalories = 0
	for index, entry := range caloriesTop {
		totalCalories += entry
		fmt.Println("CalorieElf ", index+1, " to ask: ", entry)
	}
	fmt.Println("They are carrying: ", totalCalories, " calories.")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
