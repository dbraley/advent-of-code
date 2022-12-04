package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strings"
)

func main() {
	rucksacks, err := file.ReadFile("day3/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	sumOfCompartmentalErrors := 0
	sumOfBadges := 0

	for i, rucksack := range rucksacks {
		//fmt.Printf("Rucksack %v has size %d\n", rucksack, len(rucksack))
		runes := []rune(rucksack)
		rucksackSize := len(rucksack)
		comp1 := string(runes[0 : rucksackSize/2])
		comp2 := string(runes[rucksackSize/2 : rucksackSize])

		//fmt.Printf("Rucksack %v -> %v && %v\n", rucksack, comp1, comp2)

		for _, r := range comp1 {
			if strings.ContainsRune(comp2, r) {
				score := scoreForRune(r)
				sumOfCompartmentalErrors += score
				//fmt.Printf("Found duplicate %v %q %v -> %v\n", r, string(r), score, sumOfCompartmentalErrors)
				break
			}
		}

		if i%3 == 0 {
			for _, r := range rucksack {
				if strings.ContainsRune(rucksacks[i+1], r) && strings.ContainsRune(rucksacks[i+2], r) {
					badgeScore := scoreForRune(r)
					sumOfBadges += badgeScore
					fmt.Printf("Badge for group %v - %v is %v, score %v, total %v\n", i, i+2, string(r), badgeScore, sumOfBadges)
					break
				}
				fmt.Errorf("didn't find a badge for group %v-%v\n", i, i+2)
			}
		}
	}

	fmt.Printf("sumOfCompartmentalErrors: %v\n", sumOfCompartmentalErrors)
	fmt.Printf("sumOfBadges: %v\n", sumOfBadges)
}

func scoreForRune(r int32) int {
	score := int(r) - 96
	if score < 1 {
		score = score + 58
	}
	return score
}
