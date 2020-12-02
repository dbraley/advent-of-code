package main

import (
	"fmt"
	"os"

	"github.com/dbraley/advent-of-code/2020/day1"
	"github.com/dbraley/advent-of-code/2020/util"
)

func main() {
	in, err := util.ReadFileOfInts("2020/day1/input.csv")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}
	sum := 2020
	v1, v2, err := day1.FindCommon(in, sum)
	if err != nil {
		fmt.Printf("Wasn't able to find an appropriate combination adding up to %v in %v\n", sum, in)
	}
	fmt.Printf("Part 1: %v\n", v1*v2)

	v1, v2, v3, err := day1.FindCommon3(in, sum)
	if err != nil {
		fmt.Printf("Wasn't able to find an appropriate combination adding up to %v in %v\n", sum, in)
	}
	fmt.Printf("Part 2: %v\n", v1*v2*v3)
}
