package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fileName := "example.txt"
	fileName := "input.txt"

	input, err := file.ReadFile("day14/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("%d\n", len(input))

	pourPoint := math.Point2D{X: 500}

	f := field{}
	f.initWithPourPoint(pourPoint)

	for i, line := range input {
		_ = i
		var lastCoordinate *math.Point2D
		coordinates := strings.Split(line, " -> ")
		for j, coordinateString := range coordinates {
			_ = j
			coordinate := strings.Split(coordinateString, ",")
			x, _ := strconv.Atoi(coordinate[0])
			y, _ := strconv.Atoi(coordinate[1])
			point := math.Point2D{X: x, Y: y}
			//fmt.Printf("Line %d Coordinate %d is %v\n", i, j, point)
			if lastCoordinate != nil {
				// Draw the line
				//fmt.Printf("Drawing rock line from %v to %v\n", lastCoordinate, point)

				// Horizontal line
				if lastCoordinate.X < point.X {
					for col := lastCoordinate.X + 1; col < point.X; col++ {
						f.DrawRock(col, point.Y)
					}
				}
				if lastCoordinate.X > point.X {
					for col := lastCoordinate.X - 1; col > point.X; col-- {
						f.DrawRock(col, point.Y)
					}
				}

				// Vertical line
				if lastCoordinate.Y < point.Y {
					for row := lastCoordinate.Y + 1; row < point.Y; row++ {
						f.DrawRock(point.X, row)
					}
				}
				if lastCoordinate.Y > point.Y {
					for row := lastCoordinate.Y - 1; row > point.Y; row-- {
						f.DrawRock(point.X, row)
					}
				}

			}
			f.DrawEdge(point)
			lastCoordinate = &point
		}
	}

	// Part 2
	floorY := f.maxY + 2
	floorL := f.minX - f.maxY
	floorR := f.maxX + f.maxY
	for x := floorL; x < floorR; x++ {
		f.DrawRock(x, floorY)
	}

	f.print()

	maxSand := 100000
	drawAt := []int{1, 2, 5, 6, 22, 50, 100, 500, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000}
	for s := 0; s < maxSand; s++ {
		if f.WhatIsAt(pourPoint.X, pourPoint.Y) == 'o' {
			fmt.Printf("\n\nSand has blocked the pour point after %d\n", s)
			f.print()
			return
		}
		if contains(drawAt, s) {
			//fmt.Printf("\n\nDropped %d Sand\n", s)
			f.print()
		}
		col := pourPoint.X
		for row := 0; row <= f.maxY; row++ {
			if row == f.maxY {
				fmt.Printf("\n\nDONE! %d sand fell\n", s)
				f.print()
				return
			}
			didDraw, moveCol := maybeDrawSand(f, col, row)
			if didDraw {
				break
			}
			col = col + moveCol
		}
	}

	fmt.Printf("Dropped all Sand!\n")
}

func maybeDrawSand(f field, col int, row int) (bool, int) {
	down := f.WhatIsAt(col, row+1)
	switch down {
	case ' ':
		//continue
	case '#':
		f.DrawSand(col, row)
		return true, 0
	case 'o':
		left := f.WhatIsAt(col-1, row+1)
		switch left {
		case ' ':
			// continue, but move left
			return false, -1
		default:
			right := f.WhatIsAt(col+1, row+1)
			switch right {
			case ' ':
				//continue, but move right
				return false, 1
			default:
				f.DrawSand(col, row)
				return true, 0
			}
		}
	case '^':
		left := f.WhatIsAt(col-1, row+1)
		switch left {
		case ' ':
			// continue, but move left
			return false, -1
		default:
			right := f.WhatIsAt(col+1, row+1)
			switch right {
			case ' ':
				//continue, but move right
				return false, 1
			default:
				f.DrawSand(col, row)
				return true, 0
			}
		}
	default:
		fmt.Printf("\n\nERROR\n\n")
		f.print()
		panic(fmt.Sprintf("Don't know how to work with %c", down))
	}
	return false, 0
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type field struct {
	grid map[int]map[int]rune
	minX int
	maxX int
	maxY int
}

func (f *field) addRune(x, y int, c rune) {
	if x < f.minX {
		f.minX = x
	}
	if x > f.maxX {
		f.maxX = x
	}
	if y > f.maxY {
		f.maxY = y
	}
	if f.grid[x] == nil {
		f.grid[x] = make(map[int]rune)
	}
	f.grid[x][y] = c
}

func (f *field) initWithPourPoint(pp math.Point2D) {
	f.grid = make(map[int]map[int]rune)
	f.minX = pp.X
	f.maxX = pp.X
	f.maxY = pp.Y
	f.addRune(pp.X, pp.Y, '+')
}

func (f *field) print() {
	//fmt.Printf("Grid from (%d,0) to (%d, %d)\n", f.minX, f.maxX, f.maxY)
	//for row := 0; row <= f.maxY; row++ {
	//	for col := f.minX; col <= f.maxX; col++ {
	//		if r, ok := f.grid[col][row]; ok {
	//			fmt.Printf("%c", r)
	//		} else {
	//			fmt.Printf(".")
	//		}
	//	}
	//	fmt.Println()
	//}
}

// DrawEdge draws the edge of a rock line. It's not explicitly stated, but the example seems to suggest edges cannot support sand and it must fall to either the let or right.
func (f *field) DrawEdge(p math.Point2D) {
	f.addRune(p.X, p.Y, '^')
}

func (f *field) DrawRock(x int, y int) {
	f.addRune(x, y, '#')
}

func (f *field) WhatIsAt(x, y int) rune {
	if r, ok := f.grid[x][y]; ok {
		return r
	}
	return ' '
}

func (f *field) DrawSand(x int, y int) {
	f.addRune(x, y, 'o')
}
