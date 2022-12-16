package main

import (
	"errors"
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//ept1, ept2 := day16("example.txt")
	//fmt.Printf("Input:\n\tPt1: %v\n\tPt2: %v\n", ept1, ept2)

	ipt1, ipt2 := day16("input.txt")
	fmt.Printf("Input:\n\tPt1: %v\n\tPt2: %v\n", ipt1, ipt2)

}

func day16(fileName string) (pt1 string, pt2 string) {
	input, err := file.ReadFile("day16/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("%d\n", len(input))

	r := regexp.MustCompile(`Valve (?P<name>\S{2}) has flow rate=(?P<sy>\d+); tunnels? leads? to valves? (?P<neighbors>.*)`)

	nodeMap := make(map[string]Node)
	usefulValves := []*Node{}
	for _, line := range input {
		nodeMatch := r.FindStringSubmatch(line)

		nodeName := nodeMatch[1]
		nodePressure, _ := strconv.Atoi(nodeMatch[2])
		nodeNeighbors := strings.Split(nodeMatch[3], ", ")

		node := Node{
			Name:      nodeName,
			Pressure:  nodePressure,
			Neighbors: nodeNeighbors,
		}
		nodeMap[node.Name] = node
		if nodePressure > 0 {
			usefulValves = append(usefulValves, &node)
		}
		//fmt.Println()
	}

	bestP1 := findBest(1000000, 30, usefulValves, nodeMap, 1)
	pt1 = fmt.Sprintf("%d", bestP1.Total)

	bestP2 := findBest(1000000, 26, usefulValves, nodeMap, 2)
	pt2 = fmt.Sprintf("%d", bestP2.Total)

	return pt1, pt2
}

func findBest(maxRuns int, maxTime int, usefulValves []*Node, nodeMap map[string]Node, runners int) Run {
	done := []Run{}
	runs := []Run{{Current: "AA", Released: make(map[string]int), PathTaken: "AA"}}
	for i := 0; i < maxRuns; i++ {
		if len(runs) <= i {
			break
		}
		curRun := runs[i]
		timeLeft := maxTime - curRun.MinutesSpent
		// We could just chill until we're out of time, but maybe only if we've got some pressure releasing?
		// But also lets make sure we leave some for our partners
		if len(curRun.Released) >= 1 && len(curRun.Released) <= len(usefulValves)-runners+1 {
			deadRun := Run{
				MinutesSpent: maxTime,
				Current:      curRun.Current,
				Released:     curRun.copyReleased(),
				PerMin:       curRun.PerMin,
				Total:        curRun.Total + curRun.PerMin*timeLeft,
				PathTaken:    fmt.Sprintf("%v -> wait", curRun.PathTaken),
			}
			done = append(done, deadRun)
		}

		// Or head for another valuable node

		// Figure out which valuable nodes are left
		var stillDesirableNodes []*Node
		for _, uv := range usefulValves {
			if _, released := curRun.Released[uv.Name]; !released {
				stillDesirableNodes = append(stillDesirableNodes, uv)
			}
		}
		//fmt.Println(stillDesirableNodes)

		// Pick one and go
		for _, target := range stillDesirableNodes {
			shortest, err := shortestTo(curRun, target, nodeMap)
			if err != nil {
				fmt.Printf("No path from %v to %v?\n", curRun.Current, target.Name)
			}
			if len(shortest) <= timeLeft {
				//fmt.Printf("%v", shortest)
				newRun := Run{
					MinutesSpent: curRun.MinutesSpent + len(shortest),
					Current:      target.Name,
					Released:     curRun.copyReleased(),
					PerMin:       curRun.PerMin + target.Pressure,
					Total:        curRun.Total + curRun.PerMin*len(shortest),
					PathTaken:    fmt.Sprintf("%v -> %v -> open(%v)", curRun.PathTaken, strings.Join(shortest[1:], " -> "), target.Name),
				}
				newRun.Released[target.Name] = target.Pressure
				runs = append(runs, newRun)
			}
		}
	}
	fmt.Printf("Runs: %d (maybeErr=%v)\n", len(runs), len(runs) > maxRuns)
	best := Run{}
	for i, run := range done {
		if runners == 1 {
			if run.Total > best.Total {
				best = run
			}
		} else {
			//unvisitedNodes := []*Node{}
			//for _, node := range usefulValves {
			//	if _, visitted := run.Released[node.Name]; !visitted {
			//		unvisitedNodes = append(unvisitedNodes, node)
			//	}
			//}
			//other := findBest(maxRuns, maxTime, unvisitedNodes, nodeMap, 1)

			// We don't need to go recursive, we've got all the potential solutions already!
			bestPair := Run{}
			for j := len(done) - 1; j > i; j-- {
				pair := done[j]

				// If the pair is valid
				valid := true
				for k, _ := range run.Released {
					if _, overlap := pair.Released[k]; overlap {
						valid = false
					}
				}

				if valid && pair.Total > bestPair.Total {
					bestPair = pair
				}
			}
			if run.Total+bestPair.Total > best.Total {
				newBest := Run{
					MinutesSpent: run.MinutesSpent + bestPair.MinutesSpent,
					Current:      fmt.Sprintf("%v,%v", run.Current, bestPair.Current),
					PathTaken:    fmt.Sprintf("%v\n%v", run.PathTaken, bestPair.PathTaken),
					Released:     nil,
					PerMin:       run.PerMin + bestPair.PerMin,
					Total:        run.Total + bestPair.Total,
				}
				best = newBest
			}
		}
	}
	return best
}

func shortestTo(curRun Run, target *Node, nodeMap map[string]Node) ([]string, error) {
	//fmt.Printf("%v: Trying to find Path from %v -> %v\n", curRun.PathTaken, curRun.Current, target.Name)
	paths := [][]string{{curRun.Current}}
	for true {
		if len(paths) == 0 {
			break
		}
		path := paths[0]
		paths = paths[1:]
		lastOnPath := nodeMap[path[len(path)-1]]
		if lastOnPath.Name == target.Name {
			return path, nil
		} else {
			for _, neighbor := range lastOnPath.Neighbors {
				if !contains(path, neighbor) {
					//newPath := append(path, neighbor)
					newPath := make([]string, len(path)+1)
					copy(newPath, path)
					newPath[len(path)] = neighbor
					paths = append(paths, newPath)
				}
			}
		}
	}
	return []string{}, errors.New("no path found!")
}

type Node struct {
	Name      string
	Pressure  int
	Neighbors []string
}

type Run struct {
	MinutesSpent int
	Current      string
	PathTaken    string

	Released map[string]int
	PerMin   int
	Total    int
}

func (r *Run) copyReleased() map[string]int {
	cpy := make(map[string]int)
	for k, v := range r.Released {
		cpy[k] = v
	}
	return cpy
}

func contains(path []string, node string) bool {
	for _, pn := range path {
		if node == pn {
			return true
		}
	}
	return false
}
