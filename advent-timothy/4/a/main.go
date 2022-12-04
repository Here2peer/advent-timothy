package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("4/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var pair [2][2]int
	var reconsider = 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		for index1, val1 := range strings.Split(scanner.Text(), ",") {
			for index2, val2 := range strings.Split(val1, "-") {
				pair[index1][index2], _ = strconv.Atoi(val2)
			}
		}
		if checkPair(getZones(pair[0]), pair[1]) {
			reconsider++
		} else if checkPair(getZones(pair[1]), pair[0]) {
			reconsider++
		}
	}
	fmt.Println("score = ", reconsider)
}

func checkPair(zones []int, checkPair [2]int) bool {
	var zone2 = getZones(checkPair)
	var hits []bool
	for _, zone1 := range zones {
		var found = false
		for _, zones2 := range zone2 {
			if zone1 == zones2 {
				hits = append(hits, true)
				found = true
				break
			}
		}
		if !found {
			hits = append(hits, false)
		}
	}
	for _, bools := range hits {
		if !bools {
			return false
		}
	}
	return true
}

func getZones(pair [2]int) []int {
	var zones []int

	for x := pair[0]; x <= pair[1]; x++ {
		zones = append(zones, x)
	}
	return zones
}
