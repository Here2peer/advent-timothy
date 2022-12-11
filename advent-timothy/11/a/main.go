package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var monkeys []Monkey
var monkeyBus = [8]int{0, 0, 0, 0, 0, 0, 0, 0}

func main() {

	input, err := os.Open("11/a/input")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	var count = 0
	for scanner.Scan() {
		if count == 6 {
			count = 0
		} else {
			count++
			var input = strings.ReplaceAll(scanner.Text(), ",", "")
			var str = strings.Fields(input)
			if count == 1 {
				m := Monkey{}
				m.id = str[1]
				monkeys = append(monkeys, m)
			}
			if count == 2 {
				var items = false
				var itemCount = 0
				for _, value := range str {
					if items {
						var item, _ = strconv.Atoi(value)
						it := Item{}
						var monkey = monkeys[len(monkeys)-1]
						it.worry = item
						monkey.items = append(monkey.items, it)
						monkeys[len(monkeys)-1] = monkey
						itemCount++
					}
					if value == "items:" {
						items = true
					}
				}
			}
			if count == 3 {
				monkeys[len(monkeys)-1].operator = str[3] + " " + str[4] + " " + str[5]
			}
			if count == 4 {
				var test, _ = strconv.Atoi(str[3])
				monkeys[len(monkeys)-1].test = test
			}
			if count == 5 {
				var ifTrue, _ = strconv.Atoi(str[5])
				monkeys[len(monkeys)-1].ifTrue = ifTrue
			}
			if count == 6 {
				var ifFalse, _ = strconv.Atoi(str[5])
				monkeys[len(monkeys)-1].ifFalse = ifFalse
			}
		}
	}
	var rounds = 20
	for i := 0; i < rounds; i++ {
		println(i)
		for index, monkey := range monkeys {
			monkeys[index] = inspect(monkey, index)
		}
	}

	fmt.Printf("%#v", monkeyBus)
	println(monkeyBus[0] * monkeyBus[5])
}

func inspect(monkey Monkey, in int) Monkey {
	var str = strings.Fields(monkey.operator)
	var items = monkey.items
	for index, item := range items {
		var v1 = 0
		var v2 = 0
		for point, operation := range str {
			if operation == "old" && point == 0 {
				v1 = item.worry
			}
			if operation == "old" && point == 2 {
				v2 = item.worry
			} else {
				v2, _ = strconv.Atoi(str[2])
			}
		}
		if str[1] == "*" {
			monkey.items[index].worry = (v1 * v2) / 3
		} else {
			monkey.items[index].worry = (v1 + v2) / 3
		}
	}

	for index, item := range items {
		monkeyBus[in]++
		if monkeys[in].items[index].worry%monkey.test == 0 {
			monkey = throw(item, in, monkey.ifTrue, monkey)
		} else {
			monkey = throw(item, in, monkey.ifFalse, monkey)
		}
	}
	return monkey
}

func throw(item Item, index int, dest int, monkey Monkey) Monkey {
	monkeys[dest].items = append(monkeys[dest].items, item)
	for val, it := range monkey.items {
		if it == item {
			if len(monkey.items) == 1 {
				monkey.items = make([]Item, 0)
				return monkey
			}
			ret := make([]Item, 0)
			ret = append(ret, monkey.items[:val]...)
			test := append(ret, monkey.items[val+1:]...)
			monkey.items = test
			return monkey
		}
	}
	return monkey
}

type Monkey struct {
	id       string
	items    []Item
	operator string
	test     int
	ifTrue   int
	ifFalse  int
}

type Item struct {
	worry int
}
