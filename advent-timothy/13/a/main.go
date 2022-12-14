package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var pairs []interface{}

func main() {
	input, _ := os.ReadFile("13/a/input")
	// Split input on newline
	lines := strings.Split(string(input), "\n\n")
	var count = 0
	for index, line := range lines {

		// Append pair of arrays to slice of pairs
		if line != "" {
			pairs := strings.Split(line, "\n")
			var left = unString(pairs[0])
			var right = unString(pairs[1])
			//fmt.Printf("%+v\n", pairs[0])
			//fmt.Printf("%+v\n", pairs[1])
			if compareValues(left, right) {
				count += index + 1
				println(index)
			}
		}
	}
	println(count)
}

func compareValues(a, b []interface{}) bool {
	for i := 0; i < len(a); i++ {
		switch a[i].(type) {
		case float64:
			if len(b) == i {
				return false
			}
			switch b[i].(type) {
			case float64:
				if a[i].(float64) < b[i].(float64) {
					return true
				} else if a[i].(float64) > b[i].(float64) {
					return false
				}
			case []interface{}:
				if compareValues(a, b[i].([]interface{})) {
					return true
				} else {
					continue
				}
			}
		case []interface{}:
			if len(a) > 0 && len(b) == 0 {
				return false
			}
			if _, ok := b[i].(float64); ok {
				if compareValues(a[i].([]interface{}), b) {
					return true
				} else {
					return false
				}
			}
			if compareValues(a[i].([]interface{}), b[i].([]interface{})) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func compareValues1(a, b []interface{}) bool {
	for i := 0; i < len(a); i++ {
		for x := 0; x < len(b); x++ {

			switch a[i].(type) {
			case float64:

				return true
			case []interface{}:

				return false
			}
		}
	}
	return false
}

func check(a, b any) bool {
	if valueB, ok := b.(float64); ok {
		if a.(float64) < valueB {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func unString(in string) []interface{} {
	var str []interface{}
	err := json.Unmarshal([]byte(in), &str)
	if err != nil {
		println("")
		fmt.Println(err)
	}
	return str
}
