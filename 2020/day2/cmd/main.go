package main

import (
	"fmt"
	"os"

	"github.com/dbraley/advent-of-code/2020/day2"
	"github.com/dbraley/advent-of-code/2020/util"
)

func main() {
	in, err := util.ReadSSV("2020/day2/input.csv")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}
	count, err := day2.CountValid(in)
	if err != nil {
		fmt.Printf("Error analyzing file %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", count)
}
