package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"time"
)

func main() {
	//fmt.Println("Example:")
	//ept1, ept2 := day17("example.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day17("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day17(fileName string) (pt1 string, pt2 string) {
	input, err := file.ReadFile("day17/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	wind := input[0]
	windcount := len(wind)

	i := 0

	rocks := map[int]rock{
		0: {
			points: []math.Point2D{
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
				{X: 4, Y: 1}},
		},
		1: {
			points: []math.Point2D{
				{X: 2, Y: 1},
				{X: 1, Y: 2},
				{X: 2, Y: 2},
				{X: 3, Y: 2},
				{X: 2, Y: 3}},
		},
		2: {
			points: []math.Point2D{
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
				{X: 3, Y: 2},
				{X: 3, Y: 3}},
		},
		3: {
			points: []math.Point2D{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 1, Y: 3},
				{X: 1, Y: 4}},
		},
		4: {
			points: []math.Point2D{
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 1, Y: 2},
				{X: 2, Y: 2}},
		},
	}

	rockAt := map[int]map[int]rune{
		1: {0: '#'},
		2: {0: '#'},
		3: {0: '#'},
		4: {0: '#'},
		5: {0: '#'},
		6: {0: '#'},
		7: {0: '#'},
	}

	highestPoint := 0

	numberOfRocks := 1000000000000
	pattern := []int{}
	start := time.Now()
	// rn == 1220
	preambleHeight := 0
	// rn == 1220 + 1725n
	patternHeight := 0
	delta := 0
	deltaCount := (numberOfRocks-1220)%1725 + 1220
	patternReps := (numberOfRocks - 1220) / 1725
	for rn := 0; rn < numberOfRocks; rn++ {
		if rn == 1220 {
			preambleHeight = highestPoint
		}
		if rn == 1220+1725 {
			patternHeight = highestPoint - preambleHeight
			fmt.Printf("\nEstimating: %d\n", preambleHeight+delta+patternHeight*patternReps)
			//break
		}
		if rn == deltaCount {
			delta = highestPoint - preambleHeight
		}
		if rn%(5) == 0 {
			pattern = append(pattern, highestPoint)
			//fmt.Printf("%12d (%v):\n", rn, time.Now())
			//for _, c := range rockAt {
			//	for ri, _ := range c {
			//		if ri > (highestPoint - 20) {
			//			//fmt.Printf("%d,", highestPoint-ri)
			//		} else {
			//			delete(c, ri)
			//		}
			//	}
			//fmt.Println()
			//}

		}
		rockPos := math.Point2D{X: 2, Y: highestPoint + 3}
		rock, _ := rocks[rn%5]
		pts := []math.Point2D{}
		for _, p := range rock.points {
			pts = append(pts, p.Translate(rockPos.X, rockPos.Y))
		}
		for j := 0; j < rockPos.Y+1; j++ {
			// wind blows
			switch wind[i%windcount] {
			case '>':
				newPts, ok := canMove(pts, 1, 0, rockAt)
				if ok {
					pts = newPts
				}
			case '<':
				newPts, ok := canMove(pts, -1, 0, rockAt)
				if ok {
					pts = newPts
				}
			}
			i++

			// rock drops
			newPts, ok := canMove(pts, 0, -1, rockAt)
			if ok {
				pts = newPts
			} else {
				// rock stopped!
				for _, p := range pts {
					highestPoint = max(highestPoint, p.Y)
					rockAt[p.X][p.Y] = '#'
				}
				break
			}
		}
		if rn == 2021 {
			fmt.Printf("%12d (%v):\n", rn, time.Now().Sub(start))
			pt1 = fmt.Sprintf("%d", highestPoint)
		}
	}

	fmt.Printf("%12d (%v):\n", numberOfRocks, time.Now().Sub(start))
	for p := 1; p < len(pattern); p++ {
		fmt.Printf("%d ", pattern[p]-pattern[p-1])
	}
	fmt.Println()
	pt2 = fmt.Sprintf("%d", highestPoint)

	return pt1, pt2
}

func canMove(pts []math.Point2D, x int, y int, rockAt map[int]map[int]rune) ([]math.Point2D, bool) {
	var newPts []math.Point2D
	canMove := true
	for _, p := range pts {
		newPt := p.Translate(x, y)
		if newPt.X > 7 || newPt.X < 1 {
			return []math.Point2D{}, false
		}
		if _, ok := rockAt[newPt.X][newPt.Y]; ok {
			return []math.Point2D{}, false
		}
		newPts = append(newPts, newPt)
	}
	return newPts, canMove
}

func maxValue(m map[int]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

type rock struct {
	name   string
	points []math.Point2D
	width  int
	height int
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
