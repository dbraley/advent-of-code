package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strings"
)

func main() {
	input, err := file.ReadFile("day6/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	signal := input[0]

	fmt.Printf("input: %v (%d)", signal, len(signal))

	//markerSize := 4
	markerSize := 14
	for i := markerSize; i <= len(signal); i++ {
		maybeMarker := signal[i-markerSize : i]
		//fmt.Println(maybeMarker)
		isMarker := true
		for j, r := range maybeMarker {
			if strings.Contains(maybeMarker[j+1:], string(r)) {
				fmt.Printf("index %d:%v is NOT a marker because %v is in %v\n", i, maybeMarker, string(r), maybeMarker[j+1:])
				isMarker = false
				break
			}
		}
		if isMarker == true {
			fmt.Printf("Found marker at index %d:%v\n", i, maybeMarker)
			break
		}
	}

}
