package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := file.ReadFile("day7/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	dirSize := make(map[string]int)

	var curDir []string
	for lineNumber, line := range input {
		//fmt.Printf("%d: %v\n", lineNumber, line)
		lineParts := strings.Split(line, " ")
		if lineParts[0] == "$" {
			callArgs := lineParts[1:]
			switch callArgs[0] {
			case "cd":
				switch callArgs[1] {
				case "/":
					curDir = []string{}
				case "..":
					curDir = curDir[0 : len(curDir)-1]
				default:
					curDir = append(curDir, callArgs[1])
				}
				fmt.Printf("%d: cd %v -> (%q)\n", lineNumber, callArgs[1], "/"+strings.Join(curDir, "/"))
			case "ls":
				fmt.Printf("%d: ls\n", lineNumber)
			}
		} else {
			// This assumes the previous command was 'ls'
			switch lineParts[0] {
			case "dir":
				// Do nothing
			default:
				size, err := strconv.Atoi(lineParts[0])
				if err != nil {
					fmt.Printf("Error converting %v to int at line %d: %v\n", lineParts[0], lineNumber, err)
					os.Exit(2)
				}
				for i := len(curDir); i >= 0; i-- {
					path := "/" + strings.Join(curDir[:i], "/")
					dirSize[path] = dirSize[path] + size
					fmt.Printf("Adding %d to %v (now: %d)\n", size, path, dirSize[path])
				}
			}
			fmt.Printf("%d: output=%q\n", lineNumber, line)
		}
	}
	fmt.Println("")
	sum := 0
	neededFree := dirSize["/"] - 40000000
	bestDir := "/"
	bestDirSize := dirSize["/"]
	fmt.Printf("Need to Free Up: %d\n", neededFree)
	for path, size := range dirSize {
		//fmt.Printf("%v: %d\n", path, size)
		if size < 100000 {
			sum += size
		}
		if size > neededFree {
			fmt.Printf("%v is a candidate at size %d... ", path, size)
			if size < bestDirSize {
				fmt.Printf("Replacing %v with %v!\n", bestDir, path)
				bestDir = path
				bestDirSize = size
			} else {
				fmt.Printf("Larger than %v:%d\n", bestDir, bestDirSize)
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)

	fmt.Printf("Best to delete %v:%d\n", bestDir, bestDirSize)
}
