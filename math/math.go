package math

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func WithinOne(a int, b int) bool {
	return Abs(a-b) <= 1
}

type Point2D struct {
	X int
	Y int
}

func (p Point2D) Translate(xVec int, yVec int) Point2D {
	return Point2D{X: p.X + xVec, Y: p.Y + yVec}
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
