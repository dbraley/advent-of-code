package main

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				fileName: "testdata/basic",
			},
			want:    [][]string{{"R1"}, {"L1"}},
			wantErr: false,
		},
		{
			name: "Multiple",
			args: args{
				fileName: "testdata/multiple",
			},
			want:    [][]string{{"R1", "U1"}, {"L1", "D1"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLines(t *testing.T) {
	type args struct {
		dirs []string
	}
	tests := []struct {
		name string
		args args
		want []Line
	}{
		{
			name: "Right",
			args: args{
				dirs: []string{"R1"},
			},
			want: []Line{
				{
					x0: 0,
					y0: 0,
					x1: 1,
					y1: 0,
					l0: 0,
					l1: 1,
				},
			},
		},
		{
			name: "R1,U2,L3,D4",
			args: args{
				dirs: []string{"R1", "U2", "L3", "D4"},
			},
			want: []Line{
				{0, 0, 1, 0, 0, 1},
				{1, 0, 1, 2, 1, 3},
				{1, 2, -2, 2, 3, 6},
				{-2, 2, -2, -2, 6, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ToLines(tt.args.dirs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		l1 Line
		l2 Line
	}
	tests := []struct {
		name string
		args args
		want []Coordinate
	}{
		{
			name: "Parallel",
			args: args{
				l1: Line{0, 0, 2, 0, 0, 0},
				l2: Line{0, 1, 2, 1, 0, 0},
			},
			want: []Coordinate{},
		},
		{
			name: "Cross",
			args: args{
				l1: Line{-2, 0, 2, 0, 0, 0},
				l2: Line{0, -1, 0, 1, 0, 0},
			},
			want: []Coordinate{{0, 0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIntersections(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}
