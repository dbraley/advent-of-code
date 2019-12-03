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
