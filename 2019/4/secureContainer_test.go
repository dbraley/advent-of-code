package main

import (
	"reflect"
	"testing"
)

func Test_toDecArray(t *testing.T) {
	type args struct {
		guess int
	}
	tests := []struct {
		name    string
		args    args
		want    [6]int
		wantErr bool
	}{
		{
			name:    "0",
			args:    args{0},
			want:    [6]int{0, 0, 0, 0, 0, 0},
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{1},
			want:    [6]int{0, 0, 0, 0, 0, 1},
			wantErr: false,
		},
		{
			name:    "21",
			args:    args{21},
			want:    [6]int{0, 0, 0, 0, 2, 1},
			wantErr: false,
		},
		{
			name:    "123456",
			args:    args{123456},
			want:    [6]int{1, 2, 3, 4, 5, 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toDecArray(tt.args.guess)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDecArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDecArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAscending(t *testing.T) {
	type args struct {
		decArray [6]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "000000",
			args: args{decArray: [6]int{0, 0, 0, 0, 0, 0}},
			want: true,
		},
		{
			name: "123456",
			args: args{decArray: [6]int{1, 2, 3, 4, 5, 6}},
			want: true,
		},
		{
			name: "222221",
			args: args{decArray: [6]int{2, 2, 2, 2, 2, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAscending(tt.args.decArray); got != tt.want {
				t.Errorf("isAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsImediateDuplicate(t *testing.T) {
	type args struct {
		guess int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "None",
			args: args{guess: 123456},
			want: false,
		},
		{
			name: "One Duplicate",
			args: args{guess: 123455},
			want: true,
		},
		{
			name: "One Duplicate",
			args: args{guess: 123455},
			want: true,
		},
		{
			name: "One Duplicate",
			args: args{guess: 123455},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsImediateDuplicate(tt.args.guess); got != tt.want {
				t.Errorf("containsImediateDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsImediatePerfectDuplicate(t *testing.T) {
	type args struct {
		guess int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "None",
			args: args{guess: 123456},
			want: false,
		},
		{
			name: "One Duplicate",
			args: args{guess: 123455},
			want: true,
		},
		{
			name: "One Triplet",
			args: args{guess: 123555},
			want: false,
		},
		{
			name: "One Triplet and Dupicate",
			args: args{guess: 111233},
			want: true,
		},
		{
			name: "One Quad and Dupicate",
			args: args{guess: 111133},
			want: true,
		},
		{
			name: "One Quint",
			args: args{guess: 111113},
			want: false,
		},
		{
			name: "Six of the same digit",
			args: args{guess: 111111},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsImediatePerfectDuplicate(tt.args.guess); got != tt.want {
				t.Errorf("containsImediatePerfectDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
