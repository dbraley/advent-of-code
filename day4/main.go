package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairs, err := file.ReadFile("day4/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	completeEncasedPairs := 0
	overlappingPairs := 0

	for line, pair := range pairs {
		_ = line
		ranges := strings.Split(pair, ",")
		e1range := strings.Split(ranges[0], "-")
		e2range := strings.Split(ranges[1], "-")
		//fmt.Printf("Range for line %v: %v & %v\n", line, e1range, e2range)
		e1low, err := strconv.Atoi(e1range[0])
		if err != nil {
			fmt.Errorf("Error translating val to int\n")
		}
		e2low, err := strconv.Atoi(e2range[0])
		if err != nil {
			fmt.Errorf("Error translating val to int\n")
		}
		e1High, err := strconv.Atoi(e1range[1])
		if err != nil {
			fmt.Errorf("Error translating val to int\n")
		}
		e2high, err := strconv.Atoi(e2range[1])
		if err != nil {
			fmt.Errorf("Error translating val to int\n")
		}

		inside := false
		if e1low >= e2low && e1High <= e2high {
			inside = true
			completeEncasedPairs += 1
		} else if e1low <= e2low && e1High >= e2high {
			inside = true
			completeEncasedPairs += 1
		}

		overlap := false
		if !(e1High < e2low || e2high < e1low) {
			overlap = true
			overlappingPairs += 1
		}

		fmt.Printf("Ranges %d - %d && %d - %d ; overlap: %v ; contains inside range %v\n", e1low, e1High, e2low, e2high, overlap, inside)
	}

	fmt.Printf("Completely Encased Pairs: %v\n", completeEncasedPairs)
	fmt.Printf("Overlap at all: %v\n", overlappingPairs)
}
