package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("7/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)
	var path []string
	values := make(map[string]int)
	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text())
		if string(scanner.Text()[0]) == "$" {
			if cmd[1] == "cd" {
				dir := cmd[2]
				if cmd[2] == "/" {
					//directories.Set("/")
					path = appendArray(path, cmd[2])
				} else if cmd[2] == ".." {
					path = popArray(path)
				} else {
					if path[len(path)-1] != "/" {
						path = appendArray(path, path[len(path)-1]+"/"+dir)
					} else {
						path = appendArray(path, path[len(path)-1]+""+dir)
					}
				}
			}
		}
		if string(cmd[0]) != "$" && string(cmd[0]) != "dir" {
			for _, dir := range path {
				var value, _ = strconv.Atoi(cmd[0])
				values[dir] += value
			}

		}
	}
	for _, value := range path {
		println(value)
	}

	var totala = 0
	var unused = 30000000
	var curSpace = 70000000
	var small = 0
	for _, value := range values {
		if value <= 100000 {
			totala += value
		}
		if value >= unused-(curSpace-values["/"]) {
			if small == 0 {
				small = value
			} else if value < small {
				small = value
			}
		}
	}
	println(totala)

}

func appendArray(array []string, value string) []string {
	return append(array, value)
}

func popArray(array []string) []string {
	var _, a = array[len(array)-1], array[:len(array)-1]
	return a
}
