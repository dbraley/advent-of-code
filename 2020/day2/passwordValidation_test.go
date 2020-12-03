package day2

import "testing"

func TestCountValidByRange(t *testing.T) {
	tests := []struct {
		name    string
		in      [][]string
		wantCnt int
		wantErr error
	}{
		{
			name: "Invalid Row",
			in: [][]string{
				{"1-3", "a:", "aaa", "bbb"},
			},
			wantCnt: 0,
			wantErr: ErrInvalidRow,
		},
		{
			name: "One Match",
			in: [][]string{
				{"1-3", "a:", "aaa"},
			},
			wantCnt: 1,
			wantErr: nil,
		},
		{
			name: "First 3 Rows",
			in: [][]string{
				{"9-12", "q:", "qqqxhnhdmqqqqjz"},
				{"12-16", "z:", "zzzzzznwlzzjzdzf"},
				{"4-7", "s:", "sssgssw"},
			},
			wantCnt: 1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCnt, gotErr := CountValidByRange(tt.in)
			if tt.wantErr != gotErr {
				t.Errorf("Expected %v, got %w\n", tt.wantErr, gotErr)
			}
			if tt.wantCnt != gotCnt {
				t.Errorf("Expected valid count of %v, got %v\n", tt.wantCnt, gotCnt)
			}
		})
	}
}

func TestCountValidPosition(t *testing.T) {
	tests := []struct {
		name    string
		in      [][]string
		wantCnt int
		wantErr error
	}{
		{
			name: "Invalid Row",
			in: [][]string{
				{"1-3", "a:", "aaa", "bbb"},
			},
			wantCnt: 0,
			wantErr: ErrInvalidRow,
		},
		{
			name: "Matches Neither position",
			in: [][]string{
				{"1-3", "a:", "bbb"},
			},
			wantCnt: 0,
			wantErr: nil,
		},
		{
			name: "Matches First position",
			in: [][]string{
				{"1-3", "a:", "abb"},
			},
			wantCnt: 1,
			wantErr: nil,
		},
		{
			name: "Matches Second position",
			in: [][]string{
				{"1-3", "a:", "bba"},
			},
			wantCnt: 1,
			wantErr: nil,
		},
		{
			name: "Matches Both positions",
			in: [][]string{
				{"1-3", "a:", "aaa"},
			},
			wantCnt: 0,
			wantErr: nil,
		},
		{
			name: "First 3 Rows",
			in: [][]string{
				{"9-12", "q:", "qqqxhnhdmqqqqjz"},   // 9:m, 12:q, valid
				{"12-16", "z:", "zzzzzznwlzzjzdzf"}, // 12: j, 15: f, invalid
				{"4-7", "s:", "sssgssw"},            // 4: g, 7: w, invalid
			},
			wantCnt: 1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCnt, gotErr := CountValidByPosition(tt.in)
			if tt.wantErr != gotErr {
				t.Errorf("Expected %v, got %w\n", tt.wantErr, gotErr)
			}
			if tt.wantCnt != gotCnt {
				t.Errorf("Expected valid count of %v, got %v\n", tt.wantCnt, gotCnt)
			}
		})
	}
}

func Test_parseRange(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantLB  int
		wantUB  int
		wantErr error
	}{
		{
			name:    "1-3",
			in:      "1-3",
			wantLB:  1,
			wantUB:  3,
			wantErr: nil,
		},
		{
			name:    "Double Digit UB",
			in:      "1-12",
			wantLB:  1,
			wantUB:  12,
			wantErr: nil,
		},
		{
			name:    "Double Digit LB",
			in:      "10-12",
			wantLB:  10,
			wantUB:  12,
			wantErr: nil,
		},
		{
			name:    "Bad LB",
			in:      "x-3",
			wantLB:  0,
			wantUB:  0,
			wantErr: ErrInvalidRow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLB, gotUB, gotErr := parseRange(tt.in)
			if tt.wantErr != gotErr {
				t.Errorf("Expected %v, got %v\n", tt.wantErr, gotErr)
			}
			if tt.wantLB != gotLB {
				t.Errorf("Expected Lower Bound %v, got %v\n", tt.wantLB, gotLB)
			}
			if tt.wantUB != gotUB {
				t.Errorf("Expected Upper Bound %v, got %v\n", tt.wantUB, gotUB)
			}
		})
	}
}

func Test_parseChar(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantChr rune
		wantErr error
	}{
		{
			name:    "valid",
			in:      "a:",
			wantChr: 'a',
			wantErr: nil,
		},
		{
			name:    "too small",
			in:      "a",
			wantChr: ' ',
			wantErr: ErrInvalidRow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChr, gotErr := parseChar(tt.in)
			if tt.wantErr != gotErr {
				t.Errorf("Expected %v, got %v\n", tt.wantErr, gotErr)
			}
			if tt.wantChr != gotChr {
				t.Errorf("Expected Rune %v, got %v\n", tt.wantChr, gotChr)
			}

		})

	}
}

func Test_check(t *testing.T) {
	tests := []struct {
		name   string
		c      rune
		lb, ub int
		in     string
		want   bool
	}{
		{
			name: "Basic",
			c:    'a',
			lb:   1,
			ub:   3,
			in:   "aaa",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != checkRange(tt.c, tt.lb, tt.ub, tt.in) {
				t.Errorf("Expected %v, got %v\n", tt.want, !tt.want)
			}
		})
	}
}
