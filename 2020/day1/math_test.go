package day1

import (
	"testing"
)

func TestFindCommon(t *testing.T) {
	tests := []struct {
		name  string
		in    []int
		sum   int
		want1 int
		want2 int
		wantE error
	}{
		{
			name:  "Only 2",
			in:    []int{1, 2},
			sum:   3,
			want1: 1,
			want2: 2,
			wantE: nil,
		},
		{
			name:  "Randomly placed",
			in:    []int{0, 1, 0, 2, 0},
			sum:   3,
			want1: 1,
			want2: 2,
			wantE: nil,
		},
		{
			name:  "No Valid combo",
			in:    []int{1, 1},
			sum:   3,
			want1: 0,
			want2: 0,
			wantE: ErrorNoValidSum,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, gotE := FindCommon(tt.in, tt.sum)
			if tt.wantE != gotE {
				t.Errorf("Expected %w, got %w", tt.wantE, gotE)
			}
			switch tt.want1 {
			case got1:
				if tt.want2 != got2 {
					t.Errorf("Expected %v %v, got %v %v", tt.want1, tt.want2, got1, got2)
				}
			case got2:
				if tt.want2 != got1 {
					t.Errorf("Expected %v %v, got %v %v", tt.want1, tt.want2, got1, got2)
				}
			default:
				t.Errorf("Expected %v %v, got %v %v", tt.want1, tt.want2, got1, got2)
			}
		})
	}
}
