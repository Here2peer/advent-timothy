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
	//directories := orderedmap.NewOrderedMap()
	//dirSizeTest := orderedmap.NewOrderedMap()
	//var currentDir = ""
	//var count = 0
	//var prev = 0
	//prevDir := orderedmap.NewOrderedMap()
	//var ls = false
	//var totalTest = 0
	//var prevDir []string
	//var prev string
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
	println(small)

	//for scanner.Scan() {
	//
	//	if string(scanner.Text()[0]) == "$" {
	//		if string(scanner.Text()[2]) == "c" {
	//			ls = false
	//			if string(scanner.Text()[5]) == "." {
	//				prevDir.Delete(prev)
	//				prev--
	//				currentDir = prevDir.Back().Value.(string)
	//			} else {
	//				prevDir.Set(prev, currentDir)
	//				currentDir = strings.Split(scanner.Text(), " ")[2]
	//				prev++
	//				//println(prev, currentDir, prevDir.GetElement(prev-1).Value.(string))
	//				//println(string(currentDir), "<--- ")
	//			}
	//		} else if string(scanner.Text()[2]) == "l" {
	//			ls = true
	//			if currentDir != "" {
	//				prevDir.Set(prev, currentDir)
	//			}
	//			count = 0
	//			//println(string(scanner.Text()[2]), "<--- ")
	//			scanner.Scan()
	//		}
	//	}
	//
	//	if ls {
	//		if string(scanner.Text()[0]) == "d" {
	//			//directories[currentDir+","+strconv.Itoa(count)] = strings.Split(scanner.Text(), " ")[1]
	//			directoriesTest.Set(any(prevDir.GetElement(prev-1).Value.(string)+"-"+currentDir+","+strconv.Itoa(count)), any(strings.Split(scanner.Text(), " ")[1]))
	//		} else {
	//
	//			//println(strconv.Itoa(count))
	//			//println((prevDir.GetElement(prev-1).Value.(string) + "-" + currentDir + "," + strconv.Itoa(count)), (scanner.Text()))
	//			var k, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	//			totalTest = totalTest + k
	//			directoriesTest.Set(any(prevDir.GetElement(prev-1).Value.(string)+"-"+currentDir+","+strconv.Itoa(count)), any(scanner.Text()))
	//			//directories[currentDir+","+strconv.Itoa(count)] = scanner.Text()
	//		}
	//		count++
	//	}
	//}
	//var prevDir1 = ""
	//for el := directoriesTest.Back(); el != nil; el = el.Prev() {
	//	var dir = fmt.Sprintf("%v", el.Key)
	//	var value = fmt.Sprintf("%v", el.Value)
	//	var parsedDir = strings.Split(dir, ",")
	//	var fileSize = strings.Split(value, " ")
	//	//println(parsedDir[0], value, fileSize)
	//	//println(dir, value)
	//
	//	if parsedDir[0] != "" {
	//		//println(value)
	//		if len(fileSize) != 1 {
	//			var size, _ = strconv.Atoi(fileSize[0])
	//			//println(fileSize[0])
	//			var initialSize, initialized = dirSizeTest.Get(parsedDir[0])
	//			if !initialized {
	//				dirSizeTest.Set(parsedDir[0], size)
	//			} else {
	//				dirSizeTest.Set(parsedDir[0], initialSize.(int)+size)
	//			}
	//		} else {
	//			var test, _ = dirSizeTest.Get(prevDir1)
	//			var test1, _ = dirSizeTest.Get(parsedDir[0])
	//			if test1 == nil {
	//				dirSizeTest.Set(parsedDir[0], test.(int))
	//			} else {
	//				dirSizeTest.Set(parsedDir[0], test.(int)+test1.(int))
	//			}
	//		}
	//	}
	//	if parsedDir[1] == "0" {
	//		prevDir1 = parsedDir[0]
	//	}
	//	//println(parsedDir[0], value)
	//}
	//println(" ")
	//var sumTest = 0
	//for el := directoriesTest.Back(); el != nil; el = el.Prev() {
	//	var dir = fmt.Sprintf("%v", el.Key)
	//	var value = fmt.Sprintf("%v", el.Value)
	//	var parsedDir = strings.Split(dir, ",")[0]
	//	if string(parsedDir[0]) == "/" || string(parsedDir[0]) == "-" {
	//		parsedDir = "/"
	//	}
	//	var parsedValue = strings.Split(value, " ")
	//	if len(parsedValue) == 1 {
	//		//println(string(dir), parsedValue[0])
	//
	//		var initialSize, _ = dirSizeTest.Get(parsedDir)
	//		var sumSize, _ = dirSizeTest.Get(parsedValue[0])
	//		if initialSize == nil && sumSize != nil {
	//			dirSizeTest.Set(parsedDir, sumSize.(int))
	//		} else if sumSize == nil && initialSize != nil {
	//			dirSizeTest.Set(parsedDir, initialSize.(int))
	//		} else if sumSize != nil && initialSize != nil {
	//			dirSizeTest.Set(parsedDir, initialSize.(int)+sumSize.(int))
	//		} else {
	//			//println(sumSize.(int))
	//			parsedDir = strings.Split(parsedDir, "-")[0]
	//			//println(string(parsedDir[0]))
	//		}
	//	} else {
	//		var test, _ = dirSizeTest.Get(parsedDir)
	//		var test2, _ = strconv.Atoi(parsedValue[0])
	//		if test == nil {
	//			dirSizeTest.Set(parsedDir, test2)
	//		} else {
	//			var parse = test.(int) + test2
	//			//println(test2 + test.(int))
	//			dirSizeTest.Set(parsedDir, parse)
	//		}
	//	}
	//}
	//var total = 0
	//println("Test: ", sumTest)
	//
	//var parsePrev string
	//for el := dirSizeTest.Front(); el != nil; el = el.Next() {
	//	var value = el.Value.(int)
	//	var curr = strings.Split(el.Key.(string), "-")[0]
	//	if parsePrev == curr {
	//		if value <= 100000 {
	//			//println(dir, value)
	//			total = total + value
	//		}
	//	}
	//	println(el.Key.(string), el.Value.(int))
	//	if value <= 100000 {
	//		//println(dir, value)
	//		total = total + value
	//	}
	//	parsePrev = strings.Split(el.Key.(string), "-")[0]
	//}
	//println("Total value: ", total)
	//
	//println(totalTest)
}

func appendArray(array []string, value string) []string {
	return append(array, value)
}

func popArray(array []string) []string {
	var _, a = array[len(array)-1], array[:len(array)-1]
	return a
}
