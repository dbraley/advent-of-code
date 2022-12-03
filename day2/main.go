package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
)

func main() {
	in, err := file.ReadSSV("day2/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	scoreFirstWay := 0
	scoreSecondWay := 0

	for _, match := range in {
		matchScoreFirstWay := 0
		matchScoreSecondWay := 0
		switch match[1] {
		case "X": // Rock(*) OR Lose(**)
			matchScoreFirstWay += 1
			matchScoreSecondWay = 0

			switch match[0] {
			case "A": // Rock
				// tie(*)
				matchScoreFirstWay += 3
				// (**) To Lose to rock, we pick Scissors
				matchScoreSecondWay += 3
			case "B": // Paper
				// lose(*)
				// (**) To Lose to Paper, we pick Rock
				matchScoreSecondWay += 1
			case "C": // Scissors
				// win(*)
				matchScoreFirstWay += 6
				// (**) To Lose to Scissors, we pick Paper
				matchScoreSecondWay += 2
			}
		case "Y": // Paper(*) OR Tie(**)
			matchScoreFirstWay += 2
			matchScoreSecondWay = 3

			switch match[0] {
			case "A": // Rock
				// win(*)
				matchScoreFirstWay += 6
				matchScoreSecondWay += 1
			case "B": // Paper
				// tie(*)
				matchScoreFirstWay += 3
				matchScoreSecondWay += 2
			case "C": // Scissors
				// lose(*)
				matchScoreSecondWay += 3
			}
		case "Z": // Scissors(*) OR Win(3)
			matchScoreFirstWay += 3
			matchScoreSecondWay = 6

			switch match[0] {
			case "A": // Rock
				// lose(*)
				// (**) To Win to rock, we pick Paper
				matchScoreSecondWay += 2
			case "B": // Paper
				// win(*)
				matchScoreFirstWay += 6
				// (**) To Win to Paper, we pick scissors
				matchScoreSecondWay += 3
			case "C": // Scissors
				// tie(*)
				matchScoreFirstWay += 3
				// (**) To Win to scissors, we pick rock
				matchScoreSecondWay += 1
			}
		}
		fmt.Printf("%v vs %v -> %d (*) OR %d (**)\n", match[0], match[1], matchScoreFirstWay, matchScoreSecondWay)
		scoreFirstWay += matchScoreFirstWay
		scoreSecondWay += matchScoreSecondWay
	}

	fmt.Printf("Final Score: %v (*) OR %d (**)", scoreFirstWay, scoreSecondWay)
}
