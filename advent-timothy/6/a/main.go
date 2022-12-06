package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	input, err := os.Open("6/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)
	var y = 3
	for scanner.Scan() {
		for i := 0; y < len(scanner.Text()); i++ {
			var a []int
			for x := i; x <= y; x++ {
				a = append(a, int(scanner.Text()[x]))
			}
			if !duplicateInArray(a) {
				println(string(a[0]), string(a[1]), string(a[2]), string(a[3]), " index: ", y+1)
				return
			}
			y++
		}

	}

}

func duplicateInArray(arr []int) bool {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
