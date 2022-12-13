package main

import (
	"errors"
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"strings"
)

type ValuePoint struct {
	Point math.Point2D
	Value rune
}

func (p ValuePoint) Translate(x, y int) ValuePoint {
	return ValuePoint{Point: p.Point.Translate(x, y)}
}

func (p ValuePoint) canTravel(maybe ValuePoint) bool {
	if maybe.Value == 'S' {
		return p.Value == 'a' || p.Value == 'b'
	}
	if p.Value == 'E' {
		return maybe.Value == 'z' || maybe.Value == 'y'
	}
	return int(p.Value-maybe.Value) <= 1
}

func main() {
	//f := "example.txt"
	f := "input.txt"

	input, err := file.ReadFile("day12/" + f)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", len(input))

	startPos, err := findFirst(input, 'E')
	if err != nil {
		fmt.Printf("Error finding start %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Start Position %v\n", startPos)

	paths := [][]ValuePoint{{startPos}}
	var solution []ValuePoint
	seen := []ValuePoint{startPos}

	i := 0
	maxI := 10000
	for {
		i++ // Safety
		pathInspect := paths[0]
		paths = paths[1:]

		fmt.Printf("Inspecting %v\n", pathInspect)
		lastPos := pathInspect[len(pathInspect)-1]

		rightPos := lastPos.Translate(1, 0)
		if isValidPoint(rightPos.Point, input) {
			rightPos.Value = []rune(input[rightPos.Point.Y])[rightPos.Point.X]
			if lastPos.canTravel(rightPos) {
				fmt.Printf("\tCan travel Right %v!\n", rightPos)
				if !seenAlready(seen, rightPos) {
					newPath := makeNewPath(pathInspect, rightPos)
					if isEnd(rightPos) {
						solution = newPath
						break
					}
					seen = append(seen, rightPos)
					paths = append(paths, newPath)
				}
			}
		}

		downPos := lastPos.Translate(0, 1)
		if isValidPoint(downPos.Point, input) {
			downPos.Value = []rune(input[downPos.Point.Y])[downPos.Point.X]
			if lastPos.canTravel(downPos) {
				fmt.Printf("\tCan travel Down! %v\n", downPos)
				if !seenAlready(seen, downPos) {
					newPath := makeNewPath(pathInspect, downPos)
					if isEnd(downPos) {
						solution = newPath
						break
					}
					seen = append(seen, downPos)
					paths = append(paths, newPath)
				}
			}
		}

		leftPos := lastPos.Translate(-1, 0)
		if isValidPoint(leftPos.Point, input) {
			leftPos.Value = []rune(input[leftPos.Point.Y])[leftPos.Point.X]
			if lastPos.canTravel(leftPos) {
				fmt.Printf("\tCan travel Left! %v\n", leftPos)
				if !seenAlready(seen, leftPos) {
					newPath := makeNewPath(pathInspect, leftPos)
					if isEnd(leftPos) {
						solution = newPath
						break
					}
					seen = append(seen, leftPos)
					paths = append(paths, newPath)
				}
			}
		}

		upPos := lastPos.Translate(0, -1)
		if isValidPoint(upPos.Point, input) {
			upPos.Value = []rune(input[upPos.Point.Y])[upPos.Point.X]
			if lastPos.canTravel(upPos) {
				fmt.Printf("\tCan travel Up!%v\n", upPos)
				if !seenAlready(seen, upPos) {
					newPath := makeNewPath(pathInspect, upPos)
					if isEnd(upPos) {
						solution = newPath
						break
					}
					seen = append(seen, upPos)
					paths = append(paths, newPath)
				}
			}
		}

		fmt.Printf("number of paths left: %v\n", len(paths))

		if len(paths) == 0 {
			fmt.Printf("No More Paths!\n")
			break
		}

		if i > maxI { // Safety
			fmt.Printf("Breaking for max iterations\n")
			break
		}
	}

	fmt.Printf("Solution: %v\n", solution)
	fmt.Printf("Length: %v\n", len(solution)-1)

}

func seenAlready(seen []ValuePoint, pos ValuePoint) bool {
	for _, s := range seen {
		if pos.Point.X == s.Point.X && pos.Point.Y == s.Point.Y {
			fmt.Printf("\t\tCan already get to %v\n", pos)
			return true
		}
	}
	return false
}

func isEnd(pos ValuePoint) bool {
	return pos.Value == 'S' || pos.Value == 'a'
}

func makeNewPath(pathInspect []ValuePoint, newPos ValuePoint) []ValuePoint {
	newPath := make([]ValuePoint, len(pathInspect)+1)
	copy(newPath, pathInspect)
	newPath[len(pathInspect)] = newPos
	return newPath
}

func isValidPoint(maybe math.Point2D, input []string) bool {
	return maybe.Y >= 0 &&
		maybe.X >= 0 &&
		maybe.Y < len(input) &&
		maybe.X < len(input[maybe.Y])
}

func findFirst(in []string, searchFor rune) (ValuePoint, error) {
	for y, line := range in {
		x := strings.IndexFunc(line, func(r rune) bool {
			return r == searchFor
		})
		if x != -1 {
			return ValuePoint{Point: math.Point2D{X: x, Y: y}, Value: searchFor}, nil
		}
	}
	errorMsg := fmt.Sprintf("rune %v not found", searchFor)
	return ValuePoint{}, errors.New(errorMsg)
}
