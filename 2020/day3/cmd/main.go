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

	treeCount3_1 := day3.CountTreesOnPath(in, 3, 1)
	fmt.Printf("Part 1: %v\n", treeCount3_1)

	treeCount1_1 := day3.CountTreesOnPath(in, 1, 1)
	fmt.Printf("1-1 trees: %v\n", treeCount1_1)

	treeCount5_1 := day3.CountTreesOnPath(in, 5, 1)
	fmt.Printf("5-1 trees: %v\n", treeCount5_1)

	treeCount7_1 := day3.CountTreesOnPath(in, 7, 1)
	fmt.Printf("7-1 trees: %v\n", treeCount7_1)

	treeCount1_2 := day3.CountTreesOnPath(in, 1, 2)
	fmt.Printf("1-2 trees: %v\n", treeCount1_2)

	fmt.Printf("Part 2: %v\n", treeCount1_1*treeCount3_1*treeCount5_1*treeCount7_1*treeCount1_2)
}
