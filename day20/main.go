package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Example:")
	ept1, ept2 := day20("example.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day20("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day20(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day20/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	e1, e2, e3 := findCoords(input, 1, 1)

	pt1 = e1.val + e2.val + e3.val

	e21, e22, e23 := findCoords(input, 811589153, 10)

	pt2 = e21.val + e22.val + e23.val

	return pt1, pt2
}

func findCoords(input []string, decryptionKey int, mixCount int) (*entry, *entry, *entry) {
	entryMap := make(map[int]*entry)
	var zeroEntry *entry

	var last *entry

	for i, line := range input {
		n, _ := strconv.Atoi(line)
		e := newEntry(n * decryptionKey)
		entryMap[i] = e
		if n == 0 {
			zeroEntry = e
		}
		e.linkLeft(last)
		last = e
	}
	entryMap[0].linkLeft(last)

	//fmt.Printf("Zero entry: %v\n", zeroEntry)

	fmt.Printf("List: %v\n", zeroEntry.SprintEntries(7))

	for j := 0; j < mixCount; j++ {
		for i := 0; i < len(entryMap); i++ {
			e := entryMap[i]
			//fmt.Printf("Moving %v\n", e)
			if e.val != 0 {
				// Join old left and right
				e.right.linkLeft(e.left)
				var newLeft, newRight *entry
				move := e.val % (len(entryMap) - 1)
				if move > 0 {
					// going right
					newLeft = e.GetFromRight(move)
					newRight = newLeft.right
				} else {
					newRight = e.GetFromLeft(move * -1)
					newLeft = newRight.left
				}
				// Join to new Right
				newRight.linkLeft(e)
				e.linkLeft(newLeft)
			}
			//fmt.Printf("List: %v\n", zeroEntry.SprintEntries(7))
		}
		//fmt.Printf("List: %v\n", zeroEntry.SprintEntries(7))
	}

	e1 := zeroEntry.GetFromRight(1000)
	e2 := e1.GetFromRight(1000)
	e3 := e2.GetFromRight(1000)
	return e1, e2, e3
}

type entry struct {
	val   int
	left  *entry
	right *entry
}

func newEntry(i int) *entry {
	return &entry{i, nil, nil}
}

func (e *entry) String() string {
	l := "nil"
	if e.left != nil {
		l = strconv.Itoa(e.left.val)
	}
	r := "nil"
	if e.right != nil {
		r = strconv.Itoa(e.right.val)
	}
	return fmt.Sprintf("<-%v (%d) %v->", l, e.val, r)

}

func (e *entry) linkLeft(l *entry) {
	if l != nil {
		e.left = l
		l.right = e
	}
}

func (e *entry) SprintEntries(i int) string {
	if i <= 1 {
		return fmt.Sprintf("%d", e.val)
	}
	if e.right == nil {
		panic(fmt.Sprintf("nil right entry found at %v", e))
	}
	return fmt.Sprintf("%d, %v", e.val, e.right.SprintEntries(i-1))
}

func (e *entry) GetFromRight(i int) *entry {
	if i <= 0 {
		return e
	}
	if e.right == nil {
		panic(fmt.Sprintf("nil right entry found at %v", e))
	}
	return e.right.GetFromRight(i - 1)
}
func (e *entry) GetFromLeft(i int) *entry {
	if i <= 0 {
		return e
	}
	if e.left == nil {
		panic(fmt.Sprintf("nil left entry found at %v", e))
	}
	return e.left.GetFromLeft(i - 1)
}
