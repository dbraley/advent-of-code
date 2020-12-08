package main

import (
	"fmt"
	"os"

	"github.com/dbraley/advent-of-code/2020/day5"
	"github.com/dbraley/advent-of-code/2020/util"
)

func main() {
	in, err := util.ReadFile("2020/day5/input.csv")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	max, seat := day5.Find(in)
	fmt.Printf("Part 1: %v\n", max)
	fmt.Printf("Part 2: %v\n", seat)
}
