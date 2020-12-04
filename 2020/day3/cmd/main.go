package main

import (
	"fmt"
	"os"

	"github.com/dbraley/advent-of-code/2020/day3"
	"github.com/dbraley/advent-of-code/2020/util"
)

func main() {
	in, err := util.ReadFile("2020/day3/input.csv")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	treeCount := day3.CountTreesOnPath(in)
	fmt.Printf("Part 1: %v\n", treeCount)
}
