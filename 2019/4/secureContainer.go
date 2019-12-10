package main

import "fmt"

const min = 372037
const max = 905157

func main() {
	possibilities := []int{}
	possibilities2 := []int{}

	for guess := min; guess <= max; guess++ {
		decArray, err := toDecArray(guess)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if isAscending(decArray) && containsImediateDuplicate(guess) {
			possibilities = append(possibilities, guess)
		}
		if isAscending(decArray) && containsImediatePerfectDuplicate(guess) {
			possibilities2 = append(possibilities2, guess)
		}
	}

	fmt.Println(len(possibilities))
	fmt.Println(len(possibilities2))
	// // // for _, v := range possibilities {
	// // 	fmt.Printf("%v,", v)
	// }
	// fmt.Println(possibilities)

}

func toDecArray(guess int) ([6]int, error) {
	decArray := [6]int{}
	if guess < 0 || guess > 999999 {
		return decArray, fmt.Errorf("invalid guess %v", guess)
	}
	for i := 5; i >= 0; i-- {
		decArray[i] = guess % 10
		guess = guess / 10
	}
	return decArray, nil
}

func isAscending(decArray [6]int) bool {
	previousValue := -1
	for _, val := range decArray {
		if previousValue > val {
			return false
		}
		previousValue = val
	}
	return true
}

func containsImediateDuplicate(guess int) bool {
	for guess > 0 {
		if (guess%100)%11 == 0 {
			return true
		}
		guess = guess / 10
	}
	return false
}

func containsImediatePerfectDuplicate(guess int) bool {
	for guess > 0 {
		if (guess%1000000)%111111 == 0 {
			return false
		}
		if (guess%100000)%11111 == 0 {
			return false
		}
		if (guess%10000)%1111 == 0 {
			guess = guess / 1000
			continue
		}
		if (guess%1000)%111 == 0 {
			guess = guess / 100
			continue
		}
		if (guess%100)%11 == 0 {
			fmt.Println()
			return true
		}
		guess = guess / 10
	}
	return false
}
