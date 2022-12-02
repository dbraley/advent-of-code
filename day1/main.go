package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
)

func main() {
	in, err := file.ReadFile("day1/input1.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	var max, max2, max3 int
	cur := 0
	// This makes me cringe, but I don't want to refactor it yet, as it might well be YAGNI
	for l, line := range in {
		if line == "" {
			if cur > max {
				fmt.Printf("replacing max %v with %v\n", max, cur)
				max3 = max2
				max2 = max
				max = cur
			} else if cur > max2 {
				fmt.Printf("replacing max2 %v with %v\n", max2, cur)
				max3 = max2
				max2 = cur
			} else if cur > max3 {
				fmt.Printf("replacing max3 %v with %v\n", max3, cur)
				max3 = cur
			}
			fmt.Printf("Maxes: %v %v %v\n", max, max2, max3)
			cur = 0
		} else {
			cal, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Error at line %v: %v\n", l, line)
				os.Exit(2)
			}
			cur += cal
		}
	}

	fmt.Printf("Max: %v\n", max)

	fmt.Printf("Top3: %v + %v + %v = %v", max, max2, max3, max+max2+max3)

}
