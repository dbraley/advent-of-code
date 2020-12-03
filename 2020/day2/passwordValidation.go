package day2

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ErrInvalidRow suggests that a row was invalid
var ErrInvalidRow = errors.New("Invalid Row")

// CountValidByRange counts the valid number of passwords in an array of rule and
// password combinations, where each one of those combinations looks like:
// 1-3 a: abcde
// The above is a valid rule/password combination if the letter a appears
// between one and 3 times in the password abcde. In this case, it appears
// once, and is therefor valid
func CountValidByRange(in [][]string) (int, error) {
	var count = 0
	for _, rule := range in {
		if len(rule) != 3 {
			return 0, ErrInvalidRow
		}

		lb, ub, err := parseRange(rule[0])
		if err != nil {
			return 0, err
		}

		c, err := parseChar(rule[1])
		if err != nil {
			return 0, err
		}

		if checkRange(c, lb, ub, rule[2]) {
			count = count + 1
		}
		// fmt.Println(rule, string(c), lb, ub, rule[2], count)
	}
	return count, nil
}

// CountValidByPosition counts the valid number of passwords in an array of rule and
// password combinations, where each one of those combinations looks like:
// 1-3 a: abcde
// The above is a valid rule/password combination if the letter a appears
// in exactly one of the available posistions 1 & 3. Tha character `a` apperas
// appears at position 1, but not 3
func CountValidByPosition(in [][]string) (int, error) {
	var count = 0
	for _, rule := range in {
		if len(rule) != 3 {
			return 0, ErrInvalidRow
		}

		lb, ub, err := parseRange(rule[0])
		if err != nil {
			return 0, err
		}

		c, err := parseChar(rule[1])
		if err != nil {
			return 0, err
		}

		if checkPosition(c, lb, ub, rule[2]) {
			count = count + 1
		}
		// fmt.Println(rule, string(c), lb, ub, rule[2], count)
	}
	return count, nil
}

func parseRange(in string) (int, int, error) {
	re := regexp.MustCompile(`(\d+)\-(\d+)`)
	res := re.FindStringSubmatch(in)
	if len(res) == 3 {
		// There's a better way to do this, but I'm not rememberring it at the moment...
		lb, err := strconv.Atoi(res[1])
		if err != nil {
			return 0, 0, ErrInvalidRow
		}
		ub, err := strconv.Atoi(res[2])
		if err != nil {
			return 0, 0, ErrInvalidRow
		}
		return lb, ub, nil
	}
	return 0, 0, ErrInvalidRow
}

func parseChar(in string) (rune, error) {
	if len(in) != 2 {
		return ' ', ErrInvalidRow
	}
	return []rune(in)[0], nil
}

func checkRange(c rune, lb, ub int, in string) bool {
	count := strings.Count(in, string(c))
	return count >= lb && count <= ub
}

func checkPosition(c rune, lb, ub int, in string) bool {
	runes := []rune(in)
	// TODO: this should probably return an error, but I suspect the input data won't actually hit this condition so meh
	if len(runes) < ub {
		fmt.Printf("Error! input string %q is less than %v runes long!", in, ub)
		return false
	}
	first := runes[lb-1]
	second := runes[ub-1]
	return (first == c && second != c) || (first != c && second == c)
}
