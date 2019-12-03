package main

import "testing"

func TestFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"12",
			args:args{mass:12},
			want:2,
		},
		{
			name:"14",
			args:args{mass:14},
			want:2,
		},
		{
			name:"1969",
			args:args{mass:1969},
			want:654,
		},
		{
			name:"100756",
			args:args{mass:100756},
			want:33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fuel(tt.args.mass); got != tt.want {
				t.Errorf("Fuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformAndSum(t *testing.T) {
	type args struct {
		masses []int
		transform func(int)int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"12+14",
			args:args{masses:[]int{12, 14}, transform:Fuel},
			want:4,
		},
		{
			name:"12+14+1969+100756",
			args:args{masses:[]int{12, 14, 1969, 100756},transform:Fuel},
			want:2+2+654+33583,
		},
		{
			name:"12+14+1969+100756",
			args:args{masses:[]int{12, 14, 1969, 100756},transform:Fuel2},
			want:2+2+966+50346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformAndSum(tt.args.masses, tt.args.transform); got != tt.want {
				t.Errorf("TransformAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuel2(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"14",
			args:args{mass:14},
			want:2,
		},
		{
			name:"1969",
			args:args{mass:1969},
			want:654+216+70+21+5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fuel2(tt.args.mass); got != tt.want {
				t.Errorf("Fuel2() = %v, want %v", got, tt.want)
			}
		})
	}
}