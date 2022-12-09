package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
)

func main() {
	input, err := file.ReadFile("day8/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", len(input))

	total := 0
	bestScore := 0

	for rowNum, line := range input {
		for colNum, r := range line {
			height := convertRuneToInt(r, rowNum, colNum)
			//fmt.Printf("Checking visiblity of %d %d (%d)\n", rowNum, colNum, height)
			visible, score := visibilityAndScore(rowNum, colNum, height, input)
			if visible {
				//fmt.Printf("visible!\n")
				total++
			} else {
				//fmt.Printf("not visible\n")
			}
			if score > bestScore {
				//fmt.Printf("Replacing best score %d with new best %d\n", bestScore, score)
				bestScore = score
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Best Score: %d\n", bestScore)

}

func convertRuneToInt(r rune, rowNum int, colNum int) int {
	height, err := strconv.Atoi(string(r))
	if err != nil {
		fmt.Printf("Tree at %d %d %v cannot be converted to int\n", rowNum, colNum, string(r))
		os.Exit(2)
	}
	return height
}

func visibilityAndScore(row, col, height int, input []string) (bool, int) {
	if row == 0 || col == 0 || row == len(input) || col == len(input[row])-1 {
		return true, 0
	}

	visibleUp, upScore := findUp(row, col, height, input)
	visibleDown, downScore := findDown(row, col, height, input)
	visibleLeft, leftScore := findLeft(row, col, height, input)
	visibleRight, rightScore := findRight(row, col, height, input)

	return visibleUp || visibleDown || visibleLeft || visibleRight,
		upScore * downScore * leftScore * rightScore
}

func findUp(row int, col int, height int, input []string) (bool, int) {
	score := 0
	for i := row - 1; i >= 0; i-- {
		score++
		inQuestion := convertRuneToInt(([]rune(input[i]))[col], i, col)
		if height <= inQuestion {
			//fmt.Printf("Height %d at %d %d is less than %d (up) at %d %d\n", height, row, col, inQuestion, i, col)
			return false, score
		}
	}
	return true, score
}

func findDown(row int, col int, height int, input []string) (bool, int) {
	score := 0
	for i := row + 1; i <= len(input)-1; i++ {
		score++
		inQuestion := convertRuneToInt(([]rune(input[i]))[col], i, col)
		if height <= inQuestion {
			//fmt.Printf("Height %d at %d %d is less than %d (down) at %d %d\n", height, row, col, inQuestion, i, col)
			return false, score
		}
	}
	return true, score
}

func findLeft(row int, col int, height int, input []string) (bool, int) {
	score := 0
	for i := col - 1; i >= 0; i-- {
		score++
		inQuestion := convertRuneToInt(([]rune(input[row]))[i], row, i)
		if height <= inQuestion {
			//fmt.Printf("Height %d at %d %d is less than %d (left) at %d %d\n", height, row, col, inQuestion, row, i)
			return false, score
		}
	}
	return true, score
}

func findRight(row int, col int, height int, input []string) (bool, int) {
	score := 0
	for i := col + 1; i <= len(input[row])-1; i++ {
		score++
		inQuestion := convertRuneToInt(([]rune(input[row]))[i], row, i)
		if height <= inQuestion {
			//fmt.Printf("Height %d at %d %d is less than %d (right) at %d %d\n", height, row, col, inQuestion, row, i)
			return false, score
		}
	}
	return true, score
}
