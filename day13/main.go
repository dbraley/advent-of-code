package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

const GOOD = 1
const UNSURE = 0
const BAD = -1

func main() {
	//f := "example.txt"
	f := "input.txt"

	input, err := file.ReadFile("day13/" + f)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", len(input))

	var goodIndexes []int

	lowDivider := packetItem{items: []packetItem{{items: []packetItem{{isNumber: true, num: 2}}}}}
	highDivider := packetItem{items: []packetItem{{items: []packetItem{{isNumber: true, num: 6}}}}}
	sortedPackets := []packetItem{lowDivider, highDivider}

	fmt.Printf("Sorted: %v", sortedPackets)

	for i := 0; i+1 < len(input); i = i + 3 {
		index := i/3 + 1
		fmt.Printf("== Pair %d ==\n", index)

		p0 := packetFromString(input[i])
		p1 := packetFromString(input[i+1])

		if p0.compare(p1, 0) == GOOD {
			goodIndexes = append(goodIndexes, index)
			sortedPackets = merge(sortedPackets, []packetItem{p0, p1})
		} else {
			sortedPackets = merge(sortedPackets, []packetItem{p1, p0})
		}
	}

	fmt.Printf("Good Indexes: %v\n", goodIndexes)
	var sum int
	for _, index := range goodIndexes {
		sum += index
	}
	fmt.Printf("Sum: %d\n", sum)

	fmt.Println("Sorted Set")
	lowIndex := -1
	highIndex := -1
	for i, pi := range sortedPackets {
		if pi.String() == lowDivider.String() {
			lowIndex = i + 1
		}
		if pi.String() == highDivider.String() {
			highIndex = i + 1
		}
		fmt.Println(pi)
	}

	fmt.Printf("Decoder Key: %d * %d = %d\n", lowIndex, highIndex, lowIndex*highIndex)

}

func merge(a, b []packetItem) []packetItem {
	final := []packetItem{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].compare(b[j], 0) > 0 {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final

}

type packetItem struct {
	isNumber bool
	num      int
	items    []packetItem
}

func packetFromString(p string) packetItem {
	ret := packetItem{}
	if p == "" {
		return packetItem{}
	}
	if p[0] == '[' {
		subItemStrs := chunkString(p[1 : len(p)-1])

		for _, subItemStr := range subItemStrs {
			ret.items = append(ret.items, packetFromString(subItemStr))
		}
		return ret
	}
	i, err := strconv.Atoi(p)
	if err != nil {
		panic(fmt.Sprintf("Item %v is neither list nor int", p))
	}
	ret.isNumber = true
	ret.num = i
	return ret
}

func chunkString(p string) []string {
	var subItemStrs []string
	var chunk []rune
	bracketCount := 0
	for _, r := range p {
		switch r {
		case ',':
			if bracketCount == 0 {
				subItemStrs = append(subItemStrs, string(chunk))
				chunk = []rune{}
			} else {
				chunk = append(chunk, r)
			}
		case '[':
			bracketCount++
			chunk = append(chunk, r)
		case ']':
			bracketCount--
			chunk = append(chunk, r)
		default:
			chunk = append(chunk, r)
		}
	}
	subItemStrs = append(subItemStrs, string(chunk))
	return subItemStrs
}

func (pi packetItem) compare(other packetItem, depth int) int {
	fmt.Printf("%v- Compare %v vs %v\n", strings.Repeat("  ", depth), pi, other)
	if pi.isNumber && other.isNumber {
		if pi.num < other.num {
			fmt.Printf("%v- Left side is smaller, so inputs are in the other order\n", strings.Repeat("  ", depth+1))
			return GOOD
		}
		if pi.num == other.num {
			return UNSURE
		}
		fmt.Printf("%v- Right side is smaller, so inputs are NOT in the other order\n", strings.Repeat("  ", depth+1))
		return BAD
	}
	if !pi.isNumber && other.isNumber {
		newOther := packetItem{items: []packetItem{other}}
		fmt.Printf("%v- Mixed types; convert right to %v and retry comparison\n", strings.Repeat("  ", depth+1), newOther)
		return pi.compare(newOther, depth+1)
	}
	if pi.isNumber && !other.isNumber {
		newPi := packetItem{items: []packetItem{pi}}
		fmt.Printf("%v- Mixed types; convert left to %v and retry comparison\n", strings.Repeat("  ", depth+1), newPi)
		return newPi.compare(other, depth+1)
	}
	if !pi.isNumber && !other.isNumber {
		for i, item := range pi.items {
			if i >= len(other.items) {
				fmt.Printf("%v- Right side ran out of items, so inputs are NOT in the right order\n", strings.Repeat("  ", depth+1))
				return BAD
			}
			res := item.compare(other.items[i], depth+1)
			if res == GOOD || res == BAD {
				return res
			}
		}
		if len(pi.items) < len(other.items) {
			fmt.Printf("%v- Left side ran out of items, so inputs are in the right order\n", strings.Repeat("  ", depth+1))
			return GOOD
		}
		return UNSURE
	}
	return UNSURE
}

func (pi packetItem) String() string {
	if pi.isNumber {
		return fmt.Sprintf("%d", pi.num)
	}
	strItems := make([]string, len(pi.items))
	for i, item := range pi.items {
		strItems[i] = item.String()
	}
	return fmt.Sprintf("[%v]", strings.Join(strItems, ","))
}
