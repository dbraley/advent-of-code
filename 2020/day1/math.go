package day1

import (
	"errors"
)

// ErrorNoValidSum indicates no valid sum was found
var ErrorNoValidSum = errors.New("No Valid Sum")

// FindCommon takes a list of integers and a desired sum, finds the first two ints in the list to add up to that sum. If none do, provides an ErrorNoValidSum
func FindCommon(in []int, sum int) (int, int, error) {
	for i, v1 := range in {
		for _, v2 := range in[i+1:] {
			if v1+v2 == sum {
				return v1, v2, nil
			}
		}
	}
	return 0, 0, ErrorNoValidSum
}

// FindCommon3 takes a list of integers and a desired sum, finds the first three ints in the list to add up to that sum. If none do, provides an ErrorNoValidSum
func FindCommon3(in []int, sum int) (int, int, int, error) {
	// TODO: This code isn't tested, but it's similar enough to FindCommon that I consider it fine for a quick one off. There's some pretty obvious refactoring that could be done, but I'll hold off on that for now.
	for i, v1 := range in {
		// fmt.Printf("Checking (%v) %v\n", i, v1)
		for j, v2 := range in[i+1:] {
			for _, v3 := range in[j+1:] {
				if v1+v2+v3 == sum {
					return v1, v2, v3, nil
				}
			}
		}
	}
	return 0, 0, 0, ErrorNoValidSum
}
