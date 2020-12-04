package day3

const tree = '#'

// CountTreesOnPath counts the trees on the 3 over, 1 down path, wrapping around the edges.
func CountTreesOnPath(in []string) int {
	treeCount := 0
	for i, row := range in {
		offset := (i * 3) % len(row)
		// fmt.Println(i, offset)
		c := []rune(row)[offset]
		// fmt.Println(i, string(c))
		if c == tree {
			treeCount = treeCount + 1
		}
	}
	return treeCount
}
