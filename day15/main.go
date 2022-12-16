package main

import (
	"errors"
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//fileName := "example.txt"
	//rowInQuestion := 10
	//fieldMax := 20
	//
	fileName := "input.txt"
	rowInQuestion := 2000000
	fieldMax := 4000000

	beaconFreePoints, score := day15(fileName, rowInQuestion, fieldMax)

	fmt.Printf("There are %d spaces where no beacon can exist on row %d\n", beaconFreePoints, rowInQuestion)
	fmt.Printf("Score: %d\n", score)

}

func day15(fileName string, rowInQuestion int, fieldMax int) (int, int) {
	input, err := file.ReadFile("day15/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("%d\n", len(input))

	r := regexp.MustCompile(`Sensor at x=(?P<sx>-?\d+), y=(?P<sy>-?\d+): closest beacon is at x=(?P<bx>-?\d+), y=(?P<by>-?\d+)`)

	var prs []pointradius
	minX, maxX := 100000000, -10000000
	beaconsOnRowInQuestion := make(map[math.Point2D]int)

	for _, line := range input {
		fmt.Printf("%v\n", line)
		substrings := r.FindStringSubmatch(line)
		sx, _ := strconv.Atoi(substrings[1])
		sy, _ := strconv.Atoi(substrings[2])
		bx, _ := strconv.Atoi(substrings[3])
		by, _ := strconv.Atoi(substrings[4])

		minX = min(minX, sx, bx)
		maxX = max(maxX, sx, bx)

		if by == rowInQuestion {
			beaconsOnRowInQuestion[math.Point2D{X: bx, Y: by}]++
		}

		mdist := math.Abs(sx-bx) + math.Abs(sy-by)
		prs = append(prs, pointradius{sx, sy, mdist})
	}

	beaconFreePoints := part1BeaconFreePoints(minX, maxX, rowInQuestion, prs, beaconsOnRowInQuestion)

	score, _ := part2Score(prs, fieldMax)
	return beaconFreePoints, score
}

func part1BeaconFreePoints(minX int, maxX int, rowInQuestion int, prs []pointradius, beaconsOnRowInQuestion map[math.Point2D]int) int {
	sum := 0
	for c := minX; c <= maxX; c++ {
		if !check(c, rowInQuestion, maxX, prs) {
			sum++
		}
	}

	beaconlessPoints := sum - len(beaconsOnRowInQuestion)
	return beaconlessPoints
}

func part2Score(prs []pointradius, fieldMax int) (int, error) {
	score := 0
	for i, pr := range prs {
		fmt.Printf("checking pr %d - (%d, %d) %d\n", i, pr.x, pr.y, pr.r)

		// topRight, travels [12-3)
		tr := math.Point2D{X: pr.x, Y: pr.y - pr.r - 1}
		// botRight, travels [3-6)
		br := math.Point2D{X: pr.x + pr.r + 1, Y: pr.y}
		// botLeft, travels [6-9)
		bl := math.Point2D{X: pr.x, Y: pr.y + pr.r + 1}
		// topLeft, travels [9-12)
		tl := math.Point2D{X: pr.x - pr.r - 1, Y: pr.y}

		pointsOutsideSignal := make(map[math.Point2D]int)
		for j := 0; j < pr.r+1; j++ {
			pointsOutsideSignal[tr.Translate(j, j)]++
			pointsOutsideSignal[br.Translate(-j, j)]++
			pointsOutsideSignal[bl.Translate(-j, -j)]++
			pointsOutsideSignal[tl.Translate(j, -j)]++
		}
		for p := range pointsOutsideSignal {
			if check(p.X, p.Y, fieldMax, prs) {
				score = p.X*4000000 + p.Y
				return 0, nil
			}
		}
	}
	return score, errors.New("no empty point found")
}

// TODO: Move to math
func min(a int, b ...int) int {
	m := a
	for _, other := range b {
		if other < m {
			m = other
		}
	}
	return m
}
func max(a int, b ...int) int {
	m := a
	for _, other := range b {
		if other > m {
			m = other
		}
	}
	return m
}

func check(x, y, max int, prs []pointradius) bool {
	if x < 0 || y < 0 || x > max || y > max {
		return false
	}
	for i, pr := range prs {
		_ = i
		distToSensor := math.Abs(pr.x-x) + math.Abs(pr.y-y)
		if distToSensor <= pr.r {
			return false
		}
	}
	fmt.Printf("Empty space found at (%d, %d)\n", x, y)
	fmt.Printf("Score: %d\n", x*4000000+y)
	return true
}

type pointradius struct {
	x int
	y int
	r int
}
