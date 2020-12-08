package day5

import (
	"testing"
)

func TestRowAndCol(t *testing.T) {
	tests := []struct {
		ticket  string
		wantRow int
		wantCol int
	}{
		{
			ticket:  "FFFFFFFLLL",
			wantRow: 0,
			wantCol: 0,
		},
		{
			ticket:  "FFFFFFFLLR",
			wantRow: 0,
			wantCol: 1,
		},
		{
			ticket:  "FFFFFFBLLR",
			wantRow: 1,
			wantCol: 1,
		},
		{
			ticket:  "BFFFFFFRLL",
			wantRow: 64,
			wantCol: 4,
		},
		{
			ticket:  "BBBBBBBRRR",
			wantRow: 127,
			wantCol: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.ticket, func(t *testing.T) {
			gotRow, gotCol := rowAndCol(tt.ticket)
			sameInt(t, "row", tt.wantRow, gotRow)
			sameInt(t, "col", tt.wantCol, gotCol)
		})
	}
}

func sameInt(t *testing.T, name string, want, got int) {
	if want != got {
		t.Errorf("Expected %s %v, got %v", name, want, got)
	}
}
