package main

import (
	"reflect"
	"testing"
)

func TestShouldContinue(t *testing.T) {
	type args struct {
		program     []int
		instruction int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "OutOfBounds Negative",
			args: args{
				program:     []int{0},
				instruction: -1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "OutOfBounds Too High",
			args: args{
				program:     []int{0},
				instruction: 4,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Not done",
			args: args{
				program:     []int{1, 0, 0, 0, 99},
				instruction: 0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Done",
			args: args{
				program:     []int{1, 0, 0, 0, 99},
				instruction: 4,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShouldContinue(tt.args.program, tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldContinue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShouldContinue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteInstruction(t *testing.T) {
	type args struct {
		program     []int
		instruction int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Basic Addition",
			args: args{
				program:     []int{1, 0, 0, 0, 99},
				instruction: 0,
			},
			want:    []int{2, 0, 0, 0, 99},
			wantErr: false,
		},
		{
			name: "Basic Multiplication",
			args: args{
				program:     []int{2, 3, 0, 3, 99},
				instruction: 0,
			},
			want:    []int{2, 3, 0, 6, 99},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecuteInstruction(tt.args.program, tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteInstruction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteInstruction2(t *testing.T) {
	type args struct {
		program mem
		index   int
	}
	tests := []struct {
		name      string
		args      args
		want      []int
		nextIndex int
		wantErr   bool
	}{
		{
			name: "Basic Addition",
			args: args{
				program: []int{1, 0, 0, 0, 99},
				index:   0,
			},
			want:      []int{2, 0, 0, 0, 99},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Basic Addition",
			args: args{
				program: []int{1, 0, 0, 0, 99},
				index:   0,
			},
			want:      []int{2, 0, 0, 0, 99},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Basic Multiplication",
			args: args{
				program: []int{2, 3, 0, 3, 99},
				index:   0,
			},
			want:      []int{2, 3, 0, 6, 99},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Basic Multiplication with immediate mode",
			args: args{
				program: []int{1002, 4, 3, 4, 33},
				index:   0,
			},
			want:      []int{1002, 4, 3, 4, 99},
			nextIndex: 4,
			wantErr:   false,
		},
		//	did some q$d hardcoding so opcode 3 and 4 aren't really testable
		{
			name: "Jump if True - false",
			args: args{
				program: []int{1105, 0, 0, 99, 99},
				index:   0,
			},
			want:      []int{1105, 0, 0, 99, 99},
			nextIndex: 3,
			wantErr:   false,
		},
		{
			name: "Jump if True - true",
			args: args{
				program: []int{1105, 1, 0, 99, 99},
				index:   0,
			},
			want:      []int{1105, 1, 0, 99, 99},
			nextIndex: 0,
			wantErr:   false,
		},
		{
			name: "Jump if False - false",
			args: args{
				program: []int{1106, 0, 0, 99, 99},
				index:   0,
			},
			want:      []int{1106, 0, 0, 99, 99},
			nextIndex: 0,
			wantErr:   false,
		},
		{
			name: "Jump if False - true",
			args: args{
				program: []int{1106, 1, 0, 99, 99},
				index:   0,
			},
			want:      []int{1106, 1, 0, 99, 99},
			nextIndex: 3,
			wantErr:   false,
		},
		{
			name: "Less Than - true",
			args: args{
				program: []int{1107, 1, 2, 5, 99, 2},
				index:   0,
			},
			want:      []int{1107, 1, 2, 5, 99, 1},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Less Than - false",
			args: args{
				program: []int{1107, 2, 2, 5, 99, 2},
				index:   0,
			},
			want:      []int{1107, 2, 2, 5, 99, 0},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Equal to - true",
			args: args{
				program: []int{1108, 1, 1, 5, 99, 2},
				index:   0,
			},
			want:      []int{1108, 1, 1, 5, 99, 1},
			nextIndex: 4,
			wantErr:   false,
		},
		{
			name: "Equal to - false",
			args: args{
				program: []int{1108, 1, 2, 5, 99, 2},
				index:   0,
			},
			want:      []int{1108, 1, 2, 5, 99, 0},
			nextIndex: 4,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotNextIndex, err := ExecuteInstruction2(tt.args.program, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteInstruction2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteInstruction2() got = %v, want %v", got, tt.want)
			}
			if gotNextIndex != tt.nextIndex {
				t.Errorf("ExecuteInstruction2() gotNextIndex = %v, want %v", gotNextIndex, tt.nextIndex)
			}
		})
	}
}
