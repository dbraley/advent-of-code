package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"strconv"
)

type point struct {
	math.Point2D
}

func (p point) follow(head point) (point, bool) {
	needToMove := !math.WithinOne(p.X, head.X) || !math.WithinOne(p.Y, head.Y)
	if !needToMove {
		return p, false
	}
	xVec, yVec := 0, 0
	if head.X > p.X {
		xVec = 1
	} else if head.X < p.X {
		xVec = -1
	}
	if head.Y > p.Y {
		yVec = 1
	} else if head.Y < p.Y {
		yVec = -1
	}
	return point{p.Translate(xVec, yVec)}, true
}

func main() {
	//f := "example.txt"
	f := "input.txt"

	//length := 2
	length := 10

	input, err := file.ReadSSV("day9/" + f)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("%d\n", len(input))

	tailVisits := make(map[point]int)

	knots := make(map[int]point)

	for i := 0; i < length; i++ {
		knots[i] = point{}
	}
	tailVisits[knots[length-1]] = 1

	for r, line := range input {
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[1])
		fmt.Printf("%v: %v %d\n", r, direction, magnitude)

		xVec, yVec := 0, 0
		switch direction {
		case "R":
			xVec = 1
		case "L":
			xVec = -1
		case "U":
			yVec = 1
		case "D":
			yVec = -1
		default:
			fmt.Printf("Unable to move in direction %v", direction)
			os.Exit(2)
		}

		for m := 0; m < magnitude; m++ {
			knots[0] = point{knots[0].Translate(xVec, yVec)}
			for k := 1; k < length; k++ {
				moved := false
				knots[k], moved = knots[k].follow(knots[k-1])
				if moved && k == length-1 {
					tailVisits[knots[length-1]]++
				}
			}

			//fmt.Printf("Knots: %v\tTailVisits: %d\n", knots, len(tailVisits))
		}

	}

	fmt.Printf("\n\nTail Visited Points: %d\n", len(tailVisits))
}
