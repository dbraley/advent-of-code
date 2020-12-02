package util

import (
	"fmt"
	"io"
	"os"
)

// ReadFileOfInts reads a file of one int per line into an array of ints. Returns an error if there is a problem reading or parsing the file
func ReadFileOfInts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	var line int
	var ret []int
	for {
		_, err := fmt.Fscanln(file, &line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return []int{}, err
		}
		ret = append(ret, line)
	}
	return ret, nil
}
