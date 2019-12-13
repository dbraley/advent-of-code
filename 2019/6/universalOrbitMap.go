package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	read, err := Read("input")
	if err != nil {
		fmt.Printf("Failed %+v\n", err)
		return
	}
	orbitMap := ToNameOrbitMap(read)
	checksum := CalculateChecksum(orbitMap)
	fmt.Printf("Checksum: %v\n", checksum)

	youToCom := ToComNameSlice(orbitMap["YOU"])
	santaToCom := ToComNameSlice(orbitMap["SAN"])
	//fmt.Println(youToCom)
	//fmt.Println(santaToCom)
	fmt.Println("Minimum Transfers:", MinTransfers(youToCom, santaToCom))
}

func Read(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ')'
	return reader.ReadAll()
}

type Orbit struct {
	name   string
	parent *Orbit
	depth  int
}

func (o *Orbit) getDepth() int {
	if o.parent != nil {
		o.depth = 1 + o.parent.getDepth()
		return o.depth
	}
	return 0
}

func ToNameOrbitMap(input [][]string) map[string]*Orbit {
	nameOrbit := make(map[string]*Orbit)
	nameOrbit["COM"] = &Orbit{
		name:   "COM",
		parent: nil,
		depth:  0,
	}
	for _, orbitPair := range input {
		if len(orbitPair) != 2 {
			panic("Bad Orbit Pair")
		}
		parent, ok := nameOrbit[orbitPair[0]]
		if !ok {
			parent = &Orbit{
				name:  orbitPair[0],
				depth: -1,
			}
			nameOrbit[orbitPair[0]] = parent
		}
		child, ok := nameOrbit[orbitPair[1]]
		if !ok {
			child = &Orbit{
				name:  orbitPair[1],
				depth: -1,
			}
			nameOrbit[orbitPair[1]] = child
		}
		child.parent = parent
	}
	return nameOrbit
}

func CalculateChecksum(orbits map[string]*Orbit) int {
	var total int
	for _, orbit := range orbits {
		total += orbit.getDepth()
	}
	return total
}

func ToComNameSlice(orbit *Orbit) []string {
	var comToOrbitNames []string
	for cur := orbit; cur != nil; cur = cur.parent {
		comToOrbitNames = append(comToOrbitNames, cur.name)
	}
	return comToOrbitNames
}

func MinTransfers(you, santa []string) int {
	for i, j := len(you)-1, len(santa)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		//fmt.Printf("Comparing %v:%v to %v:%v => %v\n", you, i, santa, j, you[i] != santa[j])
		if you[i] != santa[j] {
			//fmt.Printf("Done. %v:%v == %v:%v\n", you[i:i+2], i, santa[j:j+2], j)
			return i + j
		}
	}
	return 0
}
