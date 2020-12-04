package day3

const tree = '#'

// CountTreesOnPath counts the trees on the 3 over, 1 down path, wrapping around the edges.
func CountTreesOnPath(in []string, right, down int) int {
	treeCount := 0
	y, x := 0, 0
	for {
		y += down
		if y >= len(in) {
			break
		}
		row := in[y]
		x = (x + right) % len(row)
		c := []rune(row)[x]
		if c == tree {
			treeCount = treeCount + 1
		}
	}
	return treeCount
}
