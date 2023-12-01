package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
)

func main() {
	fmt.Println("Trivial:")
	tpt1, tpt2 := day25("trivial.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", tpt1, tpt2)

	fmt.Println("Example:")
	ept1, ept2 := day25("example.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day25("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day25(fileName string) (pt1 int, pt2 string) {
	input, err := file.ReadFile("day25/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	sum := 0
	for _, line := range input {
		val := 0
		for _, char := range line {
			val *= 5
			switch char {
			case '0': // nothing
			case '1':
				val += 1
			case '2':
				val += 2
			case '-':
				val -= 1
			case '=':
				val -= 2
			default:
				panic("wut")
			}
		}
		sum += val
	}

	num := sum
	snafu := []rune{}
	for true {
		if num <= 0 {
			break
		}
		switch num % 5 {
		case 0:
			snafu = append([]rune{'0'}, snafu...)
		case 1:
			snafu = append([]rune{'1'}, snafu...)
		case 2:
			snafu = append([]rune{'2'}, snafu...)
		case 3:
			snafu = append([]rune{'='}, snafu...)
		case 4:
			snafu = append([]rune{'-'}, snafu...)
		}
		num = (num + 2) / 5
	}

	return sum, string(snafu)
}
