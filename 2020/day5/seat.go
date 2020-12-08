package day5

func Max(input []string) int {
	m := 0
	for _, ticket := range input {
		i := seatID(ticket)
		if i > m {
			m = i
		}
	}
	return m
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
