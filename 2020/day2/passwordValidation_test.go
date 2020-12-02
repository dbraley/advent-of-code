package day2

import "testing"

func TestCountValid(t *testing.T) {
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
			gotCnt, gotErr := CountValid(tt.in)
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
			if tt.want != check(tt.c, tt.lb, tt.ub, tt.in) {
				t.Errorf("Expected %v, got %v\n", tt.want, !tt.want)
			}
		})
	}
}
