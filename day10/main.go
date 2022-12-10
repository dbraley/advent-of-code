package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	//f := "example.txt"
	f := "input.txt"

	input, err := file.ReadFile("day10/" + f)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", len(input))

	x := 1
	clock := 0

	sum := 0
	screen := ""

	for _, line := range input {
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "addx":
			screen += render(x, clock)
			clock++
			sum += strength(x, clock)

			screen += render(x, clock)
			clock++
			sum += strength(x, clock)

			v, _ := strconv.Atoi(instruction[1])
			x += v
		case "noop":
			screen += render(x, clock)
			clock++
			sum += strength(x, clock)
		}
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Screen:\n%v", screen)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func strength(x int, clock int) int {
	if (clock+20)%40 == 0 {
		//fmt.Printf("Evaluating strength! Clock %d\tX %d\n", clock, x)
		return clock * x
	}
	return 0
}

func render(x int, clock int) string {
	pixel := "."
	col := clock % 40
	if Abs(x-col) < 2 {
		pixel = "#"
	}
	if col == 0 {
		return "\n" + pixel
	}
	return pixel
}
