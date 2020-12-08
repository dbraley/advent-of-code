package day5

// Find identifies the missing seat ID. Returns the max as well, for checksum reasons or something.
func Find(input []string) (int, int) {
	max, min := 0, 1000000
	var sum int // int64?
	for _, ticket := range input {
		i := seatID(ticket)
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
		sum = sum + i
	}
	// fmt.Printf("min: %v\nmax: %v\nsum: %v\n", min, max, sum)
	fullSum := ((max - min + 1) * (max + min)) >> 1
	// fmt.Printf("Full sum: %v\nMissing: %v\n", fullSum, fullSum-sum)
	return max, fullSum - sum
}

func seatID(ticket string) int {
	row, col := rowAndCol(ticket)
	return row*8 + col
}

func rowAndCol(ticket string) (int, int) {
	row, col := 0, 0
	for i, rune := range []rune(ticket) {
		if i < 7 {
			row = row << 1
			if rune == 'B' {
				row = row | 1
			}
		} else {
			col = col << 1
			if rune == 'R' {
				col = col | 1
			}
		}
	}
	return row, col
}
