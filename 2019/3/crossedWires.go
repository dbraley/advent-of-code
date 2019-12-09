package main

import (
	"fmt"
	"os"
	"strconv"
)
import "encoding/csv"

func main() {
	strings, err := Read("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	var wires [][]Line
	for _, vectors := range strings {
		lines, err := ToLines(vectors)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wires = append(wires, lines)
	}
	//fmt.Println(wires)
	if len(wires) != 2 {
		fmt.Println("Expected 2 wires, got", len(wires))
		return
	}

	var intersections []Coordinate

	for _, l1 := range wires[0] {
		for _, l2 := range wires[1] {
			intersections = append(intersections, FindIntersections(l1, l2)...)
		}
	}

	fmt.Println(intersections)

	best := intersections[0]
	smallest := intersections[0]
	for _, coord := range intersections {
		if coord.Manhattan() < best.Manhattan() {
			best = coord
		}
		if coord.dist < smallest.dist {
			smallest = coord
		}
	}

	fmt.Println(best, best.Manhattan())
	fmt.Println(smallest)
}

func Read(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

type Line struct {
	x0, y0, x1, y1 int
	l0, l1 int
}

func ToLines(vectors []string) ([]Line, error) {
	x0, y0 := 0, 0
	l0 := 0
	var lines []Line
	for _, vector := range vectors {

		mag, err := strconv.Atoi(vector[1:])
		if err != nil {
			return []Line{}, err
		}

		var newLine Line

		switch vector[0] {
		case 'R':
			newLine = Line{x0, y0, x0 + mag, y0, l0, l0 + mag}
		case 'L':
			newLine = Line{x0, y0, x0 - mag, y0, l0, l0 + mag}
		case 'U':
			newLine = Line{x0, y0, x0, y0 + mag, l0, l0 + mag}
		case 'D':
			newLine = Line{x0, y0, x0, y0 - mag, l0, l0 + mag}
		default:
			fmt.Println("Didn't understand", vector)
			return []Line{}, fmt.Errorf("unknown direction %v", vector[0])
		}

		x0, y0 = newLine.x1, newLine.y1
		l0 = newLine.l1
		lines = append(lines, newLine)
	}
	return lines, nil
}

type Coordinate struct {
	x, y int
	dist int
}

func (c Coordinate) Manhattan() int {
	return abs(c.x) + abs(c.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (l Line) isVertical() bool {
	return l.x0 == l.x1
}

func (l Line) isHorizontal() bool {
	return l.y0 == l.y1
}

func FindIntersections(l1, l2 Line) []Coordinate {
	intersections := []Coordinate{}
	if l1.isHorizontal() && l2.isVertical() {
		if overlapsOnX(l1, l2) && overlapsOnY(l1, l2) {
			intersections = append(intersections, Coordinate{l2.x0, l1.y0,
				l1.l0 + abs(l1.x0 - l2.x0) + l2.l0 + abs(l2.y0 - l1.y0)})
		}
	}
	if l2.isHorizontal() && l1.isVertical() {
		if overlapsOnX(l2, l1) && overlapsOnY(l2, l1) {
			intersections = append(intersections, Coordinate{l1.x0, l2.y0,
				l1.l0 + abs(l1.y0 - l2.y0) + l2.l0 + abs(l2.x0 - l1.x0)})
		}
	}
	return intersections
}

func overlapsOnX(horizontal Line, vertical Line) bool {
	if horizontal.x0 < horizontal.x1 {
		return horizontal.x0 <= vertical.x0 && vertical.x0 <= horizontal.x1
	}
	return horizontal.x0 >= vertical.x0 && vertical.x0 >= horizontal.x1
}

func overlapsOnY(horizontal Line, vertical Line) bool {
	if vertical.y0 < vertical.y1 {
		return vertical.y0 <= horizontal.y0 && horizontal.y0 <= vertical.y1
	}
	return vertical.y0 >= horizontal.y0 && horizontal.y0 >= vertical.y1
}
