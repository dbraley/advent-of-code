package main

import (
	"fmt"
	"os"

	"github.com/dbraley/advent-of-code/2020/day4"
	"github.com/dbraley/advent-of-code/2020/util"
)

func main() {
	in, err := util.ReadFile("2020/day4/input.csv")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	count, err := day4.CountValidPassports(in)
	if err != nil {
		fmt.Printf("Error parsing file %v \n", err)
	}
	fmt.Printf("Part 1: %v\n", count)
}
