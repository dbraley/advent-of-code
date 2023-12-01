package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("Example:")
	//ept1, ept2 := day22("example.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day22("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day22(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day22/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	//r, c, d := runMaze(input, 1)
	r, c, d := 0, 0, '>'

	pt1 = score(r, c, d)

	r, c, d = runMaze(input, 2)

	pt2 = score(r, c, d)

	return pt1, pt2
}

func runMaze(input []string, pt int) (int, int, rune) {
	directions := input[len(input)-1]
	board := input[:len(input)-2]

	posRow, posCol := findStart(board)
	facing := '>'

	fmt.Println(posRow, posCol, facing)

	for true {
		newDirIndex := strings.IndexFunc(directions, func(r rune) bool {
			return r == 'R' || r == 'L'
		})
		if newDirIndex == -1 {
			dist, _ := strconv.Atoi(directions)
			posRow, posCol, facing = move(board, pt, facing, posRow, posCol, dist)
			fmt.Printf("Went %c %d, now done!\n", facing, dist)
			break
		}
		dist, _ := strconv.Atoi(directions[:newDirIndex])
		posRow, posCol, facing = move(board, pt, facing, posRow, posCol, dist)
		turn := rune(directions[newDirIndex])
		fmt.Printf("went %c %d, now turn %c\n", facing, dist, turn)
		facing = changeDir(turn, facing)
		directions = directions[newDirIndex+1:]
	}

	return posRow + 1, posCol + 1, facing
}

func move(board []string, pt int, facing rune, startRow int, startCol int, dist int) (int, int, rune) {
	posRow := startRow
	posCol := startCol
	newFacing := facing
	couldMove := true
	for i := 0; i < dist; i++ {
		fmt.Printf("  Moving %c from %d, %d\n", newFacing, posRow, posCol)
		posRow, posCol, newFacing, couldMove = moveOne(board, pt, newFacing, posRow, posCol)
		if !couldMove {
			return posRow, posCol, newFacing
		}
	}
	return posRow, posCol, newFacing
}

func moveOne(board []string, pt int, facing rune, row int, col int) (int, int, rune, bool) {
	newRow := row
	newCol := col
	newFacing := facing
	switch facing {
	case '>':
		newCol = col + 1
		if !isValid(board, row, newCol) {
			if pt == 1 {
				for i := 1; i < len(board[row]); i++ {
					maybeCol := (col + i) % len(board[row])
					if isValid(board, row, maybeCol) {
						newCol = maybeCol
						break
					}
				}
			} else {
				newRow, newCol, newFacing = wrapAroundCube(board, col, row, facing)
			}
		}
	case 'v':
		newRow = row + 1
		if !isValid(board, newRow, col) {
			if pt == 1 {
				for i := 1; i < len(board); i++ {
					maybeRow := (row + i) % len(board)
					if isValid(board, maybeRow, col) {
						newRow = maybeRow
						break
					}
				}
			} else {
				newRow, newCol, newFacing = wrapAroundCube(board, col, row, facing)
			}
		}
	case '<':
		newCol = col - 1
		if !isValid(board, row, newCol) {
			if pt == 1 {
				for i := 1; i < len(board[row]); i++ {
					maybeCol := (len(board[row]) + col - i) % len(board[row])
					if isValid(board, row, maybeCol) {
						newCol = maybeCol
						break
					}
				}
			} else {
				newRow, newCol, newFacing = wrapAroundCube(board, col, row, facing)
			}
		}
	case '^':
		newRow = row - 1
		if !isValid(board, newRow, newCol) {
			if pt == 1 {
				for i := 1; i < len(board); i++ {
					maybeRow := (len(board) + row - i) % len(board)
					if isValid(board, maybeRow, col) {
						newRow = maybeRow
						break
					}
				}
			} else {
				newRow, newCol, newFacing = wrapAroundCube(board, col, row, facing)
			}
		}
	}
	if board[newRow][newCol] == '#' {
		fmt.Printf("Hit a wall at %d, %d\n", newRow, col)
		return row, col, facing, false
	}
	return newRow, newCol, newFacing, true
}

func wrapAroundCube(board []string, col int, row int, facing rune) (int, int, rune) {
	var newRow, newCol int
	var newFacing rune
	// do part 2
	cur := PointAndDir{
		X:   col,
		Y:   row,
		Dir: facing,
	}
	var folds map[PointAndDir]PointAndDir
	if len(board) < 20 {
		// example
		folds = cube4Edges
	} else {
		// input
		folds = cube50Edges
	}
	if newPoint, ok := folds[cur]; ok {
		newRow = newPoint.Y
		newCol = newPoint.X
		newFacing = newPoint.Dir
	} else {
		if row == 0 && col >= 50 && col < 100 && facing == '^' {
			// from top of 1 to left of 6
			newRow = col + 100
			newCol = 0
			newFacing = '>'
			return newRow, newCol, newFacing
		}
		if row == 0 && col >= 100 && col < 150 && facing == '^' {
			// from top of 2 to bottom of 6
			newRow = 199
			newCol = col - 100
			newFacing = '^'
			return newRow, newCol, newFacing
		}
		if row == 100 && col >= 0 && col < 50 && facing == '^' {
			// from top of 5 to left of 3
			newRow = 50 + col
			newCol = 50
			newFacing = '>'
			return newRow, newCol, newFacing
		}
		if row == 49 && col >= 100 && col < 150 && facing == 'v' {
			// from bottom of 2 to right of 3
			newRow = col - 50
			newCol = 99
			newFacing = '<'
			return newRow, newCol, newFacing
		}
		if row == 149 && col >= 50 && col < 100 && facing == 'v' {
			// from bottom of 4 to right of 6
			newRow = col + 100
			newCol = 49
			newFacing = '<'
			return newRow, newCol, newFacing
		}
		if row == 199 && col >= 0 && col < 50 && facing == 'v' {
			// from bottom of 6 to top of 2
			newRow = 0
			newCol = col + 100
			newFacing = 'v'
			return newRow, newCol, newFacing
		}
		if col == 0 && row >= 100 && row < 150 && facing == '<' {
			// from left of 5 to left of 1
			newRow = 149 - row
			newCol = 50
			newFacing = '>'
			return newRow, newCol, newFacing
		}
		if col == 0 && row >= 150 && row < 200 && facing == '<' {
			// from left of 6 to top of 1
			newRow = 0
			newCol = row - 100
			newFacing = 'v'
			return newRow, newCol, newFacing
		}
		if col == 50 && row >= 0 && row < 50 && facing == '<' {
			// from left of 1 to left of 5
			newRow = 149 - row
			newCol = 0
			newFacing = '>'
			return newRow, newCol, newFacing
		}
		if col == 50 && row >= 50 && row < 100 && facing == '<' {
			// from left of 3 to top of 5
			newRow = 100
			newCol = row - 50
			newFacing = 'v'
			return newRow, newCol, newFacing
		}
		if col == 49 && row >= 150 && row < 200 && facing == '>' {
			// right of 6 to bottom of 4
			newRow = 149
			newCol = row - 100
			newFacing = '^'
			return newRow, newCol, newFacing
		}
		if col == 99 && row >= 50 && row < 100 && facing == '>' {
			// right of 3 to bottom of 2
			newRow = 49
			newCol = row + 50
			newFacing = '^'
			return newRow, newCol, newFacing
		}
		if col == 99 && row >= 100 && row < 150 && facing == '>' {
			// right of 4 to right of 2
			newRow = 149 - row
			newCol = 149
			newFacing = '<'
			return newRow, newCol, newFacing
		}
		if col == 149 && row >= 0 && row < 50 && facing == '>' {
			// right of 2 to right of 4
			newRow = 149 - row
			newCol = 99
			newFacing = '<'
			return newRow, newCol, newFacing
		}
		panic("Fuck")
	}
	return newRow, newCol, newFacing
}

func isValid(board []string, maybeRow, maybeCol int) bool {
	if maybeRow < 0 || len(board) <= maybeRow || (maybeCol < 0 || len(board[maybeRow]) <= maybeCol) || board[maybeRow][maybeCol] == ' ' {
		return false
	}
	return true
}

func findStart(board []string) (int, int) {
	startCol := strings.IndexFunc(board[0], func(r rune) bool {
		return r == '.'
	})
	return 0, startCol
}

func changeDir(turn, facing rune) rune {
	switch turn {
	case 'L':
		switch facing {
		case '>':
			return '^'
		case 'v':
			return '>'
		case '<':
			return 'v'
		case '^':
			return '<'
		}
	case 'R':
		switch facing {
		case '>':
			return 'v'
		case 'v':
			return '<'
		case '<':
			return '^'
		case '^':
			return '>'
		}
	}
	panic(fmt.Sprintf("Unable to turn %c from %c", turn, facing))
}

func score(r int, c int, d rune) int {
	fmt.Printf("Ended at %d, %d facing %c\n", r, c, d)
	rowScore := 1000 * r
	colScore := 4 * c
	dirScore := 0
	switch d {
	case '>':
		dirScore = 0
	case 'v':
		dirScore = 1
	case '<':
		dirScore = 2
	case '^':
		dirScore = 3
	default:
		panic(fmt.Sprintf("Do not understand direction %c", d))
	}
	return rowScore + colScore + dirScore
}

type PointAndDir struct {
	X   int
	Y   int
	Dir rune
}

func (p PointAndDir) Opposite() PointAndDir {
	var r rune
	switch p.Dir {
	case '>':
		r = '<'
	case '<':
		r = '>'
	case '^':
		r = 'v'
	case 'v':
		r = '^'
	default:
		panic("Fuck")
	}
	return PointAndDir{p.X, p.Y, r}
}

var cube4Edges map[PointAndDir]PointAndDir
var cube50Edges map[PointAndDir]PointAndDir

func init() {
	cube4Edges = make(map[PointAndDir]PointAndDir)
	cube4Edges[PointAndDir{11, 5, '>'}] = PointAndDir{14, 8, 'v'}
	cube4Edges[PointAndDir{15, 8, '>'}] = PointAndDir{11, 3, '<'}
	cube4Edges[PointAndDir{11, 3, '>'}] = PointAndDir{15, 8, '<'}
	cube4Edges[PointAndDir{10, 11, 'v'}] = PointAndDir{1, 7, '^'}
	cube4Edges[PointAndDir{6, 4, '^'}] = PointAndDir{8, 2, '>'}

	cube50Edges = make(map[PointAndDir]PointAndDir)
}
