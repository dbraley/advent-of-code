package main

import (
	"errors"
	"fmt"
	cu "github.com/dbraley/advent-of-code/colutil"
	"github.com/dbraley/advent-of-code/file"
	"github.com/dbraley/advent-of-code/math"
	"os"
	"strings"
)

func main() {
	//f := "example.txt"
	f := "input.txt"

	input, err := file.ReadFile("day12/" + f)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", len(input))

	startPos, err := findFirst(input, 'S')
	if err != nil {
		fmt.Printf("Error finding start %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Start Position %v\n", startPos)

	endPos, err := findFirst(input, 'E')
	if err != nil {
		fmt.Printf("Error finding end %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("End Position %v\n", endPos)

	path, err := Find(&startPos, &endPos)
	if err != nil {
		fmt.Printf("Error finding path %v\n", err)
		os.Exit(3)
	}

	fmt.Printf("Distance: %d\nPath: %v", len(path)-1, path)
}

type TraversableNodeSet struct {
	NodeMap map[string]AStarNode
}

func (t *TraversableNodeSet) Find(start AStarNode, end AStarNode) ([]*AStarNode, error) {
	openSet := make(cu.PriorityQueue, 1)
	openSet.Push(cu.Item{
		Value:    start.UniqueKey(),
		Priority: int(-end.TravelCost(start)),
		Index:    0,
	})

	cameFrom := make(map[string]AStarNode)
	_ = cameFrom

	gscore := make(map[string]float64)
	gscore[start.UniqueKey()] = 0

	fscore := make(map[string]float64)
	fscore[start.UniqueKey()] = end.TravelCost(start)

	for true {
		if len(openSet) == 0 {
			break
		}

		current := openSet.Pop().(*cu.Item).Value
	}

	return []*AStarNode{}, errors.New("no path found")
}

type AStarNode interface {
	Neighbors() []string
	TravelCost(to AStarNode) float64
	EstimatedPathCost(to AStarNode) float64
	UniqueKey() string
}

type Tile struct {
	point math.Point2D
	field []string
}

func (t *Tile) UniqueKey() string {
	return t.point.String()
}

func (t *Tile) Neighbors() []string {
	//TODO implement me
	panic("implement me")
}

func (t *Tile) TravelCost(to AStarNode) float64 {
	return 1.0
}

func (t *Tile) EstimatedPathCost(to AStarNode) float64 {
	//TODO implement me
	panic("implement me")
}

func findFirst(field []string, searchFor rune) (Tile, error) {
	for y, line := range field {
		x := strings.IndexFunc(line, func(r rune) bool {
			return r == searchFor
		})
		if x != -1 {
			return Tile{math.Point2D{X: x, Y: y}, field}, nil
		}
	}
	errorMsg := fmt.Sprintf("rune %v not found", searchFor)
	return Tile{}, errors.New(errorMsg)
}
