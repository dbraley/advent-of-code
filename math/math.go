package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func WithinOne(a int, b int) bool {
	return Abs(a-b) <= 1
}
