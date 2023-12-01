package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Example:")
	ept1, ept2 := day18("example.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day18("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day18(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day18/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	blocks := make([]Block, len(input))
	// ////////////// x ///// y ///// z / type
	grid := make(map[int]map[int]map[int]string)

	minX, minY, minZ := 100, 100, 100
	maxX, maxY, maxZ := 0, 0, 0

	for i, b := range input {
		bxyz := strings.Split(b, ",")
		x, _ := strconv.Atoi(bxyz[0])
		y, _ := strconv.Atoi(bxyz[1])
		z, _ := strconv.Atoi(bxyz[2])

		minX, maxX = min(minX, x-1), max(maxX, x+1)
		minY, maxY = min(minY, y-1), max(maxY, y+1)
		minZ, maxZ = min(minZ, z-1), max(maxZ, z+1)

		set(grid, x, y, z, "rock")
		setIfUnset(grid, x-1, y, z, "air")
		setIfUnset(grid, x+1, y, z, "air")
		setIfUnset(grid, x, y-1, z, "air")
		setIfUnset(grid, x, y+1, z, "air")
		setIfUnset(grid, x, y, z-1, "air")
		setIfUnset(grid, x, y, z+1, "air")

		blocks[i] = Block{
			X:   x,
			Y:   y,
			Z:   z,
			Mat: "rock",
		}

	}

	for _, block := range blocks {
		pt1 += block.ExposedSurface(grid)
	}

	maybeBubbles := []*Block{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				at := materialAt(grid, x, y, z)
				if at == "air" || at == "unknown" {
					setIfUnset(grid, x, y, z, "air")
					maybeBubbles = append(maybeBubbles, &Block{
						X:   x,
						Y:   y,
						Z:   z,
						Mat: "air",
					})
				}
			}
		}
	}

	for i := 0; i < 30; i++ {
		didSomething := false
		for _, mb := range maybeBubbles {
			if mb.Mat == "air" {
				ns := mb.Neighbors(grid)
				for _, n := range ns {
					if n.Mat == "unknown" || n.Mat == "free-air" {
						didSomething = true
						set(grid, mb.X, mb.Y, mb.Z, "free-air")
						mb.Mat = "free-air"
						break
					}
				}
			}
		}
		if didSomething == false {
			break
		}
	}

	bubbleSurface := 0
	for _, mb := range maybeBubbles {
		// Actually a bubble now
		if mb.Mat == "air" {
			bubbleSurface += mb.ExposedSurface(grid)
			mb.Mat = "bubble"
			set(grid, mb.X, mb.Y, mb.Z, "bubble")
		}
	}

	//pt2 = pt1 - bubbleSurface
	for _, block := range blocks {
		pt2 += block.ExposedSurface(grid)
	}

	fmt.Println("Display")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%2d   01234567890123456789\n", x)
		for y := minY; y <= maxY; y++ {
			fmt.Printf("%2d: ", y)
			for z := minZ; z <= maxZ; z++ {
				switch materialAt(grid, x, y, z) {
				case "rock":
					b := blockAt(grid, x, y, z)
					fmt.Printf("%d", b.ExposedSurface(grid))
				case "air":
					fmt.Printf("X")
				case "bubble":
					fmt.Printf("o")
				case "free-air":
					fmt.Printf("~")
				case "unknown":
					fmt.Printf(" ")
				default:
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return pt1, pt2
}

func set(grid map[int]map[int]map[int]string, x, y, z int, material string) {
	if _, ok := grid[x]; !ok {
		grid[x] = make(map[int]map[int]string)
	}
	if _, ok := grid[x][y]; !ok {
		grid[x][y] = make(map[int]string)
	}
	if old, ok := grid[x][y][z]; ok {
		fmt.Printf("Replacing (%d, %d, %d) %v with %v\n", x, y, z, old, material)
	}
	grid[x][y][z] = material
}

func setIfUnset(grid map[int]map[int]map[int]string, x, y, z int, material string) {
	if materialAt(grid, x, y, z) == "unknown" {
		set(grid, x, y, z, material)
	}
}

func materialAt(grid map[int]map[int]map[int]string, x, y, z int) string {
	if m, ok := grid[x][y][z]; ok {
		return m
	}
	return "unknown"
}
func blockAt(grid map[int]map[int]map[int]string, x, y, z int) Block {
	return Block{x, y, z, materialAt(grid, x, y, z)}
}

type Block struct {
	X   int
	Y   int
	Z   int
	Mat string
}

func (b *Block) String() string {
	return fmt.Sprintf("((%d,%d,%d), %q)", b.X, b.Y, b.Z, b.Mat)
}

func (b *Block) ExposedSurface(grid map[int]map[int]map[int]string) int {
	sa := 0
	if b.isEdge(materialAt(grid, b.X-1, b.Y, b.Z)) {
		sa++
	}
	if b.isEdge(materialAt(grid, b.X+1, b.Y, b.Z)) {
		sa++
	}
	if b.isEdge(materialAt(grid, b.X, b.Y-1, b.Z)) {
		sa++
	}
	if b.isEdge(materialAt(grid, b.X, b.Y+1, b.Z)) {
		sa++
	}
	if b.isEdge(materialAt(grid, b.X, b.Y, b.Z-1)) {
		sa++
	}
	if b.isEdge(materialAt(grid, b.X, b.Y, b.Z+1)) {
		sa++
	}
	return sa
}

func (b *Block) isEdge(material string) bool {
	if b.Mat == "rock" {
		switch material {
		case "rock":
			return false
		case "air":
			return true
		case "free-air":
			return true
		case "unknown":
			return true
		case "bubble":
			return false
		default:
			return true
		}
	}
	return b.Mat == material
}

func (b *Block) Neighbors(grid map[int]map[int]map[int]string) []Block {
	ret := make([]Block, 6)
	ret[4] = blockAt(grid, b.X-1, b.Y, b.Z)
	ret[5] = blockAt(grid, b.X+1, b.Y, b.Z)
	ret[2] = blockAt(grid, b.X, b.Y-1, b.Z)
	ret[3] = blockAt(grid, b.X, b.Y+1, b.Z)
	ret[0] = blockAt(grid, b.X, b.Y, b.Z-1)
	ret[1] = blockAt(grid, b.X, b.Y, b.Z+1)
	return ret
}

// TODO: Move to math
func min(a int, b ...int) int {
	m := a
	for _, other := range b {
		if other < m {
			m = other
		}
	}
	return m
}
func max(a int, b ...int) int {
	m := a
	for _, other := range b {
		if other > m {
			m = other
		}
	}
	return m
}
