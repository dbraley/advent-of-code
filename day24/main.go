package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
)

func main() {
	//fmt.Println("Trivial:")
	//tpt1, tpt2 := day24("trivial.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", tpt1, tpt2)
	//
	//fmt.Println("Example:")
	//ept1, ept2 := day24("example.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)
	//
	fmt.Println("Input:")
	ipt1, ipt2 := day24("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day24(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day24/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	max := math.Point2D{
		X: len(input[0]),
		Y: len(input),
	}
	goal := max.Translate(-1, 0)

	blizzards := makeBlizzards(input)

	checkedAlready := make(map[int]map[math.Point2D]bool)

	startPath := []math.Point2D{
		{2, 1},
	}

	finalPath := findBestPath(startPath, checkedAlready, max, goal, blizzards)

	backPath := findBestPath(finalPath, checkedAlready, max, math.Point2D{2, 1}, blizzards)
	backAgain := findBestPath(backPath, checkedAlready, max, goal, blizzards)

	return len(finalPath) - 1, len(backAgain) - 1
}

func makeBlizzards(input []string) []blizzard {
	var blizzards []blizzard
	for row, line := range input {
		for offset, char := range line {
			isBliz := isBlizzard(char)
			if isBliz {
				newBliz := blizzard{
					start: math.Point2D{
						X: offset + 1,
						Y: row + 1,
					},
					direction: char,
				}
				blizzards = append(blizzards, newBliz)
			}
		}
	}
	return blizzards
}

func findBestPath(startPath []math.Point2D, checkedAlready map[int]map[math.Point2D]bool, max math.Point2D, goal math.Point2D, blizzards []blizzard) []math.Point2D {
	start := startPath[len(startPath)-1]
	paths := [][]math.Point2D{
		startPath,
	}
	var finalPath []math.Point2D

	for true {
		path := paths[0]
		if checkedAlready[len(path)] == nil {
			checkedAlready[len(path)] = make(map[math.Point2D]bool)
		}
		paths = paths[1:]

		curPos := path[len(path)-1]
		if _, ok := checkedAlready[len(path)][curPos]; ok {
			continue
		}

		nextPos := map[math.Point2D]bool{
			curPos.Translate(1, 0):  true, // right
			curPos.Translate(-1, 0): true, // left
			curPos.Translate(0, 1):  true, // down
			curPos.Translate(0, -1): true, // up
			curPos:                  true, // wait
		}
		// prune for invalids
		for p, _ := range nextPos {
			if p.X <= 1 { // in left wall
				nextPos[p] = false
			} else if p.X >= max.X { // in right wall
				nextPos[p] = false
			} else if p.Y <= 1 { // in top wall, unless start
				nextPos[p] = false
			} else if p.Y == max.Y { // in bottom wall, no worries about end, we already dealt with that
				nextPos[p] = false
			}
			if p == start {
				nextPos[p] = true
			}
			if p == goal {
				return append(path, goal)
			}
		}
		nextRound := len(path) // Should be minus 1 for current round, then plus one for next
		for _, b := range blizzards {
			if math.Abs(b.start.X-curPos.X) <= 1 || math.Abs(b.start.Y-curPos.Y) <= 1 {
				bNext := b.posInRound(nextRound, max)
				if _, ok := nextPos[bNext]; ok {
					nextPos[bNext] = false
				}
			}
		}
		for nextPoint, possible := range nextPos {
			if possible {
				paths = append(paths, makeNewPath(path, nextPoint))
			}
		}
		checkedAlready[len(path)][curPos] = true
	}
	return finalPath
}

type blizzard struct {
	start     math.Point2D
	direction rune
}

func (b blizzard) posInRound(round int, max math.Point2D) math.Point2D {
	switch b.direction {
	case '>':
		x := 2 + nnMod(b.start.X-2+round, max.X-2)
		return math.Point2D{X: x, Y: b.start.Y}
	case '<':
		x := 2 + nnMod(b.start.X-2-round, max.X-2)
		return math.Point2D{X: x, Y: b.start.Y}
	case 'v':
		y := 2 + nnMod(b.start.Y-2+round, max.Y-2)
		return math.Point2D{X: b.start.X, Y: y}
	case '^':
		y := 2 + nnMod(b.start.Y-2-round, max.Y-2)
		return math.Point2D{X: b.start.X, Y: y}
	default:
		panic("Blizzard with invalid direction")
	}
	return b.start
}

func isBlizzard(char rune) bool {
	switch char {
	case '>':
		return true
	case '<':
		return true
	case 'v':
		return true
	case '^':
		return true
	default:
		return false
	}
}

func nnMod(x, mod int) int {
	ret := x % mod
	if ret < 0 {
		return ret + mod
	}
	return ret
}

func makeNewPath(pathInspect []math.Point2D, newPos math.Point2D) []math.Point2D {
	newPath := make([]math.Point2D, len(pathInspect)+1)
	copy(newPath, pathInspect)
	newPath[len(pathInspect)] = newPos
	return newPath
}
