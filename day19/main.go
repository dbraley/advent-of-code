package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"regexp"
	"strconv"
)

var best = make(map[int]int)

func main() {
	agg := 0
	for i := 1; i <= 32; i++ {
		agg += i
		best[i] = agg
	}
	//fmt.Println("Example:")
	//ept1, ept2 := day19("example.txt")
	//fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)
	//
	fmt.Println("Input:")
	ipt1, ipt2 := day19("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day19(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day19/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	bps := makeBlueprints(input)

	var topLevelRuns []run
	for _, bp := range bps {
		if bp.number < 4 {
			topLevelRuns = append(topLevelRuns,
				NewFromBluePrint(bp, 1),
				NewFromBluePrint(bp, 2),
			)
		}
	}

	maxGeodes := make(map[int]int)
	for _, tlrun := range topLevelRuns {
		runs := []run{tlrun}
		i := 0
		for true {
			r := runs[i]
			//fmt.Printf("Passing time for Blueprint %d at minute %d with next robot %d\n", r.bp.number, r.minutesPast, r.nextRobot)
			bestForBp := maxGeodes[r.bp.number]
			for m := r.minutesPast; m < 32; m++ {
				bestRemaining := r.BestRemaining(32)
				if bestRemaining <= bestForBp {
					break
				}
				nrs := r.passMinute()
				if len(nrs) > 0 {
					runs = append(runs, nrs...)
				}
			}
			if r.geodes > maxGeodes[r.bp.number] {
				maxGeodes[r.bp.number] = r.geodes
			}

			i++
			if i >= len(runs) {
				fmt.Printf("Breaking after %d runs.\n", i)
				break
			}
		}
	}

	pt2 = 1
	for k, v := range maxGeodes {
		fmt.Printf("BP %d MAX %d = score %d\n", k, v, k*v)
		pt1 += k * v
		pt2 *= v
	}

	return pt1, pt2
}

func makeBlueprints(input []string) []blueprint {
	r := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

	blueprints := make([]blueprint, len(input))
	for i, line := range input {
		m := r.FindStringSubmatch(line)
		blueprints[i] = blueprint{
			number:     intOrPanic(m[1]),
			r1OreCost:  intOrPanic(m[2]),
			r2OreCost:  intOrPanic(m[3]),
			r3OreCost:  intOrPanic(m[4]),
			r3ClayCost: intOrPanic(m[5]),
			r4OreCost:  intOrPanic(m[6]),
			r4ObsCost:  intOrPanic(m[7]),
		}
	}

	return blueprints
}

func intOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type blueprint struct {
	number     int
	r1OreCost  int
	r2OreCost  int
	r3OreCost  int
	r3ClayCost int
	r4OreCost  int
	r4ObsCost  int
}

type run struct {
	bp          blueprint
	minutesPast int
	nextRobot   int
	building    int
	r1          int
	r2          int
	r3          int
	r4          int
	ore         int
	clay        int
	obs         int
	geodes      int
}

func NewFromBluePrint(bp blueprint, nextRobot int) run {
	return run{
		bp:          bp,
		minutesPast: 0,
		nextRobot:   nextRobot,
		building:    -1,
		r1:          1,
		r2:          0,
		r3:          0,
		r4:          0,
		ore:         0,
		clay:        0,
		obs:         0,
		geodes:      0,
	}
}
func NewFromRun(r *run, nextRobot int) run {
	return run{
		bp:          r.bp,
		minutesPast: r.minutesPast,
		nextRobot:   nextRobot,
		building:    -1,
		r1:          r.r1,
		r2:          r.r2,
		r3:          r.r3,
		r4:          r.r4,
		ore:         r.ore,
		clay:        r.clay,
		obs:         r.obs,
		geodes:      r.geodes,
	}
}

func (r *run) passMinute() []run {
	r.minutesPast++
	//fmt.Print("== Minute %d ==\n", r.minutesPast)
	if r.canBuild() {
		// Don't actually need to do anything here I don't think
	}
	r.updateMaterials()
	if r.building != -1 {
		switch r.building {
		case 1:
			r.r1++
			//fmt.Print("The new ore-collecting robot is ready; you now have %d of them.\n", r.r1)
		case 2:
			r.r2++
			//fmt.Print("The new clay-collecting robot is ready; you now have %d of them.\n", r.r2)
		case 3:
			r.r3++
			//fmt.Print("The new obsidian-collecting robot is ready; you now have %d of them.\n", r.r3)
		case 4:
			r.r4++
			//fmt.Print("The new geode-collecting robot is ready; you now have %d of them.\n", r.r4)
		default:
			fmt.Printf("!!!No idea what type of robot %d is!!!", r.building)
		}
		r.building = -1

		nextRobotCouldBe := []int{}
		// More ore production?
		if r.r1 < max(r.bp.r4OreCost, r.bp.r3OreCost, r.bp.r2OreCost, r.bp.r1OreCost) {
			nextRobotCouldBe = append(nextRobotCouldBe, 1)
		}
		// More clay production?
		if r.r2 < r.bp.r3ClayCost {
			nextRobotCouldBe = append(nextRobotCouldBe, 2)
		}
		// More obsidian production?
		if r.r2 > 0 && r.r3 < r.bp.r4ObsCost {
			nextRobotCouldBe = append(nextRobotCouldBe, 3)
		}
		if r.r3 > 0 {
			nextRobotCouldBe = append(nextRobotCouldBe, 4)
		}
		nextRun := make([]run, len(nextRobotCouldBe)-1)
		for i := 0; i < len(nextRobotCouldBe)-1; i++ {
			nextRun[i] = NewFromRun(r, nextRobotCouldBe[i])
		}
		r.nextRobot = nextRobotCouldBe[len(nextRobotCouldBe)-1]
		return nextRun
	}
	return []run{}
}

func (r *run) canBuild() bool {
	switch r.nextRobot {
	case 1:
		if r.ore >= r.bp.r1OreCost {
			//fmt.Print("Spend %d ore to start building a ore-collecting robot.\n", r.bp.r1OreCost)
			r.ore -= r.bp.r1OreCost
			r.nextRobot = -1
			r.building = 1
			return true
		}
	case 2:
		if r.ore >= r.bp.r2OreCost {
			//fmt.Print("Spend %d ore to start building a clay-collecting robot.\n", r.bp.r2OreCost)
			r.ore -= r.bp.r2OreCost
			r.nextRobot = -1
			r.building = 2
			return true
		}
	case 3:
		if r.ore >= r.bp.r3OreCost && r.clay >= r.bp.r3ClayCost {
			//fmt.Print("Spend %d ore and %d clay to start building an obsidian-collecting robot.\n", r.bp.r3OreCost, r.bp.r3ClayCost)
			r.ore -= r.bp.r3OreCost
			r.clay -= r.bp.r3ClayCost
			r.nextRobot = -1
			r.building = 3
			return true
		}
	case 4:
		if r.ore >= r.bp.r4OreCost && r.obs >= r.bp.r4ObsCost {
			//fmt.Print("Spend %d ore and %d obsidian to start building a geode-collecting robot.\n", r.bp.r4OreCost, r.bp.r4ObsCost)
			r.ore -= r.bp.r4OreCost
			r.obs -= r.bp.r4ObsCost
			r.nextRobot = -1
			r.building = 4
			return true
		}
	default:
		fmt.Printf("!!! Unknown Robot type %d!!!", r.nextRobot)
	}
	return false
}

func (r *run) updateMaterials() {
	if r.r1 > 0 {
		r.ore += r.r1
		//fmt.Print("%d ore-collecting robot(s) collects %d ore; you now have %d ore.\n", r.r1, r.r1, r.ore)
	}
	if r.r2 > 0 {
		r.clay += r.r2
		//fmt.Print("%d clay-collecting robot(s) collects %d clay; you now have %d clay.\n", r.r2, r.r2, r.clay)
	}
	if r.r3 > 0 {
		r.obs += r.r3
		//fmt.Print("%d obsidian-collecting robot(s) collects %d obsidian; you now have %d obsidian.\n", r.r3, r.r3, r.obs)
	}
	if r.r4 > 0 {
		r.geodes += r.r4
		//fmt.Print("%d geode-collecting robot(s) collects %d geode(s); you now have %d geodes.\n", r.r4, r.r4, r.geodes)
	}
}

func (r *run) BestRemaining(to int) int {
	turnsLeft := to - r.minutesPast
	willProduce := r.geodes + turnsLeft*r.r4
	mightProduce := best[turnsLeft]
	switch r.nextRobot {
	case 1:
		doneAfterTurns := (r.bp.r1OreCost - r.ore) / r.r1
		// no additional obsidian production until this bot is done, plus one for next bot
		mightProduce = best[turnsLeft-doneAfterTurns]
	case 2:
		doneAfterTurns := (r.bp.r2OreCost - r.ore) / r.r1
		// no additional obsidian production until this bot is done, plus one for next bot
		mightProduce = best[turnsLeft-doneAfterTurns]
	case 3:
		doneAfterTurns := max((r.bp.r3OreCost-r.ore)/r.r1, (r.bp.r3ClayCost-r.clay)/r.r2)
		// no additional obsidian production until this bot is done, plus one for next bot
		mightProduce = best[turnsLeft-doneAfterTurns]
	case 4:
		doneAfterTurns := max((r.bp.r4OreCost-r.ore)/r.r1, (r.bp.r4ObsCost-r.obs)/r.r3)
		// no additional obsidian production until this bot is done, plus one for next bot
		mightProduce = best[turnsLeft-doneAfterTurns]
	}
	bestRemaining := willProduce + mightProduce
	return bestRemaining
}

func max(a int, b ...int) int {
	m := a
	for _, other := range b {
		if other > m {
			m = other
		}
	}
	return m
}
