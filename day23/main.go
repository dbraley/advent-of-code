package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
)

func main() {
	//fmt.Println("Example:")
	//ept1, ept2 := day23("example.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day23("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day23(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day23/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	elves := readElves(input)

	//print(elves)

	return moveElves(elves, 10)
}

func score(elves map[math.Point2D]*Elf) int {
	minX, maxX, minY, maxY := findMinsAndMaxes(elves)
	xWidth := maxX - minX + 1
	yWidth := maxY - minY + 1
	return xWidth*yWidth - len(elves)
}

func moveElves(elves map[math.Point2D]*Elf, maxRounds int) (int, int) {
	curElves := elves
	round := 0
	round10Score := 0
	for true {
		//fmt.Printf("Round %d\n", round+1)
		// Propose
		proposedCounts := make(map[math.Point2D]int)
		for _, elf := range curElves {
			proposed := elf.ProposeNext(round, curElves)
			if proposed != nil {
				proposedCounts[*proposed]++
			}
		}
		// Move
		moveCount := 0
		collisionCount := 0
		nextElves := make(map[math.Point2D]*Elf)
		for _, elf := range curElves {
			if elf.Proposed != nil {
				if c, ok := proposedCounts[*elf.Proposed]; ok {
					if c > 1 {
						collisionCount++
						//fmt.Printf("  %v: I colided with another elf moving to %v, staying at %v\n", elf.Name, elf.Proposed, elf.Current)
						elf.Proposed = nil
					} else {
						moveCount++
						elf.Current = *elf.Proposed
						elf.Proposed = nil
					}
				} else {
					panic("I messed something up with the proposals")
				}
			}
			nextElves[elf.Current] = elf
		}

		if moveCount == 0 {
			return round10Score, round + 1
		}

		// Print
		//print(nextElves)
		curElves = nextElves
		round++
		if round == 10 {
			round10Score = score(curElves)
		}
	}
	return round10Score, round
}

func print(elves map[math.Point2D]*Elf) {
	minX, maxX, minY, maxY := findMinsAndMaxes(elves)
	fmt.Println(minX, minY)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if e, ok := elves[math.Point2D{X: x, Y: y}]; ok {
				switch e.Name {
				case "Elf-0001":
					fmt.Printf("1")
				case "Elf-2533":
					fmt.Printf("L")
				default:
					fmt.Printf("#")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func findMinsAndMaxes(elves map[math.Point2D]*Elf) (int, int, int, int) {
	minX, maxX, minY, maxY := 1000, 0, 1000, 0
	for c, _ := range elves {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	return minX, maxX, minY, maxY
}

func readElves(input []string) map[math.Point2D]*Elf {
	elfLocMap := make(map[math.Point2D]*Elf)
	elfNumber := 1
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				e := Elf{
					Name:     fmt.Sprintf("Elf-%04d", elfNumber),
					Current:  math.Point2D{X: x, Y: y},
					Proposed: nil,
				}
				elfLocMap[e.Current] = &e
				elfNumber++
			}
		}
	}
	return elfLocMap
}

type Elf struct {
	Name     string
	Current  math.Point2D
	Proposed *math.Point2D
}

func (e *Elf) ProposeNext(round int, allElves map[math.Point2D]*Elf) *math.Point2D {
	north := e.Current.Translate(0, -1)
	south := e.Current.Translate(0, 1)
	west := e.Current.Translate(-1, 0)
	east := e.Current.Translate(1, 0)
	_, isNW := allElves[e.Current.Translate(-1, -1)]
	_, isN := allElves[north]
	_, isNE := allElves[e.Current.Translate(1, -1)]
	_, isE := allElves[east]
	_, isSE := allElves[e.Current.Translate(1, 1)]
	_, isS := allElves[south]
	_, isSW := allElves[e.Current.Translate(-1, 1)]
	_, isW := allElves[west]

	if !(isNW || isN || isNE || isE || isSE || isS || isSW || isW) {
		//fmt.Printf("  %v: I don't need to move from %v!\n", e.Name, e.Current)
		return nil
	}

	for check := 0; check < 4; check++ {
		switch (check + round) % 4 {
		case 0:
			// check N
			if !isN && !isNW && !isNE {
				e.Proposed = &north
				//fmt.Printf("  %v: I can go North from %v to %v!\n", e.Name, e.Current, e.Proposed)
				return &north
			}
		case 1:
			// check S
			if !isS && !isSW && !isSE {
				e.Proposed = &south
				//fmt.Printf("  %v: I can go South from %v to %v!\n", e.Name, e.Current, e.Proposed)
				return &south
			}
		case 2:
			// check W
			if !isW && !isNW && !isSW {
				e.Proposed = &west
				//fmt.Printf("  %v: I can go West from %v to %v!\n", e.Name, e.Current, e.Proposed)
				return &west
			}
		case 3:
			// check E
			if !isE && !isNE && !isSE {
				e.Proposed = &east
				//fmt.Printf("  %v: I can go East from %v to %v!\n", e.Name, e.Current, e.Proposed)
				return &east
			}
		}
	}
	//fmt.Printf("  %v: I cannot move from %v!\n", e.Name, e.Current)
	return nil
}
