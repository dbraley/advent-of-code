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
			args: args{fileName: "testdata/basic"},
			want: [][]string{
				{"COM", "B"},
				{"B", "C"},
				{"C", "D"},
				{"D", "E"},
				{"E", "F"},
				{"B", "G"},
				{"G", "H"},
				{"D", "I"},
				{"E", "J"},
				{"J", "K"},
				{"K", "L"},
			},
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

func TestToNameOrbitMap(t *testing.T) {
	type o struct {
		parentName string
	}
	getParentName := func(orbiter *Orbit) string {
		if orbiter.parent != nil {
			return orbiter.parent.name
		}
		return ""
	}
	type args struct {
		input [][]string
	}
	tests := []struct {
		name string
		args args
		want map[string]o
	}{
		{
			name: "Empty",
			args: args{[][]string{}},
			want: map[string]o{
				"COM": {""},
			},
		},
		{
			name: "One Orbiter",
			args: args{[][]string{{"COM", "B"}}},
			want: map[string]o{
				"COM": {""},
				"B":   {"COM"},
			},
		},
		{
			name: "Double Orbiter",
			args: args{[][]string{{"COM", "B"}, {"B", "C"}}},
			want: map[string]o{
				"COM": {""},
				"B":   {"COM"},
				"C":   {"B"},
			},
		},
		{
			name: "Double Orbiter Out of Order",
			args: args{[][]string{{"B", "C"}, {"COM", "B"}}},
			want: map[string]o{
				"COM": {""},
				"B":   {"COM"},
				"C":   {"B"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := ToNameOrbitMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ToNameOrbitMap() = %v, want %v", got, tt.want)
			//}
			got := ToNameOrbitMap(tt.args.input)
			if len(got) != len(tt.want) {
				t.Errorf("ToNameOrbitMap() = %v, want %v", got, tt.want)
			}
			for name, wantOrbiter := range tt.want {
				gotOrbiter, ok := got[name]
				if !ok {
					t.Errorf("Did not get wanted orbit with name %v", name)
				}
				if wantOrbiter.parentName != getParentName(gotOrbiter) {
					t.Errorf("Expected %v to have parent %v, was %v", name, wantOrbiter.parentName, getParentName(gotOrbiter))
				}
			}
		})
	}
}

func TestOrbit_getDepth(t *testing.T) {

	tests := []struct {
		name    string
		orbiter Orbit
		want    int
	}{
		{
			name: "COM",
			orbiter: Orbit{
				name:   "COM",
				parent: nil,
				depth:  0,
			},
			want: 0,
		},
		{
			name: "B",
			orbiter: Orbit{
				name:   "B",
				parent: &Orbit{"COM", nil, 0},
				depth:  -1,
			},
			want: 1,
		},
		{
			name: "C",
			orbiter: Orbit{
				name:   "B",
				parent: &Orbit{"B", &Orbit{"COM", nil, 0}, -1},
				depth:  -1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Orbit{
				name:   tt.orbiter.name,
				parent: tt.orbiter.parent,
				depth:  tt.orbiter.depth,
			}
			if got := o.getDepth(); got != tt.want {
				t.Errorf("getDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateChecksum(t *testing.T) {
	basic, _ := Read("testdata/basic")
	basicOrbits := ToNameOrbitMap(basic)
	type args struct {
		orbits map[string]*Orbit
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Basic",
			args: args{
				orbits: basicOrbits,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateChecksum(tt.args.orbits); got != tt.want {
				t.Errorf("CalculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComNameSlice(t *testing.T) {
	type args struct {
		orbit *Orbit
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Basic Test",
			args: args{
				&Orbit{
					name:   "C",
					parent: &Orbit{"B", &Orbit{"COM", nil, 0}, -1},
				},
			},
			want: []string{"C", "B", "COM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToComNameSlice(tt.args.orbit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToComNameSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinTransfers(t *testing.T) {
	type args struct {
		you   []string
		santa []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Basic",
			args: args{
				you:   []string{"YOU", "COM"},
				santa: []string{"SAN", "COM"},
			},
			want: 0,
		},
		{
			name: "More Complex",
			args: args{
				you:   []string{"YOU", "B", "C", "X", "Y", "Z", "COM"},
				santa: []string{"SAN", "D", "E", "F", "X", "Y", "Z", "COM"},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinTransfers(tt.args.you, tt.args.santa); got != tt.want {
				t.Errorf("MinTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
